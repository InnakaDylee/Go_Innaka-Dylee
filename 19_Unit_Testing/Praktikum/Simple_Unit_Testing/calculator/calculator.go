package calculator

func Addition(numA, numB float32) float32{
	return numA+numB
}

func Substraction(numA, numB float32) float32{
	return numA-numB
}

func Division(numA, numB float32) float32{
	if numB <= 0{
		return 0
	}
	return numA/numB
}

func Multiple(numA, numB float32) float32 {
	return numA*numB
}