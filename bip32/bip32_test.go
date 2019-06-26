package bip32

import(
	"math/big"
	"testing"
	"strings"
)

func Test_addr(t *testing.T) {
	seed := []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f}
	test2 := new(big.Int)
	test2.SetBytes(seed)

	MasterPrivateKey := MasterPrivateKey(seed)
	MasterPublicKey := MasterPublicKey(MasterPrivateKey)

	if(strings.Compare(MasterPrivateKey.Readable(), "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi")) != 0{
		panic("wrong master private key")
	}
	
	if(strings.Compare(MasterPublicKey.Readable(), "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8")) != 0{
		panic("wrong master public key")
	}

	if(strings.Compare(MasterPrivateKey.ReadableAddress(), "15mKKb2eos1hWa6tisdPwwDC1a5J1y9nma")) != 0{
		panic("wrong address")
	}

}

