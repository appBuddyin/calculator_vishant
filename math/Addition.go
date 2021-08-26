package math

type Addition struct{}

func (d Addition) Op(v Variables) float32 { return v.var1 + v.var2 }
