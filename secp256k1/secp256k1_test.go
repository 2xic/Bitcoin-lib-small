package secp256k1

import (
	"testing"
	"math/big"
	"github.com/2xic/bip-39/base58"
	"strings"
	"crypto/hmac"
	"crypto/sha512"
)

func Test_Point(t *testing.T) {
	prime := int64(223)

	a := NewFromInt(0, prime)
	b := NewFromInt(7, prime)

	valid := [][]uint8{
		{192, 105},
		{17, 56},
		{1, 193},
	}

	invalid := [][]uint8{
		{200, 119},
		{42, 99},
	}

	for i := 0; i < len(valid); i++{
		x := NewFromInt(int64(valid[i][0]), prime)
		y := NewFromInt(int64(valid[i][1]), prime)

		_, err := NewPointCurve(&x, &y, &a, &b)
		if(err != nil){
			panic(err)
		}
	}

	for i := 0; i < len(invalid); i++{
		x := NewFromInt(int64(invalid[i][0]), prime)
		y := NewFromInt(int64(invalid[i][1]), prime)

		_, err := NewPointCurve(&x, &y, &a, &b)
		if(err == nil){
			panic(err)
		}
	}
}

func Test_Add(t *testing.T) {
	prime := int64(223)

	a := NewFromInt(0, prime)
	b := NewFromInt(7, prime)

	valid := [][]uint8{
		{192, 105, 17, 56, 170, 142},
		{47, 71, 117, 141, 60, 139},
		{143, 98, 76, 66, 47, 71},

	}

	for i := 0; i < len(valid); i++{
		x1 := NewFromInt(int64(valid[i][0]), prime)
		y1 := NewFromInt(int64(valid[i][1]), prime)

		point1, err := NewPointCurve(&x1, &y1, &a, &b)

		if(err != nil){
			panic(err)
		}

		x2 := NewFromInt(int64(valid[i][2]), prime)
		y2 := NewFromInt(int64(valid[i][3]), prime)

		point2, err := NewPointCurve(&x2, &y2, &a, &b)
		if(err != nil){
			panic(err)
		}

		x3 := NewFromInt(int64(valid[i][4]), prime)
		y3 := NewFromInt(int64(valid[i][5]), prime)

		found, err := NewPointCurve(&x3, &y3, &a, &b)
		if(err != nil){
			panic(err)
		}

		target, err := point1.AddReal(point2)

		if(err != nil){
			panic(err)
		}

		if((target).NotEqual_Real(found)){
			panic(i)
		}
	}
}

func Test_RMul(t *testing.T) {
	prime := int64(223)

	a := NewFromInt(0, prime)
	b := NewFromInt(7, prime)

	valid := [][]uint8{
		{2, 192, 105, 49, 71},
		{2, 143, 98, 64, 168},
		{2, 47, 71, 36, 111},
		{4, 47, 71, 194, 51},
		{8, 47, 71, 116, 55},
		{21, 47, 71, 0, 0},
	}

	for i := 0; i < len(valid); i++{
		s := int64(valid[i][0])
		x1 := NewFromInt(int64(valid[i][1]), prime)
		y1 := NewFromInt(int64(valid[i][2]), prime)

		test1, err := NewPointCurve(&x1, &y1, &a, &b)
		if(err != nil){
			panic(err)
		}

		target, err := test1.Mul(big.NewInt(s))
		
		if(valid[i][3] == 0){
			found, err := NewPointCurve(nil, nil, &a, &b)

			if(err != nil){
				panic(err)
			}

			if(target.NotEqual_Real(found)){
				panic("error with rmul")
			}
		}else{
			x2 := NewFromInt(int64(valid[i][3]), prime)
			y2 := NewFromInt(int64(valid[i][4]), prime)
			found, err := NewPointCurve(&x2, &y2, &a, &b)

			if(err != nil){
				panic(err)
			}
	
			if(target.NotEqual_Real(found)){
				panic("error with rmul")
			}
		}
	}
}


func Test_private_to_public(t *testing.T) {
	valid := [][]string{
		{"7", "5cbdf0646e5db4eaa398f365f2ea7a0e3d419b7e0330e39ce92bddedcac4f9bc", "6aebca40ba255960a3178d6d861a54dba813d0b813fde7b5a5082628087264da"},
		{"5cd", "c982196a7466fbbbb0e27a940b6af926c1a74d5ad07128c82824a11b5398afda", "7a91f9eae64438afb9ce6448a1c133db2d8fb9254e4546b6f001637d50901f55"},
		{"100000000000000000000000000000000", "8f68b9d2f63b5f339239c1ad981f162ee88c5678723ea3351b7b444c9ec4c0da", "662a9f2dba063986de1d90c2b6be215dbbea2cfe95510bfdf23cbf79501fff82"},
		{"1000000000000000000000000000000000000000000000000000080000000", "9577ff57c8234558f293df502ca4f09cbc65a6572c842b39b366f21717945116", "10b49c67fa9365ad7b90dab070be339a1daf9052373ec30ffae4f72d5e66d053"},
	}

	for i:= 0; i < len(valid); i++{
		secret, _ :=  new(big.Int).SetString(valid[i][0], 16)

		x, _ := new(big.Int).SetString(valid[i][1], 16)
		y, _ := new(big.Int).SetString(valid[i][2], 16)
		point, err := S256Point(x, y)
		if(err != nil){
			panic(err)
		}

		public, err := G.Mul(secret)
		if(err != nil){
			panic(err)
		}

		if((public).NotEqual_Real(point)){
			panic("error with private -> public key")
		}
	}
}


func Test_valid_private(t *testing.T){
	hmac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	_, err := hmac.Write([]byte{0})
	if err != nil {
		t.Errorf("Error with hmac")
	}else{
		hash := hmac.Sum(nil)
		if(!CheckPrivateKey(hash[:32])){
			t.Errorf("Error with private key validation")
		}
	}
}

func Test_compress(t *testing.T){
	coefficient := new(big.Int).Exp(big.NewInt(999), big.NewInt(3), nil)
	point := GetPublicPoint(coefficient)

	test := CompressPublicKey(point)
	x, y := ExpandPublicKey(test)

	if(!(point.GetY().Cmp(y) == 0 && point.GetX().Cmp(x) == 0)){
		t.Errorf("Error with public key compression")
	}

	coefficient = new(big.Int).Exp(big.NewInt(888), big.NewInt(3), nil)
	point = GetPublicPoint(coefficient)
	
	results := CompressAddress(point)
	if(!(strings.Compare(base58.ConvertBytes(results), "148dY81A9BmdpMhvYEVznrM45kWN32vSCN") == 0)){
		t.Errorf("Error with address compression")
	}
}

