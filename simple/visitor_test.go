package simple

import (
	"math/rand"
	"testing"
)

import (
	"github.com/iNamik/go_bst/visitor"
)

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

func assertVisit(r T, key int, value interface{}, result visitor.Result, t *testing.T, f visitor.F) {
	v_, result_ := r.Visit(key, f)
	if result_ != result {
		t.Fatalf("visit() returned result '%s' instead of '%s'", result_, result)
	}
	if result != visitor.NOT_FOUND && result != visitor.REMOVED {
		if v_ != value {
			t.Fatalf("visit() returned value '%v' instead of '%v'", v_, value)
		}
	}
}

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Visit_Found_Get
func Test_Visit_Found_Get(t *testing.T) {
	r := randomTree(1000)
	assertVisit(r, 500, 500, visitor.FOUND, t, func(v interface{}, _ bool) (interface{}, visitor.Action) {
		return v, visitor.GET
	})
	assertGet(r, 500, 500, true, t)
}

// Test_Visit_Found_Replace
func Test_Visit_Found_Replace(t *testing.T) {
	r := randomTree(1000)
	assertVisit(r, 500, 400, visitor.REPLACED, t, func(v interface{}, _ bool) (interface{}, visitor.Action) {
		return 400, visitor.REPLACE
	})
	assertGet(r, 500, 400, true, t)
}

// Test_Visit_Found_Remove
func Test_Visit_Found_Remove(t *testing.T) {
	r := randomTree(1000)
	assertVisit(r, 500, nil, visitor.REMOVED, t, func(v interface{}, _ bool) (interface{}, visitor.Action) {
		return 400, visitor.REMOVE
	})
	assertGet(r, 500, nil, false, t)
}

// Test_Visit_Found_Insert
func Test_Visit_Found_Insert(t *testing.T) {
	r := randomTree(1000)
	assertPanic(t, "illegal action 'INSERT' when visiting found key", func() {
		r.Visit(500, func(v interface{}, _ bool) (interface{}, visitor.Action) {
			return v, visitor.INSERT // Can't insert a found key, should panic
		})
	})
	assertGet(r, 500, 500, true, t)
}

// Test_Visit_NotFound_Insert
func Test_Visit_NotFound_Insert(t *testing.T) {
	r := randomTree(1000)
	assertVisit(r, 1000, 1000, visitor.INSERTED, t, func(v interface{}, _ bool) (interface{}, visitor.Action) {
		return 1000, visitor.INSERT
	})
	assertGet(r, 1000, 1000, true, t)
}

// Test_Visit_NotFound_Get
func Test_Visit_NotFound_Get(t *testing.T) {
	r := randomTree(1000)
	assertVisit(r, 1000, nil, visitor.NOT_FOUND, t, func(v interface{}, _ bool) (interface{}, visitor.Action) {
		return 1000, visitor.GET
	})
	assertGet(r, 1000, nil, false, t)
}

// Test_Visit_NotFound_Replace
func Test_Visit_NotFound_Replace(t *testing.T) {
	r := randomTree(1000)
	assertPanic(t, "illegal action 'REPLACE' when visiting non-found key", func() {
		r.Visit(1000, func(v interface{}, _ bool) (interface{}, visitor.Action) {
			return v, visitor.REPLACE // Can't update a non-found key, should panic
		})
	})
	assertGet(r, 1000, nil, false, t)
}

// Test_Visit_NotFound_Remove
func Test_Visit_NotFound_Remove(t *testing.T) {
	r := randomTree(1000)
	assertPanic(t, "illegal action 'REMOVE' when visiting non-found key", func() {
		r.Visit(1000, func(v interface{}, _ bool) (interface{}, visitor.Action) {
			return v, visitor.REMOVE // Can't remove a non-found key, should panic
		})
	})
	assertGet(r, 1000, nil, false, t)
}

// Test_Visit_Random
func Test_Visit_Random(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	const SIZE = 100
	const ITERATIONS = 1000000
	r := randomTree(SIZE)
	for i := 0; i < ITERATIONS; i++ {
		n := rand.Intn(SIZE)
		r.Visit(n, func(v_ interface{}, found bool) (interface{}, visitor.Action) {
			if found {
				return nil, visitor.REMOVE
			} else {
				return n, visitor.INSERT
			}
		})
		assertBST(r, t)
	}
}
