package bip32

import(
	"encoding/hex"
	"crypto/hmac"
	"fmt"
	"math/big"
	"bytes"
	"crypto/sha512"
	"github.com/2xic/bip-39/base58"
	"github.com/2xic/bip-39/secp256k1"
)

var (
	version_mainet_public, _ = hex.DecodeString("0488B21E")
	version_mainet_private, _ = hex.DecodeString("0488ADE4")
	version_testnet, _ = hex.DecodeString("0x043587CF")
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

func allZeroBytes(stream []byte) bool {
	for _, byte := range stream {
		if byte != 0 {
			return false
		}
	}
	return true
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

	if(allZeroBytes(left) || len(left) != 32 || 0 <= bytes.Compare(left, secp256k1.N.Bytes())){
		fmt.Println("mama mia")
	}

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

func(key *keyStruct) Readable() string{
	return base58.ConvertBytes(secp256k1.CheckSum(Serialize(key)))
}

func(key *keyStruct) ReadableAddress() string{
	keyDataBytes := key.KeyData
	PrivateKey := new(big.Int)
	PrivateKey.SetBytes(keyDataBytes)
	return base58.ConvertBytes(secp256k1.CompressAddress(secp256k1.GetPublicPoint(PrivateKey)))
}
