package finder_test

//import . "github.com/iNamik/go_pkg/debug/assert"
//import . "github.com/iNamik/go_pkg/debug/ping"

import (
	"fmt"
	"testing"
)

import (
	"github.com/iNamik/go_bst/finder"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Types
 **********************************************************************/

// f_find allows us to attach finder.I onto a closure
type f_find func(finder.F) (key interface{}, value interface{}, found bool)

// f_find::Find
func (w f_find) Find(f finder.F) (key interface{}, value interface{}, found bool) {
	return w(f)
}

func New(nodes []node, t *testing.T) finder.T {
	return finder.New((f_find)(func(f finder.F) (key interface{}, value interface{}, found bool) {
		return assertFind(f, nodes, t)
	}))
}

// this is to get vertical-alignment in my data :)
const (
	TRUE_ = true
	FALSE = false
)

/**********************************************************************
 ** node
 **********************************************************************/

type node struct {
	key      interface{}
	value    interface{}
	hasLeft  bool
	hasRight bool
	action   finder.Action
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

// node::HasLeft
func (n *node) HasLeft() bool {
	return n.hasLeft
}

// node::HasRight
func (n *node) HasRight() bool {
	return n.hasRight
}

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

// assertFind
func assertFind(f finder.F, nodes []node, t *testing.T) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	for i := 0; i < len(nodes); i++ {
		n := &nodes[i]
		action := f(n)
		if action != n.action {
			t.Fatalf("node[%d], find returned '%s' instead of '%s'", i, action, n.action)
		}
		if action == finder.FOUND {
			key, value, found = n.key, n.value, true
		}
	}
	return
}

// assertActionString
func assertActionString(action finder.Action, value string, t *testing.T) {
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
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: FALSE, action: finder.FOUND},
	}
	value_, found := New(nodes, t).Get(VALUE)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_Get_Miss_Left
func Test_Get_Miss_Left(t *testing.T) {
	const VALUE = 1
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: FALSE, hasRight: FALSE, action: finder.LEFT},
	}
	_, found := New(nodes, t).Get(VALUE)
	if found == true {
		t.Fatalf("find returned true")
	}
}

// Test_Get_Miss_Right
func Test_Get_Miss_Right(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: FALSE, action: finder.RIGHT},
	}
	_, found := New(nodes, t).Get(VALUE)
	if found == true {
		t.Fatalf("find returned true")
	}
}

// Test_Min
func Test_Min(t *testing.T) {
	const VALUE = 0
	nodes := []node{
		{key: 3, value: 3, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 2, value: 2, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 1, value: 1, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 0, value: 0, hasLeft: FALSE, hasRight: FALSE, action: finder.FOUND},
	}
	_, value_, found := New(nodes, t).Min()
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_Max
func Test_Max(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 3, value: 3, hasLeft: FALSE, hasRight: FALSE, action: finder.FOUND},
	}
	_, value_, found := New(nodes, t).Max()
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Match
func Test_LowerBound_Match(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 3, value: 3, hasLeft: TRUE_, hasRight: FALSE, action: finder.FOUND},
	}
	_, value_, found := New(nodes, t).LowerBound(3)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Near_Right
func Test_LowerBound_Near_Right(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 4, value: 4, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: TRUE_, hasRight: FALSE, action: finder.RIGHT},
	}
	_, value_, found := New(nodes, t).LowerBound(3)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Near_Left
func Test_LowerBound_Near_Left(t *testing.T) {
	const VALUE = 5
	nodes := []node{
		{key: 5, value: 5, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 8, value: 8, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 7, value: 7, hasLeft: FALSE, hasRight: FALSE, action: finder.LEFT},
	}
	_, value_, found := New(nodes, t).LowerBound(6)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_LowerBound_Miss
func Test_LowerBound_Miss(t *testing.T) {
	nodes := []node{
		{key: 3, value: 3, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 2, value: 2, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 1, value: 1, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 0, value: 0, hasLeft: FALSE, hasRight: FALSE, action: finder.LEFT},
	}
	_, _, found := New(nodes, t).LowerBound(-1)
	if found == true {
		t.Fatalf("find returned true")
	}
}

// Test_UpperBound_Match
func Test_UpperBound_Match(t *testing.T) {
	const VALUE = 3
	nodes := []node{
		{key: 5, value: 5, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 3, value: 3, hasLeft: FALSE, hasRight: FALSE, action: finder.FOUND},
	}
	_, value_, found := New(nodes, t).UpperBound(3)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Near_Right
func Test_UpperBound_Near_Right(t *testing.T) {
	const VALUE = 4
	nodes := []node{
		{key: 5, value: 5, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 4, value: 4, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 2, value: 2, hasLeft: FALSE, hasRight: FALSE, action: finder.RIGHT},
	}
	_, value_, found := New(nodes, t).UpperBound(3)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Near_Left
func Test_UpperBound_Near_Left(t *testing.T) {
	const VALUE = 2
	nodes := []node{
		{key: 4, value: 4, hasLeft: TRUE_, hasRight: FALSE, action: finder.LEFT},
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: FALSE, hasRight: FALSE, action: finder.LEFT},
	}
	_, value_, found := New(nodes, t).UpperBound(1)
	if found == false {
		t.Fatalf("find returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("find returned '%d' instead of '%d'", value, VALUE)
	}
}

// Test_UpperBound_Miss
func Test_UpperBound_Miss(t *testing.T) {
	nodes := []node{
		{key: 0, value: 0, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 1, value: 1, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 2, value: 2, hasLeft: FALSE, hasRight: TRUE_, action: finder.RIGHT},
		{key: 3, value: 3, hasLeft: FALSE, hasRight: FALSE, action: finder.RIGHT},
	}
	_, _, found := New(nodes, t).UpperBound(4)
	if found == true {
		t.Fatalf("find returned true")
	}
}

// Test_Action_String
func Test_Action_String(t *testing.T) {
	assertActionString(finder.LEFT /*******/, "LEFT" /*******/, t)
	assertActionString(finder.RIGHT /******/, "RIGHT" /******/, t)
	assertActionString(finder.FOUND /******/, "FOUND" /******/, t)
	assertActionString(finder.NOT_FOUND /**/, "NOT_FOUND" /**/, t)

	assertActionString(finder.Action(-1), "finder.Action(-1)", t)
}
