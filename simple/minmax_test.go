package simple

import (
	"testing"
)

import (
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Min_Empty
func Test_Min_Empty(t *testing.T) {
	r := New(cmp.F_int)
	assertKVF(-1, -1, false, r.Min, t)
}

// Test_Min
func Test_Min(t *testing.T) {
	r := randomTree(1000)
	assertKVF(0, 0, true, r.Min, t)
}

// Test_Max_Empty
func Test_Max_Empty(t *testing.T) {
	r := New(cmp.F_int)
	assertKVF(-1, -1, false, r.Max, t)
}

// Test_Max
func Test_Max(t *testing.T) {
	r := randomTree(1000)
	assertKVF(999, 999, true, r.Max, t)
}
