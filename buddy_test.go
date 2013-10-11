package buddy

import (
	"testing"
)

func TestBuddyNew(t *testing.T) {
	if b, e := New(-1); b != nil {
		t.Errorf("expect nil, but got %v\n", b)
	} else if e != InValidParameterErr {
		t.Errorf("expect error %v, but got %v\n", InValidParameterErr, e)
	}
	b, e := New(3)
	if e != nil {
		t.Fatalf("expect error nil, but got %v\n", e)
	}
	expect := &Buddy{4, []int{4, 2, 2, 1, 1, 1, 1}}
	if !isEqual(b, expect) {
		t.Errorf("expect buddy %v, but got %v\n", expect, b)
	}
}

func isEqual(a, b *Buddy) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		return false
	}
	if a.size != b.size {
		return false
	}
	if len(a.longests) != len(b.longests) {
		return false
	}
	for i, av := range a.longests {
		if av != b.longests[i] {
			return false
		}
	}
	return true
}

func TestBuddyAllocAndFree(t *testing.T) {
	b, _ := New(1024)
	if o, e := b.Alloc(0); e != InValidParameterErr {
		t.Errorf("expect error %v, but got %v\n", InValidParameterErr, e)
	} else if o != 0 {
		t.Errorf("expect offset 0, but got %v\n", o)
	}

	if o, e := b.Alloc(70); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	} else if o != 0 {
		t.Errorf("expect offset 0, but got %v\n", o)
	}

	if o, e := b.Alloc(35); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	} else if o != 128 {
		t.Errorf("expect offset 128, but got %v\n", o)
	}

	if o, e := b.Alloc(80); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	} else if o != 256 {
		t.Errorf("expect offset 256, but got %v\n", o)
	}

	if o, e := b.Alloc(550); e != NotFoundErr {
		t.Errorf("expect error %v, but got %v\n", NotFoundErr, e)
	} else if o != 0 {
		t.Errorf("expect offset 0, but got %v\n", o)
	}

	if e := b.Free(-1); e != InValidParameterErr {
		t.Errorf("expect error %v, but got %v\n", InValidParameterErr, e)
	}
	if e := b.Free(1025); e != InValidParameterErr {
		t.Errorf("expect error %v, but got %v\n", InValidParameterErr, e)
	}

	if e := b.Free(0); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	}

	if o, e := b.Alloc(60); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	} else if o != 0 {
		t.Errorf("expect offset 0, but got %v\n", o)
	}

	if e := b.Free(0); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	}

	if e := b.Free(128); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	}

	if e := b.Free(256); e != nil {
		t.Errorf("expect error nil, but got %v\n", e)
	}

	if e := b.Free(0); e != NotFoundErr {
		t.Errorf("expect error %v, but got %v\n", NotFoundErr, e)
	}
}
