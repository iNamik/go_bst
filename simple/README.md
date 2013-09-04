go_bst/simple
=============

**Extensible Binary Search Tree (BST) Implementation in Go**


About
-----

Package `simple` serves as a reference implemention of an extensible Binary Search Tree as defined in the `go_bst` package and sub-packages.


Standard BST Methods
--------------------

All of the standard BST methods required to satisfy the `bst.T` interface have been implemented:

 * Empty
 * ReplaceOrInsert
 * Get
 * Remove


Extensible BST Methods
----------------------

All of the extensible interfaces have been implemented:

 * Find  (see `finder.T`)
 * Visit (see `visitor.T`)
 * Walk  (see `walker.T`)


Additional BST Methods
----------------------

The following additional BST methods have been implemented:

 * Size (see `bst.I_Size`)
 * Min  (see `finder.I_Min`)
 * Max  (see `finder.I_Max`)


Leaning
-------

In an attempt to avoid leaning in a particular direction, this implementation uses a 'toggle' mechanism to decide if it should remove from the left or the right when both options are available.


Effeciency
----------

Instead of storing parent and sibling information on each node, this implementation uses recursion to accomplish various tree navigation functions.  Specifically, the following functions use recursion:

 * ReplaceOrInsert
 * Remove
 * Visit
 * Walk

The remaining functions do not use recursion and can be considered efficient implementations.


License
-------

This package is released under the MIT License.
See included file 'LICENSE' for more details.


Contributors
------------

David Farell <DavidPFarrell@yahoo.com>
