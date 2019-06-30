package bip39

import (
	"testing"
	"bytes"
	"encoding/hex"
)

func generateSetBytes(value byte) (stream []byte){
	byteStream := make([]byte, 32)
	for i:= 0; i < len(byteStream); i++{
		byteStream[i] = value
	}
	return byteStream
}

func testVerify(words string) bool{
	valid, err := verifyMnemonic(words)
	if(err != nil){
		return false
	}
	return valid
}

func Test(t *testing.T) {
	words, err := GenerateMnemonicBytes(generateSetBytes(0))
	if err != nil{
		t.Errorf("Error with vector 0")		
	}
	if !testVerify(words) || words != "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art"{
		t.Errorf("Error with vector 0")
	}

	words, err = GenerateMnemonicBytes(generateSetBytes(0x7F))
	if err != nil{
		t.Errorf("Error with vector 1")		
	}
	if !testVerify(words) || words != "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title"{
		t.Errorf("Error with vector 1, got %s", words)
	}

	words, err = GenerateMnemonicBytes(generateSetBytes(0x80))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}
	if !testVerify(words) || words != "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless"{
		t.Errorf("Error with vector 2, got %s", words)
	}

	words, err = GenerateMnemonicBytes(generateSetBytes(0xFF))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}

	if !testVerify(words) || words != "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote"{
		t.Errorf("Error with vector 2, got %s", words)
	}


	words, err = GenerateMnemonicBytes(generateSetBytes(0))
	seed := Memonic2Seed(words, "TREZOR")
	truth, _ := hex.DecodeString("bda85446c68413707090a52022edd26a1c9462295029f2e60cd7c4f2bbd3097170af7a4d73245cafa9c3cca8d561a7c3de6f5d4a10be8ed2a5e608d68f92fcc8")
	if(!(bytes.Compare(seed, truth) == 0)){
		t.Errorf("Error with seed 1")
	}

	if(bytes.Compare(generateRandomBytes(10), generateRandomBytes(10)) == 0){
		t.Errorf("Error with random generator")
	}

	words, err = GenerateMnemonic()
	if(err != nil){
		t.Errorf("Error with memonic generator")		
	}

}



