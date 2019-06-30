package secp256k1

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"hash"
)

func ExtendSlice(buffer []byte, extended []byte, start int)[]byte{
	for i := 0; (i + start) < len(buffer) && i < len(extended); i++{
		buffer[start + i] = extended[i];
	}
	return buffer
}

func RunHash(buffer []byte, hashType hash.Hash) []byte {
	hashType.Write(buffer)
	return hashType.Sum(nil)
}

func Hash160(buffer []byte) []byte {
	return RunHash(RunHash(buffer, sha256.New()), ripemd160.New())
}

func DoubleSha256(buffer []byte) []byte{
	return RunHash(RunHash(buffer, sha256.New()), sha256.New())
}

func CheckSum(buffer []byte) []byte{
	return append(buffer, DoubleSha256(buffer)[:4]...)
}

func ReturnCheckSum(buffer []byte) []byte{
	return DoubleSha256(buffer)[:4]
}

func CompressPublicKey(private *CurvePoint) []byte {
	data := make([]byte, 33)
	data[0] = byte(0x2) + byte(private.GetY().Bit(0))
	ExtendSlice(data, private.GetX().Bytes(), 1)	
	return data
}

func CompressAddress(point *CurvePoint) []byte{
	publicKey := CompressPublicKey(point)

	hash160 := make([]byte, 21)
	compressedAddress := make([]byte, 25)
	
	ExtendSlice(hash160, Hash160(publicKey), 1)
	ExtendSlice(compressedAddress, hash160, 0)
	ExtendSlice(compressedAddress, DoubleSha256(hash160), len(hash160))

	return compressedAddress
}
