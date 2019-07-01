package bip32

import(
	"encoding/hex"
	"crypto/hmac"
	"math/big"
	"bytes"
	"crypto/sha512"
	"github.com/2xic/bip-39/base58"
	"github.com/2xic/bip-39/secp256k1"
	"math"
	"encoding/binary"
)

var (
	versionMainetPublic, _ = hex.DecodeString("0488B21E")
	versionMainetPrivate, _ = hex.DecodeString("0488ADE4")

	versionTestnetPublic, _ = hex.DecodeString("0x043587CF")
	versionTestnetPrivate, _ = hex.DecodeString("0x04358394")
)

// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#serialization-format
type keyStruct struct{
	Version []byte
	Depth byte
	Fingerprint []byte
	ChildNumber []byte
	ChainCode []byte
	KeyData []byte 
	Private bool
}

func Serialize(key *keyStruct) []byte{
	buffer := new(bytes.Buffer)
	buffer.Write(key.Version)
	buffer.WriteByte(key.Depth)
	buffer.Write(key.Fingerprint)
	buffer.Write(key.ChildNumber)
	buffer.Write(key.ChainCode)

	keyBytes := key.KeyData
	if(key.Private){
		keyBytes = append([]byte{0x0}, keyBytes...)
	}
	buffer.Write(keyBytes)
	return buffer.Bytes()
}

//	https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#master-key-generation
func MasterPrivateKey(seed []byte) *keyStruct{
	hmac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	_, err := hmac.Write(seed)
	if err != nil {
		return nil
	}

	hash := hmac.Sum(nil)

	left := hash[:32]
	right := hash[32:]

	secp256k1.CheckPrivateKey(left)

	return &keyStruct{
		versionMainetPrivate,
		0,
		make([]byte, 4),
		make([]byte, 4),
		right,
		left,
		true}
}

func MasterPublicKey(key *keyStruct) *keyStruct{
	if(!key.Private){
		return nil
	}

	keyDataBytes := key.KeyData
	
	PrivateKey := new(big.Int)
	PrivateKey.SetBytes(keyDataBytes)

	return &keyStruct{
		versionMainetPublic,
		key.Depth,
		key.Fingerprint,
		key.ChildNumber,
		key.ChainCode,
		secp256k1.CompressPublicKey(secp256k1.GetPublicPoint(PrivateKey)),
		false}
}

func FourByteAlign(number int) []byte{
	output := make([]byte, 4)
	binary.BigEndian.PutUint32(output, uint32(number))
	return output
}

func Private2CompressedPublic(key []byte) []byte{
	PrivateKey := new(big.Int)
	PrivateKey.SetBytes(key)
	return secp256k1.CompressPublicKey(secp256k1.GetPublicPoint(PrivateKey))
}

func ChildKey(key *keyStruct, index int) (*keyStruct){
	var data [] byte
	/*
		Each key is based off a parent key
	*/
	if(math.Pow(2, 31) <= float64(index)){
		data = append([]byte{0x0}, key.KeyData...)
	}else{
		if(key.Private){
			data = Private2CompressedPublic(key.KeyData)
		}else{
			data = key.KeyData
		}
	}
	data = append(data, FourByteAlign(index)...)

	hmac := hmac.New(sha512.New, key.ChainCode)
	_, err := hmac.Write(data)
	if(err != nil){
		panic(err)
	}
	hash := hmac.Sum(nil)
	
	//	as defined in the spec
	left := hash[:32]
	right := hash[32:]

	if(key.Private){
		NewChildKey := &keyStruct{
			versionMainetPrivate,
			key.Depth + 1,
			secp256k1.FingerPrint160(Private2CompressedPublic(key.KeyData)),
			FourByteAlign(index),
			right,
			secp256k1.CombinePrivateKeys(left, key.KeyData),
			true}
		return NewChildKey
	}else{
		NewChildKey := &keyStruct{
			versionMainetPublic,
			key.Depth + 1,
			secp256k1.FingerPrint160(key.KeyData),
			FourByteAlign(index),
			right,
			secp256k1.CombinePublicKeys(Private2CompressedPublic(left), key.KeyData),
			false}
		return NewChildKey
	}
	return nil
}

func(key *keyStruct) Readable() string{
	return base58.ConvertBytes(secp256k1.CheckSum(Serialize(key)))
}

func(key *keyStruct) ReadableAddress() string{
	keyDataBytes := key.KeyData
	PrivateKey := new(big.Int)
	PrivateKey.SetBytes(keyDataBytes)
	return base58.ConvertBytes(secp256k1.CompressAddress(secp256k1.GetPublicPoint(PrivateKey)))
}
