/*

Package finder specifies the extensible BST method Find,
which allows you to direct navigation down a tree, looking for a specific value.


Example
-------

As an example of how Find works, below is the code for a
method that uses Find to implement the standard BST Get:

	// Get retrieves the value of the specified key,
	// returning value,true if found and <undefined>,false otherwise.
	// @f is an object fulfilling the finder.I interface (Find() method).
	// @key is the key to search for.
	func Get(f finder.I, key interface{}) (value interface{}, found bool) {
		_, value, found = f.Find(func(node finder.Node) finder.Action {
			switch node.Cmp(key, node.Key()) { // key <=> node
			case cmp.LT: // key < node
				return finder.LEFT
			case cmp.GT: // key > node
				return finder.RIGHT
			default: // key == node
				return finder.RETURN
			}
		})
		return
	}


BST Methods Implemented
-----------------------

This package uses Find to implement several common BST methods:

 * Get (see bst.I_Get)
 * Min (see finder.I_Min)
 * Max (see finder.I_Max)
 * LowerBound (see finder.I_LowerBound)
 * UpperBound (see finder.I_UpperBound)


Efficiency
----------

Since navigation only goes down the tree, and since we are not
modifying the tree, it is assumed that Find can generally be
implemented efficiently and without recursion.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>

*/
package finder

import "fmt"

import (
	"github.com/iNamik/go_bst"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Types
 **********************************************************************/

// T specifies an interface for all of the functions that finder supports
type T interface {
	I_Get
	I_Min
	I_Max
	I_LowerBound
	I_UpperBound
}

// I defines the extensible finder interface
type I interface {
	Find(f F) (key interface{}, value interface{}, found bool)
}

// F defines the call-back function used for the Find method
type F func(Node) Action

// t
type t struct{ i I }

/**********************************************************************
 ** Node
 **********************************************************************/

// Node
type Node interface {
	Key() interface{}
	Value() interface{}
	Cmp(a interface{}, b interface{}) int
	HasLeft() bool
	HasRight() bool
}

/**********************************************************************
 ** Action
 **********************************************************************/

// Action
type Action int

// Action:String
func (a Action) String() string {
	if 0 <= a && a < Action(len(actions)) {
		return actions[a]
	}
	return fmt.Sprintf("finder.Action(%d)", a)
}

// Action Enums
const (
	LEFT Action = iota
	RIGHT
	FOUND
	NOT_FOUND
)

// actions
var actions = []string{
	LEFT:      "LEFT",
	RIGHT:     "RIGHT",
	FOUND:     "FOUND",
	NOT_FOUND: "NOT_FOUND",
}

/**********************************************************************
 ** New
 **********************************************************************/

// New creates a finder.T from a finder.I
func New(i I) T {
	return &t{i: i}
}

/**********************************************************************
 ** Get
 **********************************************************************/

// Get
// matches bst.I_Get
type I_Get interface {
	bst.I_Get
	//Get(key interface{}) (value interface{}, found bool)
}

// t::Get
func (f *t) Get(key interface{}) (interface{}, bool) {
	return Get(f.i, key)
}

// Get
func Get(f I, key interface{}) (value interface{}, found bool) {
	_, value, found = f.Find(func(node Node) Action {
		switch node.Cmp(key, node.Key()) { // key <=> node
		case cmp.LT: // key < node
			return LEFT
		case cmp.GT: // key > node
			return RIGHT
		default: // key == node
			return FOUND
		}
	})
	return
}

/**********************************************************************
 ** Min
 **********************************************************************/

// I_Min
type I_Min interface {
	// Min finds the minimum value in the tree
	Min() (key interface{}, value interface{}, found bool)
}

// t::Min
func (f *t) Min() (interface{}, interface{}, bool) {
	return Min(f.i)
}

// Min
func Min(f I) (interface{}, interface{}, bool) {
	return f.Find(func(node Node) Action {
		if node.HasLeft() {
			return LEFT
		}
		return FOUND
	})
}

/**********************************************************************
 ** Max
 **********************************************************************/

// I_Max
type I_Max interface {
	// Max finds the maximum value in the tree
	Max() (key interface{}, value interface{}, found bool)
}

// t::Max
func (f *t) Max() (interface{}, interface{}, bool) {
	return Max(f.i)
}

// Max
func Max(f I) (interface{}, interface{}, bool) {
	return f.Find(func(node Node) Action {
		if node.HasRight() {
			return RIGHT
		}
		return FOUND
	})
}

/**********************************************************************
 ** LowerBound
 **********************************************************************/

// I_LowerBound
type I_LowerBound interface {
	// LowerBound finds the Greatest Lower Bound (GLB),
	// the greatest key that is less than or equal to the provided key
	// http://en.wikipedia.org/wiki/Greatest_lower_bound
	LowerBound(boundKey interface{}) (key interface{}, value interface{}, found bool)
}

// t::LowerBound
func (f *t) LowerBound(boundKey interface{}) (interface{}, interface{}, bool) {
	return LowerBound(f.i, boundKey)
}

// LowerBound
func LowerBound(f I, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	f.Find(func(node Node) Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.LT: // node < boundKey - We have a candidate
			// If node > working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.GT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Keep looking
			return RIGHT
		case cmp.GT: // node > boundKey
			// Keep looking
			return LEFT
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return FOUND
		}
	})
	return
}

/**********************************************************************
 ** UpperBound
 **********************************************************************/

// I_UpperBound
type I_UpperBound interface {
	// UpperBound finds the Least Upper Bound (LUB),
	// the least key that is greater than or equal to the provided key
	// http://en.wikipedia.org/wiki/Least_upper_bound
	UpperBound(boundKey interface{}) (key interface{}, value interface{}, found bool)
}

// t::UpperBound
func (f *t) UpperBound(boundKey interface{}) (interface{}, interface{}, bool) {
	return UpperBound(f.i, boundKey)
}

// UpperBound
func UpperBound(f I, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	f.Find(func(node Node) Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.GT: // node > boundKey - We have a candidate
			// If node < working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.LT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Keep looking
			return LEFT
		case cmp.LT: // boundKey < node > boundKey
			// Keep looking
			return RIGHT
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return FOUND
		}
	})
	return
}
