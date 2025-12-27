package assert

import "testing"

func TestAsserts(t *testing.T) {
	Lt(1,2,"huh?")
	Le(1,2,"huh?")
	Le(2,2,"huh?")
	Gt(3,2,"huh?")
	Ge(3,2,"huh?")
	Ge(3,3,"huh?")
	Eq(3,3,"huh?")
	Ne(3,4,"huh?")
	True(true,"huh?")
	False(false,"huh?")
}
