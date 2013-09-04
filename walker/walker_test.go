package walker_test

//import . "github.com/iNamik/go_pkg/debug/assert"
//import . "github.com/iNamik/go_pkg/debug/ping"

import (
	"fmt"
	"testing"
)

import (
	"github.com/iNamik/go_bst/walker"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Types
 **********************************************************************/

// f_walk allows us to attach walker.I onto a closure
type f_walk func(walker.F)

// f_walk::Walk
func (w f_walk) Walk(f walker.F) { w(f) }

// this is to get vertical-alignment in my data :)
const (
	TRUE_ = true
	FALSE = false
)

/**********************************************************************
 ** node
 **********************************************************************/

type node struct {
	key       interface{}
	value     interface{}
	level     int
	hasNext   bool
	hasPrev   bool
	hasLeft   bool
	hasRight  bool
	hasParent bool
	action    walker.Action
}

// node::Key
func (n *node) Key() interface{} {
	return n.key
}

// node::Value
func (n *node) Value() interface{} {
	return n.value
}

// node::Cmp
func (n *node) Cmp(a interface{}, b interface{}) int {
	return cmp.Compare_int(a.(int), b.(int))
}

// node::Level
func (n *node) Level() int {
	return n.level
}

// node::HasPrev
func (n *node) HasPrev() bool {
	return n.hasPrev
}

// node::HasNext
func (n *node) HasNext() bool {
	return n.hasNext
}

// node::HasLeft
func (n *node) HasLeft() bool {
	return n.hasLeft
}

// node::HasRight
func (n *node) HasRight() bool {
	return n.hasRight
}

// node::HasParent
func (n *node) HasParent() bool {
	return n.hasParent
}

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

// assertWalk
func assertWalk(f walker.F, nodes []node, t *testing.T) {
	for i, n := range nodes {
		action := f(&n)
		if action != n.action {
			t.Fatalf("node[%d], walk returned '%s' instead of '%s'", i, action, n.action)
		}
	}
}

// assertActionString
func assertActionString(action walker.Action, value string, t *testing.T) {
	action_string := fmt.Sprint(action)
	if action_string != value {
		t.Fatalf("action code '%d' has String() value '%s' instead of '%s'", action, action_string, value)
	}
}

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Get
func Test_Get(t *testing.T) {
	const VALUE = 1
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	value_, found := w.Get(VALUE)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_Get_Miss_Left
func Test_Get_Miss_Left(t *testing.T) {
	const VALUE = 1
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, found := w.Get(VALUE)
	if found == true {
		t.Fatalf("walk returned true")
	}
}

// Test_Get_Miss_Right
func Test_Get_Miss_Right(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 1, value: 1, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, found := w.Get(VALUE)
	if found == true {
		t.Fatalf("walk returned true")
	}
}

// Test_Min
func Test_Min(t *testing.T) {
	const VALUE = 0
	nodes := []node{
		{key: 3, value: 3, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 0, value: 0, level: 4, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.Min()
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_Max
func Test_Max(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 1, value: 1, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 2, value: 2, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 3, value: 3, level: 4, hasPrev: TRUE_, hasNext: FALSE, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.Max()
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_ForeachMin
func Test_ForeachMin(t *testing.T) {
	const VALUE = 4
	nodes := []node{
		{key: 3, value: 3, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 0, value: 0, level: 4, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.NEXT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.NEXT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.NEXT},
		{key: 3, value: 3, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	i := 0
	w.ForeachMin(func(_ interface{}, value_ interface{}) {
		value := value_.(int)
		if value != i {
			t.Fatalf("walk returned '%d' instead of '%d'", value, i)
		}
		i++
	})
	if i != VALUE {
		t.Fatalf("i = '%d' instead of '%d'", i, VALUE)
	}
}

// Test_ForeachMax
func Test_ForeachMax(t *testing.T) {
	const VALUE = -1
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 1, value: 1, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 2, value: 2, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 3, value: 3, level: 4, hasPrev: TRUE_, hasNext: FALSE, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.PREV},
		{key: 2, value: 2, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.PREV},
		{key: 1, value: 1, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.PREV},
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	i := 3
	w.ForeachMax(func(_ interface{}, value_ interface{}) {
		value := value_.(int)
		if value != i {
			t.Fatalf("walk returned '%d' instead of '%d'", value, i)
		}
		i--
	})
	if i != VALUE {
		t.Fatalf("i = '%d' instead of '%d'", i, VALUE)
	}
}

// Test_LowerBound_Match
func Test_LowerBound_Match(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 3, value: 3, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.LowerBound(3)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Near_Right
func Test_LowerBound_Near_Right(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 4, value: 4, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 2, value: 2, level: 4, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.LowerBound(3)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Near_Left
func Test_LowerBound_Near_Left(t *testing.T) {
	const VALUE = 5
	nodes := []node{
		{key: 5, value: 5, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 8, value: 8, level: 2, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 7, value: 7, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.LowerBound(6)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Miss
func Test_LowerBound_Miss(t *testing.T) {
	nodes := []node{
		{key: 3, value: 3, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 2, value: 2, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 1, value: 1, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 0, value: 0, level: 4, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, _, found := w.LowerBound(-1)
	if found == true {
		t.Fatalf("walk returned true")
	}
}

// Test_UpperBound_Match
func Test_UpperBound_Match(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 5, value: 5, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 3, value: 3, level: 2, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.UpperBound(3)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Near_Right
func Test_UpperBound_Near_Right(t *testing.T) {
	const VALUE = 4
	nodes := []node{
		{key: 5, value: 5, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 1, value: 1, level: 2, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 4, value: 4, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: TRUE_, hasRight: FALSE, hasParent: TRUE_, action: walker.LEFT},
		{key: 2, value: 2, level: 4, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.UpperBound(3)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Near_Left
func Test_UpperBound_Near_Left(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 4, value: 4, level: 1, hasPrev: TRUE_, hasNext: FALSE, hasLeft: TRUE_, hasRight: FALSE, hasParent: FALSE, action: walker.LEFT},
		{key: 0, value: 0, level: 2, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 2, value: 2, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, value_, found := w.UpperBound(1)
	if found == false {
		t.Fatalf("walk returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("walk returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Miss
func Test_UpperBound_Miss(t *testing.T) {
	nodes := []node{
		{key: 0, value: 0, level: 1, hasPrev: FALSE, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: FALSE, action: walker.RIGHT},
		{key: 1, value: 1, level: 2, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 2, value: 2, level: 3, hasPrev: TRUE_, hasNext: TRUE_, hasLeft: FALSE, hasRight: TRUE_, hasParent: TRUE_, action: walker.RIGHT},
		{key: 3, value: 3, level: 4, hasPrev: TRUE_, hasNext: FALSE, hasLeft: FALSE, hasRight: FALSE, hasParent: TRUE_, action: walker.RETURN},
	}
	w := walker.New((f_walk)(func(f walker.F) {
		assertWalk(f, nodes, t)
	}))
	_, _, found := w.UpperBound(4)
	if found == true {
		t.Fatalf("walk returned true")
	}
}

// Test_Action_String
func Test_Action_String(t *testing.T) {
	assertActionString(walker.RETURN /**/, "RETURN" /**/, t)
	assertActionString(walker.PREV /****/, "PREV" /****/, t)
	assertActionString(walker.NEXT /****/, "NEXT" /****/, t)
	assertActionString(walker.LEFT /****/, "LEFT" /****/, t)
	assertActionString(walker.RIGHT /***/, "RIGHT" /***/, t)
	assertActionString(walker.PARENT /**/, "PARENT" /**/, t)

	assertActionString(walker.Action(-1), "walker.Action(-1)", t)
}
