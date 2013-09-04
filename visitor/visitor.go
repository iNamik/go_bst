/*

Package visitor specifies the extensible BST method Visit,
which allows you to defer taking an action on a tree node
(insert, get, update, delete) until after you've visited that
node (i.e. inspect then change with one visit to the node).

Example
-------

As an example of how Visit works, below is the code for a
method that uses Visit to implement GetAndRemove:

	// GetAndRemove removes the specified key from the tree,
	// returning the value of the removed item.
	// It returns (value,true) if the key was found
	// (and hence, removed), and (<undefined>,false) otherwise.
	// @v is an object fulfilling the visitor.I interface (Visit() method).
	// @key is the key to search for.
	func GetAndRemove(v I, key interface{}) (value interface{}, _ bool) {
		_, result := v.Visit(key, func(oldValue interface{}, found bool) (interface{}, Action) {
			if found {
				value = oldValue
				return nil, REMOVE
			}
			value = nil
			return nil, GET
		})
		return value, result == REMOVED // removed == found
	}


BST Methods Implemented
-----------------------

This package uses Visit to implement several BST methods:

 * Get                   (see bst.I_Get)
 * GetOrInsert           (see visitor.I_GetOrInsert)
 * GetAndReplace         (see visitor.I_GetAndReplace)
 * GetAndReplaceOrInsert (see visitor.I_GetAndReplaceOrInsert)
 * GetAndRemove          (see visitor.I_GetAndRemove)
 * GetAndRemoveOrInsert  (see visitor.I_GetAndRemoveOrInsert)
 * Replace               (see visitor.I_Replace)
 * ReplaceOrInsert       (see bst.I_ReplaceOrInsert)
 * Remove                (see bst.I_Remove)
 * RemoveOrInsert        (see visitor.I_RemoveOrInsert)

*NOTE* These functions represent all of the default combinations
for Get/Insert/Replace/Remove.  Although some may not seem very
useful, they are included for completeness.


Efficiency
----------

The idea of Visit is that it can often be more efficient than
calling Get() followed by a call to  ReplaceOrInsert()/Remove().

For tree implementations that don't modify the tree on the
way down during inserts or deletes, Visit is almost certainly
more effecient.

For implementations that do modify the tree on the way down,
Visit may be less efficient if you are not guaranteed to do
an insert/update/remove.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>

*/
package visitor

import "fmt"
import "github.com/iNamik/go_bst"

/**********************************************************************
 ** Types
 **********************************************************************/

// T specifies an interface for all of the functions that visitor supports
type T interface {
	I_Get
	I_GetOrInsert
	I_GetAndReplace
	I_GetAndReplaceOrInsert
	I_GetAndRemove
	I_GetAndRemoveOrInsert
	I_Replace
	I_ReplaceOrInsert
	I_Remove
	I_RemoveOrInsert
}

// I defines the extensible visitor interface
type I interface {
	Visit(key interface{}, f F) (value interface{}, result Result)
}

// F defines the call-back function used for the Visit method
type F func(value interface{}, found bool) (newValue interface{}, action Action)

// t
type t struct{ i I }

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
	return fmt.Sprintf("visitor.Action(%d)", a)
}

// Action Enums
const (
	INSERT Action = iota
	GET
	REPLACE
	REMOVE
)

// actions
var actions = []string{
	INSERT:  "INSERT",
	GET:     "GET",
	REPLACE: "REPLACE",
	REMOVE:  "REMOVE",
}

/**********************************************************************
 ** Result
 **********************************************************************/

// Result
type Result int

// Result::String
func (r Result) String() string {
	if 0 <= r && r < Result(len(results)) {
		return results[r]
	}
	return fmt.Sprintf("visitor.Result(%d)", r)
}

// Result Enums
const (
	INSERTED Result = iota
	NOT_FOUND
	FOUND
	REPLACED
	REMOVED
)

// results
var results = []string{
	INSERTED:  "INSERTED",
	NOT_FOUND: "NOT_FOUND",
	FOUND:     "FOUND",
	REPLACED:  "REPLACED",
	REMOVED:   "REMOVED",
}

/**********************************************************************
 ** New
 **********************************************************************/

// New creates a visitor.T from a visitor.I
func New(i I) T {
	return &t{i: i}
}

/**********************************************************************
 ** Get
 **********************************************************************/

// I_Get
// matches bst.I_Get
type I_Get interface {
	bst.I_Get
	//Get(key interface{}) (value interface{}, found bool)
}

// t::Get
func (v *t) Get(key interface{}) (interface{}, bool) {
	return Get(v.i, key)
}

// Get
func Get(v I, key interface{}) (interface{}, bool) {
	value, result := v.Visit(key, func(_ interface{}, _ bool) (interface{}, Action) {
		return nil, GET
	})
	return value, result == FOUND
}

/**********************************************************************
 ** GetOrInsert
 **********************************************************************/

// GetOrInsert
type I_GetOrInsert interface {
	GetOrInsert(key interface{}, newValue interface{}) (value interface{}, found bool)
}

// t::GetOrInsert
func (v *t) GetOrInsert(key interface{}, newValue interface{}) (interface{}, bool) {
	return GetOrInsert(v.i, key, newValue)
}

// GetOrInsert
func GetOrInsert(v I, key interface{}, newValue interface{}) (interface{}, bool) {
	value, result := v.Visit(key, func(_ interface{}, found bool) (interface{}, Action) {
		if found {
			return nil, GET
		}
		return newValue, INSERT
	})
	return value, result == FOUND // !found = inserted
}

/**********************************************************************
 ** GetAndReplace
 **********************************************************************/

// GetAndReplace
type I_GetAndReplace interface {
	GetAndReplace(key interface{}, newValue interface{}) (oldValue interface{}, found bool)
}

// t::GetAndReplace
func (v *t) GetAndReplace(key interface{}, newValue interface{}) (interface{}, bool) {
	return GetAndReplace(v.i, key, newValue)
}

// GetAndReplace
func GetAndReplace(v I, key interface{}, newValue interface{}) (value interface{}, _ bool) {
	_, result := v.Visit(key, func(oldValue interface{}, found bool) (interface{}, Action) {
		if found {
			value = oldValue
			return newValue, REPLACE
		}
		value = nil
		return nil, GET
	})
	return value, result == FOUND // found == replaced
}

/**********************************************************************
 ** GetAndReplaceOrInsert
 **********************************************************************/

// GetAndReplaceOrInsert
type I_GetAndReplaceOrInsert interface {
	GetAndReplaceOrInsert(key interface{}, newValue interface{}) (value interface{}, found bool)
}

// t::GetAndReplaceOrInsert
func (v *t) GetAndReplaceOrInsert(key interface{}, newValue interface{}) (interface{}, bool) {
	return GetAndReplaceOrInsert(v.i, key, newValue)
}

// GetAndReplaceOrInsert
func GetAndReplaceOrInsert(v I, key interface{}, newValue interface{}) (value interface{}, _ bool) {
	_, result := v.Visit(key, func(oldValue interface{}, found bool) (interface{}, Action) {
		if found {
			value = oldValue
			return newValue, REPLACE
		}
		value = newValue
		return newValue, INSERT
	})
	return value, result == REPLACED // replaced == found, !replaced == inserted
}

/**********************************************************************
 ** GetAndRemove
 **********************************************************************/

// GetAndRemove
type I_GetAndRemove interface {
	GetAndRemove(key interface{}) (value interface{}, found bool)
}

// t::GetAndRemove
func (v *t) GetAndRemove(key interface{}) (value interface{}, found bool) {
	return GetAndRemove(v.i, key)
}

// GetAndRemove
func GetAndRemove(v I, key interface{}) (value interface{}, _ bool) {
	_, result := v.Visit(key, func(oldValue interface{}, found bool) (interface{}, Action) {
		if found {
			value = oldValue
			return nil, REMOVE
		}
		value = nil
		return nil, GET
	})
	return value, result == REMOVED // removed == found
}

/**********************************************************************
 ** GetAndRemoveOrInsert
 **********************************************************************/

// GetAndRemoveOrInsert
type I_GetAndRemoveOrInsert interface {
	GetAndRemoveOrInsert(key interface{}, newValue interface{}) (value interface{}, found bool)
}

// t::GetAndRemoveOrInsert
func (v *t) GetAndRemoveOrInsert(key interface{}, newValue interface{}) (interface{}, bool) {
	return GetAndRemoveOrInsert(v.i, key, newValue)
}

// GetAndRemoveOrInsert
func GetAndRemoveOrInsert(v I, key interface{}, newValue interface{}) (value interface{}, _ bool) {
	_, result := v.Visit(key, func(oldValue interface{}, found bool) (interface{}, Action) {
		if found {
			value = oldValue
			return nil, REMOVE
		}
		value = newValue
		return newValue, INSERT
	})
	return value, result == REMOVED // removed == found, !removed = inserted
}

/**********************************************************************
 ** Replace
 **********************************************************************/

// Replace
type I_Replace interface {
	Replace(key interface{}, newValue interface{}) (replaced bool)
}

// t::Replace
func (v *t) Replace(key interface{}, newValue interface{}) bool {
	return Replace(v.i, key, newValue)
}

// Replace
func Replace(v I, key interface{}, newValue interface{}) bool {
	_, result := v.Visit(key, func(_ interface{}, found bool) (interface{}, Action) {
		if found {
			return newValue, REPLACE
		}
		return nil, GET
	})
	return result == REPLACED // replaced == found, !replaced == !found
}

/**********************************************************************
 ** ReplaceOrInsert
 **********************************************************************/

// ReplaceOrInsert
// matches bst.I_ReplaceOrInsert
type I_ReplaceOrInsert interface {
	bst.I_ReplaceOrInsert
	//ReplaceOrInsert(key interface{}, value interface{}) (replaced bool)
}

// t::ReplaceOrInsert
func (v *t) ReplaceOrInsert(key interface{}, newValue interface{}) bool {
	return ReplaceOrInsert(v.i, key, newValue)
}

// ReplaceOrInsert
func ReplaceOrInsert(v I, key interface{}, newValue interface{}) bool {
	_, result := v.Visit(key, func(_ interface{}, found bool) (interface{}, Action) {
		if found {
			return newValue, REPLACE
		}
		return newValue, INSERT
	})
	return result == REPLACED // replaced == found, !replaced == inserted
}

/**********************************************************************
 ** Remove
 **********************************************************************/

// Remove
// matches bst.I_Remove
type I_Remove interface {
	bst.I_Remove
	//Remove(key interface{}) (removed bool)
}

// t::Remove
func (v *t) Remove(key interface{}) bool {
	return Remove(v.i, key)
}

// Remove
func Remove(v I, key interface{}) bool {
	_, result := v.Visit(key, func(_ interface{}, found bool) (interface{}, Action) {
		if found {
			return nil, REMOVE
		}
		return nil, GET
	})
	return result == REMOVED // removed = found
}

/**********************************************************************
 ** RemoveOrInsert
 **********************************************************************/

// RemoveOrInsert
type I_RemoveOrInsert interface {
	RemoveOrInsert(key interface{}, newValue interface{}) (removed bool)
}

// t::RemoveOrInsert
func (v *t) RemoveOrInsert(key interface{}, newValue interface{}) bool {
	return RemoveOrInsert(v.i, key, newValue)
}

// RemoveOrInsert
func RemoveOrInsert(v I, key interface{}, newValue interface{}) bool {
	_, result := v.Visit(key, func(_ interface{}, found bool) (interface{}, Action) {
		if found {
			return nil, REMOVE
		}
		return newValue, INSERT
	})
	return result == REMOVED // removed == found, !removed == inserted
}
