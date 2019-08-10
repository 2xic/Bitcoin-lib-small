
package tx

import (
	"testing"
	"strings"
	"bytes"
	"github.com/2xic/bip-39/secp256k1"
	"github.com/2xic/bip-39/base58"
	"encoding/hex"
)


func revert(input []byte) []byte{
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func Test_tx(t *testing.T) {

	//	basic tx example from https://bitcoin.org/en/developer-reference#raw-transaction-format

	inputs := []TransactionInput{}

	inputTX := []byte{0x7b,0x1e,0xab,0xe0,0x20,0x9b,0x1f,0xe7,0x94,0x12,0x45,0x75,0xef,0x80,0x70,0x57,0xc7,0x7a,0xda,0x21,0x38,0xae,0x4f,0xa8,0xd6,0xc4,0xde,0x03,0x98,0xa1,0x4f,0x3f}
	scripts := []byte{0x48,    0x30,0x45,0x02,0x21,0x00,0x89,0x49,0xf0,0xcb,0x40,0x00,0x94,0xad,0x2b,0x5e,0xb3,0x99,0xd5,0x9d,0x01,0xc1,0x4d,0x73,0xd8,0xfe,0x6e,0x96,0xdf,0x1a,0x71,0x50,0xde,0xb3,0x88,0xab,0x89,0x35,0x02,0x20,0x79,0x65,0x60,0x90,0xd7,0xf6,0xba,0xc4,0xc9,0xa9,0x4e,0x0a,0xad,0x31,0x1a,0x42,0x68,0xe0,0x82,0xa7,0x25,0xf8,0xae,0xae,0x05,0x73,0xfb,0x12,0xff,0x86,0x6a,0x5f,0x01}
	inputs = append(inputs, createInput(inputTX, 0, scripts))


	outputs := []OutputFormat{}
	scriptFormat := []byte{0x76, 0xa9, 0x14, 0xcb,0xc2, 0x0a, 0x76, 0x64, 0xf2, 0xf6, 0x9e, 0x53,0x55,0xaa,0x42,0x70,0x45,0xbc,0x15,0xe7,0xc6,0xc7,0x72,		0x88, 0xac}
	outputs = append(outputs, createOutput(4999990000, scriptFormat))

	tx := construct(inputs, outputs)
	if !(strings.Compare(tx, "01000000017b1eabe0209b1fe794124575ef807057c77ada2138ae4fa8d6c4de0398a14f3f00000000494830450221008949f0cb400094ad2b5eb399d59d01c14d73d8fe6e96df1a7150deb388ab8935022079656090d7f6bac4c9a94e0aad311a4268e082a725f8aeae0573fb12ff866a5f01ffffffff01f0ca052a010000001976a914cbc20a7664f2f69e5355aa427045bc15e7c6c77288ac00000000") == 0){
		t.Errorf("something is wrong with tx generator")
	}


//	fmt.Println(String2Bytes("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5603801b05e4b883e5bda9e7a59ee4bb99e9b1bcfabe6d6d1826d7955e515df6d4587632a80d1fc2bb8c965dc0d9f4417fefb7c328d0bcd6100000000000000000dde1bec70100004d696e6564206279207a6470777370ffffffff0100f90295000000001976a914c825a1ecf2a6830c4401620c3a16f1995057c2ab88ac00000000"))
//	fmt.Println()
	txid := revert(secp256k1.DoubleSha256(Serailizetx(String2Bytes("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5603801b05e4b883e5bda9e7a59ee4bb99e9b1bcfabe6d6d1826d7955e515df6d4587632a80d1fc2bb8c965dc0d9f4417fefb7c328d0bcd6100000000000000000dde1bec70100004d696e6564206279207a6470777370ffffffff0100f90295000000001976a914c825a1ecf2a6830c4401620c3a16f1995057c2ab88ac00000000"))))
	real_txid, _ := hex.DecodeString("843bf58e083503ea272f90fb24cfc1a5b3b905f32470eff09602295e6b58480d")
	//fmt.Println(txid)
	//fmt.Println(real_txid)
	if !(bytes.Compare(txid, real_txid) == 0){
		t.Errorf("something is wrong with tx generator")
	}
}

func Test_adress(t *testing.T){
	//	0x14,0x53,0x99,0xc3,0x09,0x3d,0x31,0xe4,0xb0,0xaf,0x4b,0xe1,0x21,0x5d,0x59,0xb8,0x57,0xb8,0x61,0xad,0x5d
	address_hash := []byte{0x53,0x99,0xc3,0x09,0x3d,0x31,0xe4,0xb0,0xaf,0x4b,0xe1,0x21,0x5d,0x59,0xb8,0x57,0xb8,0x61,0xad,0x5d}
	
	hash160 := make([]byte, 21)

	compressedAddress := make([]byte, 25)
	
	secp256k1.ExtendSlice(hash160, address_hash, 1)
	secp256k1.ExtendSlice(compressedAddress, hash160, 0)
	secp256k1.ExtendSlice(compressedAddress, secp256k1.DoubleSha256(hash160), len(hash160))

	//fmt.Println(base58.ConvertBytes((compressedAddress)))

	pay_2_script_hash := base58.ConvertBytes((compressedAddress))

	if !(strings.Compare(pay_2_script_hash, "18d3HV2bm94UyY4a9DrPfoZ17sXuiDQq2B") == 0){
		t.Errorf("something is wrong with P2PKH")
	}

	/*
	address_hash := []byte{0x89,0xAB,0xCD,0xEF,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA,0xAB,0xBA}
	//address_hash = secp256k1.Hash160(address_hash)
	fmt.Println(len(address_hash))

	address_hash = append(address_hash, secp256k1.ReturnCheckSum(address_hash)...)

	adress := []byte{}//{0x0}
	adress = append(adress, address_hash...)
	fmt.Println(len(address_hash))
	fmt.Println(base58.ConvertBytes(secp256k1.Hash160(adress)))
*/
}

