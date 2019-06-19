package main

import (
	"testing"
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
	words, err := generateMnemonicBytes(generateSetBytes(0))
	if err != nil{
		t.Errorf("Error with vector 0")		
	}
	if !testVerify(words) || words != "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art"{
		t.Errorf("Error with vector 0")
	}

	words, err = generateMnemonicBytes(generateSetBytes(0x7F))
	if err != nil{
		t.Errorf("Error with vector 1")		
	}
	if !testVerify(words) || words != "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title"{
		t.Errorf("Error with vector 1, got %s", words)
	}

	words, err = generateMnemonicBytes(generateSetBytes(0x80))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}
	if !testVerify(words) || words != "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless"{
		t.Errorf("Error with vector 2, got %s", words)
	}

	words, err = generateMnemonicBytes(generateSetBytes(0xFF))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}

	if !testVerify(words) || words != "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote"{
		t.Errorf("Error with vector 2, got %s", words)
	}
}