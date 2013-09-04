go_bst/walker
=============

**Declares Extensible Method 'Walk' And Uses It To Implement Several BST Methods**


About
-----

Package `walker` specifies the extensible BST method `Walk`, which allows you to freely traverse a tree, either up and down (left, right, parent) or in sequential order (previous, next).


Example
-------

As an example of how `Walk` works, below is the code for methods that use `Walk` to implement `Max` and `ForeachMin`:

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

This package uses `Walk` to implement several BST methods:

 * Get (see `bst.I_Get`)
 * Min (see `finder.I_Min`)
 * Max (see `finder.I_Max`)
 * LowerBound (see `finder.I_LowerBound`)
 * UpperBound (see `finder.I_UpperBound`)
 * ForeachMin (see `walker.I_ForeachMin`)
 * ForeachMax (see `walker.I_ForeachMax`)


Efficiency
----------

For tree implementations that maintain parent information and support threading, `Walk` will likely be implemented very efficiently.

Implementations that use recursion shouldn't require any more stack space than `ReplaceOrInsert` or `Remove` would require.

`Walk` does not modify the tree, so tree implementations that do not modify the tree during `Get` operations are likely to be more efficient than ones that do.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>
