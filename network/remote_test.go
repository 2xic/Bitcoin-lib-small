package network

import (
	"testing"
)

func Test(t *testing.T) {
	sucess := false
	for i:= 0; i < 10; i++{
		data := Getblock()
		if(data != nil){
			sucess = true
			break
		}
	}

	if(!sucess){
		t.Errorf("get block failed")
	}
}