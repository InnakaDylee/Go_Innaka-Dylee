package calculator

import "testing"

func TestAddition(t *testing.T) {
	if Addition(3,3) != 6 {
		t.Error("Wrong Answer")
	}
}
func TestSubstraction(t *testing.T) {
	if Substraction(3,2) != 1 {
		t.Error("Wrong Answer")
	}
}
func TestDivitionCannotZero(t *testing.T) {
	if Division(3,0) != 0 {
		t.Error("Wrong Answer")
	}
}
func TestDivition(t *testing.T) {
	if Division(3,3) != 1 {
		t.Error("Wrong Answer")
	}
}
func TestMultiple(t *testing.T) {
	if Multiple(3,3) != 9 {
		t.Error("Wrong Answer")
	}
}