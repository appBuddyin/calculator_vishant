package math

type Multiplication struct{}

func (p Multiplication) Op(v Variables) float32 { return v.var1 * v.var2 }
