package math

import "log"

type Divide struct{}

func (h Divide) Op(v Variables) float32 {
	validate(v)
	return v.var1 / v.var2
}

func validate(v Variables) {
	if v.var2 == 0 {
		log.Print("can't divide by zero")
		panic("can't divide by zero")
	}
}
