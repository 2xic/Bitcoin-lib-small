package secp256k1

import(
	"math/big"
)

type FieldElement struct{
	Number *big.Int
	Prime *big.Int
}

func NewFromInt(number int64, prime int64) FieldElement{
	return FieldElement{big.NewInt(number), 
						big.NewInt(prime)}
}

func New(number *big.Int, prime *big.Int) FieldElement{
	if((prime.Cmp(number) == -1 || prime.Cmp(number) == 0) || number.Cmp(big.NewInt(0)) == -1){
		panic("Bad input for field.")
	}
	return FieldElement{number, prime}
}

func(first FieldElement) Equal(second FieldElement) bool {
	if(first.Number.Cmp(second.Number) == 0){
		if(first.Prime.Cmp(second.Prime) == 0){
			return true
		}
	}
	return false
}

func(first FieldElement) NotEqual(second FieldElement) bool {
	return !first.Equal(second)
}

func(first FieldElement) Add(second FieldElement) FieldElement {
	if(first.Prime.Cmp(second.Prime) != 0){
		panic("Wrong prime combo.")
	}
	newNumber := big.NewInt(0)
	newNumber.Mod(newNumber.Add(first.Number, second.Number), first.Prime)
	return New(newNumber, first.Prime)
}

func(first FieldElement) Sub(second FieldElement) FieldElement {
	if(first.Prime.Cmp(second.Prime) != 0){
		panic("Wrong prime combo.")
	}
	newNumber := big.NewInt(0)
	newNumber.Mod(newNumber.Sub(first.Number, second.Number), first.Prime)
	return New(newNumber, first.Prime)
}

func(first FieldElement) Mul(second FieldElement) FieldElement {
	if(first.Prime.Cmp(second.Prime) != 0){
		panic("Wrong prime combo.")
	}
	newNumber := big.NewInt(0)
	newNumber.Mod(newNumber.Mul(first.Number, second.Number), first.Prime)
	return New(newNumber, first.Prime)
}

func(first FieldElement) MulWithInt(coefficient int64) FieldElement {
	newNumber := big.NewInt(0)
	newNumber.Mod(newNumber.Mul(first.Number, big.NewInt(coefficient)), first.Prime)
	return New(newNumber, first.Prime)
}

func(first FieldElement) Div(second FieldElement) FieldElement {
	if(first.Prime.Cmp(second.Prime) != 0){
		panic("Wrong prime combo.")
	}

	newNumber := big.NewInt(0)
	temp1 := big.NewInt(0)
	temp1.Sub(first.Prime, big.NewInt(2))
	newNumber.Mod((newNumber.Mul(first.Number, newNumber.Exp(second.Number, temp1, first.Prime))), first.Prime)
	return New(newNumber, first.Prime)
}

func(first FieldElement) Pow(exponent int) FieldElement {
	power := big.NewInt(0)
	mod := big.NewInt(0)
	mod.Sub(first.Prime, big.NewInt(1))
	power.Mod(big.NewInt(int64(exponent)), mod)

	newNumber := big.NewInt(0)
	newNumber.Exp(first.Number, power, first.Prime)
	return  New(newNumber, first.Prime)
}
