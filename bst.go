/*

Package bst and its sub-packages specify an extensible
Binary Search Tree (BST) API.

Instead of trying to think of every method a user might want
on a BST, we declare a set of extensible methods with the hopes
of enabling a user to create any functionality they need.


Packages and sub-packages:
--------------------------

* bst

Declares basic BST methods (Insert, Get, Remove, etc).

* bst/finder

Declares extensible method, Find, which allows you to direct
navigation down a tree, looking for a specific value.

* bst/visitor

Declares extensible method, Visit, which allows you to defer
taking an action on a tree node (insert, get, update, delete)
until after you've visited that node (i.e. inspect then change
with one visit to the node).

* bst/walker

Declares extensible method, Walk, which allows you to freely
traverse a tree, either up and down (left, right, parent) or
in sequential order (prev, next).

* bst/simple

Provides a reference implementation of an Extensible BST that
implements all of the above-declared methods.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>

*/
package bst

/**********************************************************************
 ** Tree Type
 **********************************************************************/

// T defines the minimal set of functions we'd like to see on a tree
type T interface {
	I_Empty
	I_ReplaceOrInsert
	I_Get
	I_Remove
}

/**********************************************************************
 ** Optional Interfaces that a tree might implement
 **********************************************************************/

// I_Empty
type I_Empty interface {
	Empty() bool
}

// I_Size
type I_Size interface {
	Size() int
}

// I_ReplaceOrInsert
type I_ReplaceOrInsert interface {
	ReplaceOrInsert(key interface{}, value interface{}) (replaced bool)
}

// I_Get retrieves the value of the specified key,
// returning value,true if found and <undefined>,false otherwise.
type I_Get interface {
	Get(key interface{}) (value interface{}, found bool)
}

// I_Remove
type I_Remove interface {
	Remove(key interface{}) (removed bool)
}
