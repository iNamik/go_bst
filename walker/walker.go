/*

Package walker specifies the extensible BST method Walk,
which allows you to freely traverse a tree, either up and down
(left, right, parent) or in sequential order (previous, next).


Example
-------

As an example of how Walk works, below is the code for methods
that use Walk to implement Max and ForeachMin:

	// Max returns the maximum element in tree.  If the tree
	// is empty, it will return (nil, nil, false), else it will
	// return (maxKey, maxValue, true)
	// @w is an object fulfilling the walker.I interface (Walk() method)
	func Max(w I) (key interface{}, value interface{}, found bool) {
		key, value, found = nil, nil, false
		w.Walk(func(node Node) Action {
			if node.HasRight() {
				return RIGHT
			}
			key, value, found = node.Key(), node.Value(), true
			return RETURN
		})
		return
	}

	// ForeachMin iterates sequentially over the the tree,
	// starting at the minimum value.
	// @w is an object fulfilling the walker.I interface (Walk() method)
	// @f is a call-back function, which will be called at each iteration.
	func ForeachMin(w I, f F_Visit) {
		haveMin := false
		w.Walk(func(n Node) Action {
			if haveMin == false {
				if n.HasLeft() == true {
					return LEFT
				}
				haveMin = true
			}
			f(n.Key(), n.Value())
			if n.HasNext() {
				return NEXT
			}
			return RETURN
		})
	}


BST Methods Implemented
-----------------------

This package uses Walk to implement several BST methods:

 * Get (see bst.I_Get)
 * Min (see finder.I_Min)
 * Max (see finder.I_Max)
 * LowerBound (see finder.I_LowerBound)
 * UpperBound (see finder.I_UpperBound)
 * ForeachMin (see walker.I_ForeachMin)
 * ForeachMax (see walker.I_ForeachMax)


Efficiency
----------

For tree implementations that maintain parent information and
support threading, Walk will likely be implemented very
efficiently.

Implementations that use recursion shouldn't require any more
stack space than ReplaceOrInsert or Remove would require.

Walk does not modify the tree, so tree implementations that
do not modify the tree during Get operations are likely to
be more efficient than ones that do.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>

*/
package walker

import "fmt"

import (
	"github.com/iNamik/go_bst"
	"github.com/iNamik/go_bst/finder"
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Types
 **********************************************************************/

// T specifies an interface for all of the functions that walker supports.
// extends finder.T
type T interface {
	finder.T // extends finder.T
	I_ForeachMin
	I_ForeachMax
}

// I defines the extensible walker interface
type I interface {
	Walk(F)
}

// F defines the call-back function used for the Walk method
type F func(Node) Action

// F_Visit defines a function for visiting nodes without controls (i.e. iteration)
type F_Visit func(key interface{}, value interface{})

type t struct{ i I }

/**********************************************************************
 ** Node
 **********************************************************************/

// Node extends finder.Node
type Node interface {
	finder.Node // extends finder.Node
	Level() int
	HasParent() bool
	HasPrev() bool
	HasNext() bool
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
	return fmt.Sprintf("walker.Action(%d)", a)
}

// Action Enums
const (
	RETURN Action = iota
	PREV
	NEXT
	LEFT
	RIGHT
	PARENT
)

// actions
var actions = []string{
	RETURN: "RETURN",
	PREV:   "PREV",
	NEXT:   "NEXT",
	LEFT:   "LEFT",
	RIGHT:  "RIGHT",
	PARENT: "PARENT",
}

/**********************************************************************
 ** New
 **********************************************************************/

// New creates a walker.T from a walker.I
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
}

// t::Get
func (w *t) Get(key interface{}) (value interface{}, found bool) {
	return Get(w.i, key)
}

// Get
func Get(w I, key interface{}) (value interface{}, found bool) {
	w.Walk(func(node Node) Action {
		switch node.Cmp(key, node.Key()) {
		case cmp.LT:
			if node.HasLeft() == true {
				return LEFT
			}
			value, found = nil, false
			return RETURN
		case cmp.GT:
			if node.HasRight() == true {
				return RIGHT
			}
			value, found = nil, false
			return RETURN
		default:
			value, found = node.Value(), true
			return RETURN
		}
	})
	return value, found
}

/**********************************************************************
 ** Min
 **********************************************************************/

// I_Min
// matches finder.I_Min
type I_Min interface {
	finder.I_Min
}

// t::Min
func (w *t) Min() (key interface{}, value interface{}, found bool) {
	return Min(w.i)
}

// Min
func Min(w I) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	w.Walk(func(node Node) Action {
		if node.HasLeft() {
			return LEFT
		}
		key, value, found = node.Key(), node.Value(), true
		return RETURN
	})
	return key, value, found
}

/**********************************************************************
 ** Max
 **********************************************************************/

// I_Max
// matches finder.I_Max
type I_Max interface {
	finder.I_Max
}

// t::Max
func (w *t) Max() (interface{} /* key */, interface{} /* value */, bool /* found */) {
	return Max(w.i)
}

// Max
func Max(w I) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	w.Walk(func(node Node) Action {
		if node.HasRight() {
			return RIGHT
		}
		key, value, found = node.Key(), node.Value(), true
		return RETURN
	})
	return
}

/**********************************************************************
 ** LowerBound
 **********************************************************************/

// I_LowerBound
// matches finder.I_LowerBound
type I_LowerBound interface {
	finder.I_LowerBound
}

// t::LowerBound
func (w *t) LowerBound(boundKey interface{}) (key interface{}, value interface{}, found bool) {
	return LowerBound(w.i, boundKey)
}

// LowerBound
func LowerBound(w I, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	w.Walk(func(node Node) Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.LT: // node < boundKey - We have a candidate
			// If node > working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.GT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Is there possibly a greater candidate out there?
			if node.HasRight() {
				// Keep looking
				return RIGHT
			}
			// No more keys to check, stick with what we got
			return RETURN
		case cmp.GT: // node > boundKey
			// If more possible candidates
			if node.HasLeft() == true {
				// Keep looking
				return LEFT
			}
			return RETURN
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return RETURN
		}
	})
	return key, value, found
}

/**********************************************************************
 ** UpperBound
 **********************************************************************/

// I_UpperBound
// matches finder.I_UpperBound
type I_UpperBound interface {
	finder.I_UpperBound
}

// t::UpperBound
func (w *t) UpperBound(boundKey interface{}) (key interface{}, value interface{}, found bool) {
	return UpperBound(w.i, boundKey)
}

// UpperBound
func UpperBound(w I, boundKey interface{}) (key interface{}, value interface{}, found bool) {
	key, value, found = nil, nil, false
	w.Walk(func(node Node) Action {
		switch node.Cmp(node.Key(), boundKey) {
		case cmp.GT: // node > boundKey - We have a candidate
			// If node < working then update working
			if found == false || node.Cmp(node.Key(), key) == cmp.LT {
				key, value, found = node.Key(), node.Value(), true
			}
			// Is there possibly a greater candidate out there?
			if node.HasLeft() {
				// Keep looking
				return LEFT
			}
			// No more keys to check, stick with what we got
			return RETURN
		case cmp.LT: // boundKey < node > boundKey
			// If more possible candidates
			if node.HasRight() == true {
				// Keep looking
				return RIGHT
			}
			return RETURN
		default: // node == boundKey
			key, value, found = node.Key(), node.Value(), true
			return RETURN
		}
	})
	return key, value, found
}

/**********************************************************************
 ** ForeachMin
 **********************************************************************/

// I_ForeachMin
type I_ForeachMin interface {
	// ForeachMin iterates sequentially over the the tree,
	// starting at the minimum value
	ForeachMin(f F_Visit)
}

// t::ForeachMin
func (w *t) ForeachMin(f F_Visit) {
	ForeachMin(w.i, f)
}

// ForeachMin
func ForeachMin(w I, f F_Visit) {
	haveMin := false
	w.Walk(func(n Node) Action {
		if haveMin == false {
			if n.HasLeft() == true {
				return LEFT
			}
			haveMin = true
		}
		f(n.Key(), n.Value())
		if n.HasNext() {
			return NEXT
		}
		return RETURN
	})
}

/**********************************************************************
 ** ForeachMax
 **********************************************************************/

// I_ForeachMax
type I_ForeachMax interface {
	// ForeachMax iterates sequentially over the the tree,
	// starting at the maximum value, and ending at the minimum value.
	ForeachMax(f F_Visit)
}

// t::ForeachMax
func (w *t) ForeachMax(f F_Visit) {
	ForeachMax(w.i, f)
}

// ForeachMax
func ForeachMax(w I, f F_Visit) {
	haveMax := false
	w.Walk(func(n Node) Action {
		if haveMax == false {
			if n.HasRight() == true {
				return RIGHT
			}
			haveMax = true
		}
		f(n.Key(), n.Value())
		if n.HasPrev() {
			return PREV
		}
		return RETURN
	})
}
