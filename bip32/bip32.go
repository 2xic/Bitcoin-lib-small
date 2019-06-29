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
	version_mainet_public, _ = hex.DecodeString("0488B21E")
	version_mainet_private, _ = hex.DecodeString("0488ADE4")

	version_testnet_public, _ = hex.DecodeString("0x043587CF")
	version_testnet_private, _ = hex.DecodeString("0x04358394")
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

//	xprv9uEKuUoGik8t

func Serialize(key *keyStruct) []byte{
	buffer := new(bytes.Buffer)
	buffer.Write(key.Version)
	buffer.WriteByte(key.Depth)
	buffer.Write(key.Fingerprint)
	buffer.Write(key.ChildNumber)
	buffer.Write(key.ChainCode)
//	buffer.Write(make([]byte, 32))

	keyBytes := key.KeyData
	if(key.Private){
		keyBytes = append([]byte{0x0}, keyBytes...)
	}
	buffer.Write(keyBytes)
	return buffer.Bytes()
}

func allZeroBytes(stream []byte) bool {
	for _, byte := range stream {
		if byte != 0 {
			return false
		}
	}
	return true
}

func CheckPrivateKey(input []byte) bool{
	if(allZeroBytes(input) || len(input) != 32 || 0 <= bytes.Compare(input, secp256k1.N.Bytes())){
		panic("invalid private key")
	}
	return true;
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

	CheckPrivateKey(left)

	return &keyStruct{
		version_mainet_private,
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
		version_mainet_public,
		key.Depth,
		key.Fingerprint,
		key.ChildNumber,
		key.ChainCode,
		secp256k1.CompressPublicKey(secp256k1.GetPublicPoint(PrivateKey)),
		false}
}

func parsePrivateKey(left[]byte, parrent[]byte) []byte{
	var x = new(big.Int).SetBytes(left)
	var y = new(big.Int).SetBytes(parrent)
	var z = new(big.Int)

	z.Add(x, y)
	z.Mod(z, secp256k1.N)

	key := z.Bytes()
//	padding := make([]byte, 32-len(key))
//	key = append(key, padding...)

	CheckPrivateKey(key)
	return key
}

func FourByteAlign(number int) []byte{
	output := make([]byte, 4)
	binary.BigEndian.PutUint32(output, uint32(number))
	return output
}

func ChildKey(key *keyStruct, index int) (*keyStruct){
	var data [] byte
	
	if(float64(index) <= math.Pow(2, 31)){
		if(key.Private){
			data = append([]byte{0x0}, key.KeyData...)
		}else{
			panic("failure, hardened children is not allowed")
		}
	}


	data = append(data, FourByteAlign(index)...)

	hmac := hmac.New(sha512.New, key.ChainCode)
	_, err := hmac.Write(data)
	if(err != nil){
		panic(err)
	}

	hash := hmac.Sum(nil)
	
	left := hash[:32]
	right := hash[32:]

//	NewChildKey := nil
	if(key.Private){
		PrivateKey := new(big.Int)
		PrivateKey.SetBytes(key.KeyData)
		Fingerprint := secp256k1.Hash160(secp256k1.CompressPublicKey(secp256k1.GetPublicPoint(PrivateKey)))[:4]

		NewChildKey := &keyStruct{
			version_mainet_private,
			key.Depth + 1,
			Fingerprint,
			FourByteAlign(index),
			right,
			parsePrivateKey(left, key.KeyData),
			true}
		return NewChildKey
	}else{
//		newKeyData := secp256k1.CompressPublicKey(secp256k1.GetPublicPoint(new(big.Int).SetBytes(left)))

		panic("need add new public key with old.... (need to read some crypto)")

		/*
		NewChildKey = &keyStruct{
			version_mainet_public,
			key.Depth + 1,
			make([]byte, 4),
			FourByteAlign(index),
			right,
			parsePrivateKey(left, key.KeyData),
			true}		*/
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
