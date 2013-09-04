go_bst/finder
=============

**Declares Extensible Method 'Find' And Uses It To Implement Several Common BST Methods**


About
-----

Package finder specifies the extensible BST method `Find`, which allows you to direct navigation down a tree, looking for a specific value.


Example
-------

As an example of how `Find` works, below is the code for a method that uses `Find` to implement the standard BST `Get`:

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

This package uses `Find` to implement several common BST methods:

 * Get (see `bst.I_Get`)
 * Min (see `finder.I_Min`)
 * Max (see `finder.I_Max`)
 * LowerBound (see `finder.I_LowerBound`)
 * UpperBound (see `finder.I_UpperBound`)


Efficiency
----------

Since navigation only goes down the tree, and since we are not modifying the tree, it is assumed that `Find` can generally be implemented efficiently and without recursion.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>
