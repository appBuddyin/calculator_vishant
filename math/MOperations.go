package math

type MOperation interface {
	Op(v Variables) float32
}

func GetOperation(op string) MOperation {
	switch op {
	case "*":
		return Multiplication{}
	case "+":
		return Addition{}
	case "-":
		return Subtract{}
	case "\\":
		return Divide{}
	}
	panic("The operation is permitted")
}
