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


	check = Hash160([]byte{0})
	if(bytes.Compare(check, []byte{159, 127, 208, 150, 211, 126, 210, 192, 227, 247, 240, 207, 201, 36, 190, 239, 79, 252, 235, 104}) != 0){
		fmt.Println(check)
		t.Errorf("Error with checksum")
	}

	check = FingerPrint160([]byte{0})
	if(bytes.Compare(check, []byte{159, 127, 208, 150}) != 0){
		fmt.Println(check)
		t.Errorf("Error with checksum")
	}
	
	test := []byte{1, 2, 3, 0, 0}
	results := []byte{1, 2, 3, 4, 5}
	found := ExtendSlice(test, []byte{4, 5}, 3)

	if(bytes.Compare(results, found) != 0){
		t.Errorf("Error with extender")		
	}

	if(!allZeroBytes([]byte{0, 0, 0, 0})){
		t.Errorf("Error with extender")
	}

	if(allZeroBytes([]byte{0, 1, 0, 0})){
		t.Errorf("Error with extender")
	}

}

