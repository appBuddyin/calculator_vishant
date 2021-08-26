package math

type Subtract struct{}

func (o Subtract) Op(v Variables) float32 { return v.var1 - v.var2 }
