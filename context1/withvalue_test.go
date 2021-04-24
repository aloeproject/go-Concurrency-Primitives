package context1

import "testing"

func Test_withValue(t *testing.T) {
	withValue()
}

func Test_withCancel(t *testing.T) {
	withCancel()
}

func Test_withCancel2(t *testing.T) {
	withCancel2()
}

func Test_withTimeout(t *testing.T) {
	withTimeout()
}

func Test_withDeadline(t *testing.T) {
	withDeadline()
}

func Test_childCancelCtx(t *testing.T) {
	childCancelCtx()
}

func Test_childCancelCtx2(t *testing.T) {
	childCancelCtx2()
}

func Test_withValueTransfer(t *testing.T) {
	withValueTransfer()
}