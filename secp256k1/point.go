package secp256k1

import(
	"errors"
	"math/big"
)

type CurvePoint struct{
	a *FieldElement
	b *FieldElement	
	x *FieldElement
	y *FieldElement
}


func(first *CurvePoint) GetX() *big.Int{
	if(first.x == nil){
		panic("x is nil")
	}
	return first.x.Number;
}

func(first *CurvePoint) GetY() *big.Int{
	if(first.y == nil){
		panic("y is nil")
	}
	return first.y.Number;
}

func NewPointCurve(x *FieldElement, y *FieldElement, a *FieldElement, b *FieldElement) (point * CurvePoint, err error){
	if(x == nil && y == nil){
		return &CurvePoint{a, b, x, y}, nil
	}

	if(y.Pow(2).NotEqual((x.Pow(3)).Add((a.Mul(*x))).Add(*b))){
		return nil, errors.New("invalid point")
	}

	return  &CurvePoint{a, b, x, y}, nil
}

func(first *CurvePoint) Equal_Real(second *CurvePoint) bool {
	if(first == nil && second != nil){
		return false;
	}
	if(first == nil && second == nil){
		return true;
	}

	if(first.a.Equal(*second.a)){
		if(first.b.Equal(*second.b)){
			if((first.x == nil || second.x == nil)){
				if(first.x == second.x){
					return true
				}
			}else{
				if(first.x.Equal(*second.x)){
					if(first.y.Equal(*second.y)){
						return true
					}
				}
			}
		}
	}
	return true;
}

func(first *CurvePoint) NotEqual_Real(second *CurvePoint) bool {
	return !first.Equal_Real(second)
}

func(first *CurvePoint) AddReal(second *CurvePoint) (output *CurvePoint, err error) {
	if(first == nil || second == nil){
		return nil, errors.New("input is nil")
	}

	if(first.a.NotEqual(*second.a)){
		return nil, nil
	}
	if(first.b.NotEqual(*second.b)){
		return nil, nil
	}

	if(first.x == nil){
		return second, nil
	}
	if(second.x == nil){
		return first, nil
	}

	if(first.x.Equal(*second.x) && first.y.NotEqual(*second.y)){
		return NewPointCurve(nil, nil, first.a, first.b)
	}
	

	if(first.x.NotEqual(*second.x)){
		s := (second.y.Sub(*first.y)).Div(second.x.Sub(*first.x))
		x := (s.Pow(2)).Sub(*first.x).Sub(*second.x)
		y := (s.Mul(first.x.Sub(x))).Sub(*first.y)

		return NewPointCurve(&x, &y, first.a, first.b)
	}

	if(first.Equal_Real(second) && (first.y.Equal(first.x.MulWithInt(0)))){
		return NewPointCurve(nil, nil, first.a, first.b)
	}

	if(first.Equal_Real(second)){
		s := (first.x.Pow(2).MulWithInt(3).Add(*first.a)).Div(first.y.MulWithInt(2))
		x := (s.Pow(2)).Sub(first.x.MulWithInt(2))
		y := (s.Mul(first.x.Sub(x))).Sub(*first.y)

		return NewPointCurve(&x, &y, first.a, first.b)
	}
	return nil, nil
}

func(input *CurvePoint) Mul(coefficient *big.Int) (output *CurvePoint, err error) {
	current := input
	
	if(input == nil){
		return nil, errors.New("input is nil")
	}

	result, err := NewPointCurve(nil, nil, input.a, input.b)

	for coefficient.Cmp(big.NewInt(0)) == 1{
		test := big.NewInt(0)
		test.And(coefficient, big.NewInt(1))

		if  test.Cmp(big.NewInt(1)) == 0{
			result, err = result.AddReal(current)
			if(err != nil){
				panic(err)
			}
		}

		current, err = current.AddReal(current)
		if(err != nil){
			panic(err)
		}
		coefficient.Rsh(coefficient, 1)
	}
	return result, nil
}

