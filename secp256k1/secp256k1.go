package secp256k1

import (
	"math/big"
)

var(
	A = big.NewInt(0)
	B = big.NewInt(7)
	N,_ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	P,_ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
)

func S256Field(seed *big.Int) (output *FieldElement, err error) {
	a := New(seed, P)

	return &a, nil
}

func S256Point(x *big.Int, y *big.Int) (output *CurvePoint, err error) {
	a,_ := S256Field(A)
	b,_ := S256Field(B)

	x2, _ := S256Field(x)
	y2, _ := S256Field(y)

	return NewPointCurve(x2, y2, a, b)
}

func GetPublicPoint(secret* big.Int) *CurvePoint{
	point, err := G.Mul(secret)
	if(err != nil){
		return nil
	}
	return point
}

var (
	G_X, _ = new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	G_Y, _ = new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	G, Err_g = S256Point(G_X, G_Y)
)
