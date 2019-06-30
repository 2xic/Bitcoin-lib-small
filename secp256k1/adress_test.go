package secp256k1

import (
	"testing"
	"bytes"
	"fmt"
)

func Test_CheckSum(t *testing.T) {
	check := CheckSum([]byte{0})
	if(bytes.Compare(check, []byte{0,20,6,224,88}) != 0){
		t.Errorf("Error with checksum")
	}

	check = ReturnCheckSum([]byte{0})
	if(bytes.Compare(check, []byte{20,6,224,88}) != 0){
		t.Errorf("Error with checksum")
	}

	check = DoubleSha256([]byte{0})
	if(bytes.Compare(check, []byte{20,6,224,88,129,226,153,54,119,102,211,19,226,108,5,86,78,201,27,247,33,211,23,38,189,110,70,230,6,137,83,154}) != 0){
		fmt.Println(check)
		t.Errorf("Error with checksum")
	}
}

