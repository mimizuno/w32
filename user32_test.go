package w32

import (
	"testing"
	"unsafe"
)

func Test_cINPUT_size(t *testing.T) {
	sizeHardwareInput := unsafe.Sizeof(HardwareInput{})
	sizeMouseInput := unsafe.Sizeof(MouseInput{})
	sizeKbdInput := unsafe.Sizeof(KbdInput{})
	sizeCInput := unsafe.Sizeof(cINPUT{})
	sizeDWORD := unsafe.Sizeof(DWORD(0))

	max := sizeHardwareInput
	if sizeMouseInput > max {
		max = sizeMouseInput
	}
	if sizeKbdInput > max {
		max = sizeKbdInput
	}

	expected := max + sizeDWORD
	if expected != sizeCInput {
		t.Errorf("Size not matched. expect: %v, actual: %v", expected, sizeCInput)
	}
}
