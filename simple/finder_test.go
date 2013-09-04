package simple

import (
	"testing"
)

import (
	"github.com/iNamik/go_bst/finder"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

// assertFind
func assertFind(r T, key int, value_ interface{}, found bool, t *testing.T, f finder.F) {
	k_, v_, found_ := r.Find(f)
	if found_ != found {
		t.Fatalf("find() returned %v", found_)
	}
	if found == true {
		k, ok := k_.(int)
		if ok == false {
			t.Fatal("find() did not return key of type int")
		}
		if k != key {
			t.Fatalf("find() returned key '%d' instead of '%d'", k, key)
		}
		if v_ != value_ {
			t.Fatalf("find() returned value '%v' instead of '%v'", v_, value_)
		}
	}
}

/**********************************************************************
 ** Helper Functions
 **********************************************************************/

func findLowerBound(r T, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	r.Find(func(node finder.Node) finder.Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.LT: // node < boundKey - We have a candidate
			// If node > working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.GT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Keep looking
			return finder.RIGHT
		case cmp.GT: // node > boundKey
			// Keep looking
			return finder.LEFT
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return finder.FOUND
		}
	})
	return
}

func findUpperBound(r T, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	r.Find(func(node finder.Node) finder.Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.GT: // node > boundKey - We have a candidate
			// If node < working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.LT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Keep looking
			return finder.LEFT
		case cmp.LT: // boundKey < node > boundKey
			// Keep looking
			return finder.RIGHT
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return finder.FOUND
		}
	})
	return
}

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Find_Empty
func Test_Find_Empty(t *testing.T) {
	assertFind(New(cmp.F_int), -1, -1, false, t, func(node finder.Node) finder.Action {
		t.Fatal("find() called on empty tree")
		return finder.NOT_FOUND
	})
}

// Test_Find_Min
func Test_Find_Min(t *testing.T) {
	assertFind(randomTree(1000), 0, 0, true, t, func(node finder.Node) finder.Action {
		if node.HasLeft() {
			return finder.LEFT
		}
		return finder.FOUND
	})
}

// Test_Find_Max
func Test_Find_Max(t *testing.T) {
	assertFind(randomTree(1000), 999, 999, true, t, func(node finder.Node) finder.Action {
		if node.HasRight() {
			return finder.RIGHT
		}
		return finder.FOUND
	})
}

// Test_Find_LowerBound_Found
func Test_Find_LowerBound_Found(t *testing.T) {
	boundKey := 1001
	r := randomTreeDouble(1000)
	f := func() (interface{}, interface{}, bool) { return findLowerBound(r, boundKey) }
	assertKVF(1000, 1000, true, f, t)
}

// Test_Find_LowerBound_NotFound
func Test_Find_LowerBound_NotFound(t *testing.T) {
	boundKey := -1
	r := randomTreeDouble(1000)
	f := func() (interface{}, interface{}, bool) { return findLowerBound(r, boundKey) }
	assertKVF(-1, -1, false, f, t)
}

// Test_Find_UpperBound_Found
func Test_Find_UpperBound_Found(t *testing.T) {
	boundKey := 1001
	r := randomTreeDouble(1000)
	f := func() (interface{}, interface{}, bool) { return findUpperBound(r, boundKey) }
	assertKVF(1002, 1002, true, f, t)
}

// Test_Find_UpperBound_NotFound
func Test_Find_UpperBound_NotFound(t *testing.T) {
	boundKey := 2001
	r := randomTreeDouble(1000)
	f := func() (interface{}, interface{}, bool) { return findUpperBound(r, boundKey) }
	assertKVF(-1, -1, false, f, t)
}

// Test_Find_Left_Nil
func Test_Find_Left_Nil(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertFind(r, -1, -1, false, t, func(node finder.Node) finder.Action {
		return finder.LEFT
	})
}

// Test_Find_Left_NotFound
func Test_Find_Left_NotFound(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key2, key2)
	r.ReplaceOrInsert(key1, key1)
	assertFind(r, -1, -1, false, t, func(node finder.Node) finder.Action {
		if node.HasLeft() {
			return finder.LEFT
		}
		return finder.NOT_FOUND
	})
}

// Test_Find_Right_Nil
func Test_Find_Right_Nil(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertFind(r, -1, -1, false, t, func(node finder.Node) finder.Action {
		return finder.RIGHT
	})
}

// Test_Find_Right_NotFound
func Test_Find_Right_NotFound(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	r.ReplaceOrInsert(key2, key2)
	assertFind(r, -1, -1, false, t, func(node finder.Node) finder.Action {
		if node.HasRight() {
			return finder.RIGHT
		}
		return finder.NOT_FOUND
	})
}

// Test_Find_Panic_Illegal
func Test_Find_Panic_Illegal(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "illegal find action 'finder.Action(-1)'", func() {
		r.Find(func(node finder.Node) finder.Action {
			return finder.Action(-1)
		})
	})
}
