package secp256k1

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"hash"
	"math/big"
	"bytes"
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

func FingerPrint160(buffer []byte) []byte{
	return Hash160(buffer)[:4]
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

/*
	https://github.com/tyler-smith/go-bip32/blob/master/utils.go#L147
	https://crypto.stackexchange.com/questions/8914/ecdsa-compressed-public-key-point-back-to-uncompressed-public-key-point/8916#8916
*/
func ExpandPublicKey(key []byte) (x *big.Int, y *big.Int){
	x = big.NewInt(0)
	y = big.NewInt(0)
	
	x.SetBytes(key[1:])

	ySquared := big.NewInt(0)
	ySquared.Exp(x, big.NewInt(3), nil)
	ySquared.Add(ySquared, B)

	y.ModSqrt(ySquared, P)

	Ymod2 := big.NewInt(0)
	Ymod2.Mod(y, big.NewInt(2))

	signY := uint64(key[0]) - 2
	if signY != Ymod2.Uint64() {
		y.Sub(P, y)
	}

	if(x.Sign() == 0 || y.Sign() == 0){
		panic("bad key")
	}

	return x, y
}

func CombinePublicKeys(key1 []byte, key2 []byte) []byte{
	x1, y1 := ExpandPublicKey(key1)
	x2, y2 := ExpandPublicKey(key2)

	point1,_ := S256Point(x1, y1)
	point2,_ := S256Point(x2, y2)

	results, _ := point1.AddReal(point2)

	return CompressPublicKey(results)
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
	if(allZeroBytes(input) || len(input) != 32 || 0 <= bytes.Compare(input, N.Bytes())){
		panic("invalid private key")
	}
	return true;
}

func CombinePrivateKeys(left[]byte, parrent[]byte) []byte{
	var x = new(big.Int).SetBytes(left)
	var y = new(big.Int).SetBytes(parrent)
	var z = new(big.Int)

	z.Add(x, y)
	z.Mod(z, N)

	key := z.Bytes()

	CheckPrivateKey(key)
	return key
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
