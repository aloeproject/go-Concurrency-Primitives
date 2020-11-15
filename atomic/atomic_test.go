package atomic

import "testing"

func TestAdd(t *testing.T) {
	Add()
}

func Test_Load(t *testing.T) {
	Load()
}

func Test_cas(t *testing.T) {
	CompareAndSwap()
}