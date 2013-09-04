package simple

//import . "github.com/iNamik/go_pkg/debug/assert"

import (
	"github.com/iNamik/go_bst"
	"github.com/iNamik/go_bst/finder"
	"github.com/iNamik/go_bst/visitor"
	"github.com/iNamik/go_bst/walker"
	"github.com/iNamik/go_cmp"
)

import (
	"sync"
)

/**********************************************************************
 ** Types & Interfaces
 **********************************************************************/

// T
type T interface {
	bst.T
	finder.I
	visitor.I
	walker.I
	bst.I_Size
	finder.I_Min
	finder.I_Max
}

// node
type node struct {
	key   interface{}
	value interface{}
	left  *node
	right *node
}

// tree
type tree struct {
	mutex *sync.Mutex
	root  *node
	fcmp  cmp.F
	left  bool // To randomize removal of nodes
	size  int
}

/**********************************************************************
 ** Public Functions
 **********************************************************************/

// New
func New(fcmp cmp.F) T {
	return &tree{mutex: &sync.Mutex{}, root: nil, fcmp: fcmp, left: true, size: 0}
}

// tree:Empty
func (t *tree) Empty() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.size == 0
}

// tree:Size
func (t *tree) Size() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.size
}

// tree:ReplaceOrInsert
func (t *tree) ReplaceOrInsert(key interface{}, value interface{}) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	var replaced bool
	t.root, replaced = replaceOrInsert(t.root, key, value, t.fcmp)
	if !replaced {
		t.size++
	}
	return replaced
}

// tree::Get
func (t *tree) Get(key interface{}) (interface{}, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	h := get(t.root, key, t.fcmp)
	if h != nil {
		return h.value, true
	}
	return nil, false
}

// tree::Remove
func (t *tree) Remove(key interface{}) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	var removed bool
	t.root, removed, t.left = remove(t.root, key, t.fcmp, t.left)
	if removed {
		t.size--
	}
	return removed
}

/**********************************************************************
 ** Private Functions
 **********************************************************************/

// replaceOrInsert returns true if key was replaced, false if it was inserted into the tree
func replaceOrInsert(h *node, key interface{}, value interface{}, fcmp cmp.F) (*node, bool) {
	if h == nil {
		return &node{key: key, value: value}, false
	}
	replaced := true
	switch fcmp(key, h.key) {
	case cmp.LT:
		h.left, replaced = replaceOrInsert(h.left, key, value, fcmp)
	case cmp.GT:
		h.right, replaced = replaceOrInsert(h.right, key, value, fcmp)
	default:
		h.value = value
	}
	return h, replaced
}

// get
func get(h *node, key interface{}, fcmp cmp.F) *node {
	for h != nil {
		switch fcmp(key, h.key) {
		case cmp.LT:
			h = h.left
		case cmp.GT:
			h = h.right
		default:
			return h
		}
	}
	return nil
}

// remove
func remove(h *node, key interface{}, fcmp cmp.F, left bool) (*node, bool, bool) {
	removed := false
	newLeft := left
	if h != nil {
		switch fcmp(key, h.key) {
		case cmp.LT:
			h.left, removed, newLeft = remove(h.left, key, fcmp, left)
		case cmp.GT:
			h.right, removed, newLeft = remove(h.right, key, fcmp, left)
		default:
			h, newLeft = removeNode(h, left)
			removed = true
		}
	}
	return h, removed, newLeft
}

// removeNode
func removeNode(h *node, left bool) (*node, bool) {
	// If there are any children
	if h.left != nil || h.right != nil {
		var n *node = nil // Replacement node

		// If we want left or if there is no right
		if h.left != nil && (left || h.right == nil) {
			// If we have both left and right, then use right next time
			left = !(h.right != nil)
			// If there is no left.right node
			if h.left.right == nil {
				n = h.left
			} else {
				// Find parent of max(h.left)
				var nParent *node = h.left
				for nParent.right.right != nil {
					nParent = nParent.right
				}
				n = nParent.right
				nParent.right = n.left
				n.left = h.left
			}
			n.right = h.right

			// We want right or there is no left
		} else {
			// If we have both left and right, then use left next time
			left = (h.left != nil)
			// If there is no right.left node
			if h.right.left == nil {
				n = h.right
			} else {
				// Find parent of min(h.left)
				var nParent *node = h.right
				for nParent.left.left != nil {
					nParent = nParent.left
				}
				n = nParent.left
				nParent.left = n.right
				n.right = h.right
			}
			n.left = h.left
		}
		return n, left
	}
	return nil, left
}
