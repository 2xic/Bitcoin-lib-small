package base58

import (
	"testing"
	"strings"
)

func Test_Base58(t *testing.T) {
	if(!(strings.Compare(ConvertString("test"), "3yZe7d") == 0)){
		t.Errorf("Check base58 code")
	}
}

