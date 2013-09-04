package simple

//import . "github.com/iNamik/go_pkg/debug/assert"
//import . "github.com/iNamik/go_pkg/debug/ping"

import (
	"math/rand"
	"testing"
)

import (
	"github.com/iNamik/go_bst/walker"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Walk_Empty
func Test_Walk_Empty(t *testing.T) {
	r := New(cmp.F_int)
	r.Walk(func(n walker.Node) walker.Action {
		t.Fatal("walk() called")
		return walker.RETURN
	})
}

// Test_Walk_Get
func Test_Walk_Get(t *testing.T) {
	const VALUE = 500
	r := randomTree(1000)
	var value int = -1
	var found bool = false
	r.Walk(func(n walker.Node) walker.Action {
		switch n.Cmp(VALUE, n.Key()) {
		case cmp.LT:
			if n.HasLeft() == true {
				return walker.LEFT
			}
			value, found = -1, false
			return walker.RETURN
		case cmp.GT:
			if n.HasRight() == true {
				return walker.RIGHT
			}
			value, found = -1, false
			return walker.RETURN
		default:
			value, found = n.Value().(int), true
			return walker.RETURN
		}
	})
	if found != true {
		t.Fatalf("get returned false")
	}
	if value != VALUE {
		t.Fatalf("get returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_Walk_Min
func Test_Walk_Min(t *testing.T) {
	const MIN = 0
	r := randomTree(1000)
	var min int = -1
	r.Walk(func(n walker.Node) walker.Action {
		if n.HasLeft() {
			return walker.LEFT
		}
		min = n.Value().(int)
		return walker.RETURN
	})
	if min != MIN {
		t.Fatalf("min returned '%d' instead of '%d'", min, MIN)
	}
}

// Test_Walk_Max
func Test_Walk_Max(t *testing.T) {
	const MAX = 999
	r := randomTree(1000)
	var max int = -1
	r.Walk(func(n walker.Node) walker.Action {
		if n.HasRight() {
			return walker.RIGHT
		}
		max = n.Value().(int)
		return walker.RETURN
	})
	if max != MAX {
		t.Fatalf("max returned '%d' instead of '%d'", max, MAX)
	}
}

// Test_Walk_Foreach_Min
func Test_Walk_Foreach_Min(t *testing.T) {
	const MIN = 0
	const MAX = 999
	r := randomTree(1000)
	var (
		min     int  = -1
		max     int  = -1
		i       int  = MIN
		haveMin bool = false
	)
	r.Walk(func(n walker.Node) walker.Action {
		if haveMin == false {
			if n.HasLeft() == true {
				return walker.LEFT
			}
			min = n.Value().(int)
			haveMin = true
		}
		v := n.Value().(int)
		if v != i {
			t.Fatalf("encountered '%d' instead of '%d", v, i)
		}
		max = v
		if n.HasNext() {
			i++
			return walker.NEXT
		}
		return walker.RETURN
	})
	if min != MIN {
		t.Fatalf("min returned '%d' instead of '%d'", min, MIN)
	}
	if max != MAX {
		t.Fatalf("max returned '%d' instead of '%d'", max, MAX)
	}
}

// Test_Walk_Foreach_Max
func Test_Walk_Foreach_Max(t *testing.T) {
	const MIN = 0
	const MAX = 999
	r := randomTree(1000)
	var (
		min     int  = -1
		max     int  = -1
		i       int  = MAX
		haveMax bool = false
	)
	r.Walk(func(n walker.Node) walker.Action {
		if haveMax == false {
			if n.HasRight() == true {
				return walker.RIGHT
			}
			max = n.Value().(int)
			haveMax = true
		}
		v := n.Value().(int)
		if v != i {
			t.Fatalf("encountered '%d' instead of '%d", v, i)
		}
		min = v
		if n.HasPrev() {
			i--
			return walker.PREV
		}
		return walker.RETURN
	})
	if min != MIN {
		t.Fatalf("min returned '%d' instead of '%d'", min, MIN)
	}
	if max != MAX {
		t.Fatalf("max returned '%d' instead of '%d'", max, MAX)
	}
}

// Test_Walk_Foreach_Min2 uses left/right/parent to implement foreachMin
func Test_Walk_Foreach_Min2(t *testing.T) {
	const MIN = 0
	const MAX = 999
	r := randomTree(1000)
	var (
		min     int  = -1
		max     int  = -1
		i       int  = MIN // Value we expect to see first
		haveMin bool = false
		stack   []int
	)
	r.Walk(func(n walker.Node) walker.Action {
		// Add space to stack if new level
		if len(stack) < n.Level() {
			stack = append(stack, 0) // 0 = not-visited
		}
		// Walk left
		if stack[n.Level()-1] == 0 {
			stack[n.Level()-1] = 1 // 1 == visited left
			if n.HasLeft() == true {
				return walker.LEFT
			}
		}
		// Visit / Walk right
		if stack[n.Level()-1] == 1 {
			// Visit value
			v := n.Value().(int)
			if haveMin == false {
				min = v
				haveMin = true
			}
			if v != i {
				t.Fatalf("encountered '%d' instead of '%d", v, i)
			}
			i++ // Value we expect to see next
			max = v
			// Walk right
			stack[n.Level()-1] = 2 // 2 == visited left and right
			if n.HasRight() == true {
				return walker.RIGHT
			}
		}
		// Pop stack and go up tree
		stack = stack[0 : len(stack)-1]
		if len(stack) == 0 {
			return walker.RETURN
		}
		return walker.PARENT
	})
	if min != MIN {
		t.Fatalf("min returned '%d' instead of '%d'", min, MIN)
	}
	if max != MAX {
		t.Fatalf("max returned '%d' instead of '%d'", max, MAX)
	}
}

// Test_Walk_Foreach_Max2  uses left/right/parent to implement foreachMax
func Test_Walk_Foreach_Max2(t *testing.T) {
	const MIN = 0
	const MAX = 999
	r := randomTree(1000)
	var (
		min     int  = -1
		max     int  = -1
		i       int  = MAX // Value we expect to see first
		haveMax bool = false
		stack   []int
	)
	r.Walk(func(n walker.Node) walker.Action {
		// Add space to stack if new level
		if len(stack) < n.Level() {
			stack = append(stack, 0) // 0 = not-visited
		}
		// Walk right
		if stack[n.Level()-1] == 0 {
			stack[n.Level()-1] = 1 // 1 == visited right
			if n.HasRight() == true {
				return walker.RIGHT
			}
		}
		// Visit / Walk left
		if stack[n.Level()-1] == 1 {
			// Visit value
			v := n.Value().(int)
			if haveMax == false {
				max = v
				haveMax = true
			}
			if v != i {
				t.Fatalf("encountered '%d' instead of '%d", v, i)
			}
			i-- // Value we expect to see next
			min = v
			// Walk left
			stack[n.Level()-1] = 2 // 2 == visited left and right
			if n.HasLeft() == true {
				return walker.LEFT
			}
		}
		// Pop stack and go up tree
		stack = stack[0 : len(stack)-1]
		if len(stack) == 0 && n.HasParent() == false {
			return walker.RETURN
		}
		return walker.PARENT
	})
	if min != MIN {
		t.Fatalf("min returned '%d' instead of '%d'", min, MIN)
	}
	if max != MAX {
		t.Fatalf("max returned '%d' instead of '%d'", max, MAX)
	}
}

// Test_Walk_Random
func Test_Walk_Random(t *testing.T) {
	const MIN = 0
	const MAX = 999
	r := randomTree(MAX + 1)
	var (
		count     = 1000000
		dir   int = 0
		i     int = 0
		check int = -1
	)
	r.Walk(func(n walker.Node) walker.Action {
		if count == 0 {
			return walker.RETURN
		}
		count--
		v := n.Value().(int)
		if check >= 0 && v != check {
			t.Fatalf("Expecting '%d' but found '%d'", check, v)
		}
		if i == 0 {
			if n.HasPrev() == false {
				dir = 1
			} else if n.HasNext() == false {
				dir = 0
			} else {
				dir = rand.Intn(2)
			}
			switch dir {
			case 0:
				i = rand.Intn(v) + 1
			case 1:
				i = rand.Intn(MAX-v) + 1
			default:
				panic("unreachable")
			}
			check = v
		}
		i--
		switch dir {
		case 0: // 0 = prev
			check--
			return walker.PREV
		case 1: // 1 = next
			check++
			return walker.NEXT
		default:
			panic("unreachable")
		}
	})
}

// Test_Walk_Exception_Left
func Test_Walk_Exception_Left(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "cannot walk left when hasLeft() == false", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.LEFT
		})
	})
}

// Test_Walk_Exception_Right
func Test_Walk_Exception_Right(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "cannot walk right when hasRight() == false", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.RIGHT
		})
	})
}

// Test_Walk_Exception_Prev
func Test_Walk_Exception_Prev(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "cannot walk prev when hasPrev() == false", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.PREV
		})
	})
}

// Test_Walk_Exception_Next
func Test_Walk_Exception_Next(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "cannot walk next when hasNext() == false", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.NEXT
		})
	})
}

// Test_Walk_Exception_Parent
func Test_Walk_Exception_Parent(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "cannot walk parent when hasParent() == false", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.PARENT
		})
	})
}

// Test_Walk_Exception_Illegal
func Test_Walk_Exception_Illegal(t *testing.T) {
	r := New(cmp.F_int)
	r.ReplaceOrInsert(key1, key1)
	assertPanic(t, "illegal walk action 'walker.Action(-1)'", func() {
		r.Walk(func(n walker.Node) walker.Action {
			return walker.Action(-1)
		})
	})
}
