go_bst/visitor
==============

**Declares Extensible Method 'Visit' And Uses It To Implement Several BST Methods**


About
-----

Package `visitor` specifies the extensible BST method `Visit`, which allows you to defer taking an action on a tree node (insert, get, update, delete) until after you've visited that node (i.e. inspect then change with one visit to the node).


Example
-------

As an example of how `Visit` works, below is the code for a method that uses `Visit` to implement `GetAndRemove`:

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

This package uses `Visit` to implement several BST methods:

 * Get                   (see `bst.I_Get`)
 * GetOrInsert           (see `visitor.I_GetOrInsert`)
 * GetAndReplace         (see `visitor.I_GetAndReplace`)
 * GetAndReplaceOrInsert (see `visitor.I_GetAndReplaceOrInsert`)
 * GetAndRemove          (see `visitor.I_GetAndRemove`)
 * GetAndRemoveOrInsert  (see `visitor.I_GetAndRemoveOrInsert`)
 * Replace               (see `visitor.I_Replace`)
 * ReplaceOrInsert       (see `bst.I_ReplaceOrInsert`)
 * Remove                (see `bst.I_Remove`)
 * RemoveOrInsert        (see `visitor.I_RemoveOrInsert`)

*NOTE* These functions represent all of the default combinations for Get/Insert/Replace/Remove.  Although some may not seem very useful, they are included for completeness.


Efficiency
----------

The idea of `Visit` is that it can often be more efficient than calling `Get()` followed by a call to  `ReplaceOrInsert()`/`Remove()`.

For tree implementations that don't modify the tree on the way down during inserts or deletes, `Visit` is almost certainly more effecient.

For implementations that do modify the tree on the way down, `Visit` may be less efficient if you are not guaranteed to do an insert/update/remove.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>
