package simple

import "fmt"

import (
	"github.com/iNamik/go_bst/walker"
	"github.com/iNamik/go_cmp"
)

// private walk actions
const (
	w_min walker.Action = 100 + iota
	w_max
	w_node
	w_parent
	w_lparent
	w_rparent
	w_child
	w_none walker.Action = -1
)

// wnode
type wnode struct {
	n     *node
	fcmp  cmp.F
	level int
	lp    *node
	rp    *node
}

// wnode::Key
func (w *wnode) Key() interface{} {
	return w.n.key
}

// wnode::Value
func (w *wnode) Value() interface{} {
	return w.n.value
}

// wnode::Cmp
func (w *wnode) Cmp(a interface{}, b interface{}) int {
	return w.fcmp(a, b)
}

// wnode::Level
func (w *wnode) Level() int {
	return w.level
}

// wnode::HasPrev
func (w *wnode) HasPrev() bool {
	return w.n.left != nil || w.lp != nil
}

// wnode::HasNext
func (w *wnode) HasNext() bool {
	return w.n.right != nil || w.rp != nil
}

// wnode::HasLeft
func (w *wnode) HasLeft() bool {
	return w.n.left != nil
}

// wnode::HasRight
func (w *wnode) HasRight() bool {
	return w.n.right != nil
}

// wnode::HasParent
func (w *wnode) HasParent() bool {
	// If both nil, then node is root, no parent.
	// If only one not-nil, then its the parent.
	// If both not-nil, then one is parent.
	return w.lp != nil || w.rp != nil
}

// tree::Walk
func (t *tree) Walk(f walker.F) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	// We don't walk an empty tree
	if t.root == nil {
		return
	}
	walk(t.root, nil, nil, w_node, 1, t.fcmp, f)
}

// walk uses recursion to support walking up and down the tree.
// If our tree node contained a reference to parent, this would
// probably be much easier.
func walk(h *node, lp *node, rp *node, action walker.Action, level int, fcmp cmp.F, f walker.F) walker.Action {
	var cparent, caction walker.Action
	var cnode, clp, crp *node
	for {
		switch action {
		// Visit the current node
		case w_node:
			action = f(&wnode{n: h, fcmp: fcmp, level: level, lp: lp, rp: rp})

			// Visit a child node
		case w_child:
			action = walk(cnode, clp, crp, caction, level+1, fcmp, f)

			// If next action is for a parent, and we're that parent
			if action == walker.PARENT || action == cparent {
				action = w_node // Visit ourselves
			}

			// Visit the minimum node. Used internally to support NEXT functionality
		case w_min:
			// Do I have a lesser child?
			if h.left != nil {
				action, cparent, cnode, clp, crp, caction = w_child, w_rparent, h.left, lp, h, w_min
			} else {
				action = w_node // We are the min, visit ourselves
			}

			// Visit the maximum node.  Used internally to support PREV fucionality
		case w_max:
			// Do I have a greator child?
			if h.right != nil {
				action, cparent, cnode, clp, crp, caction = w_child, w_lparent, h.right, h, rp, w_max
			} else {
				action = w_node // We are the max, visit ourselves
			}

			// Visit the left child
		case walker.LEFT:
			if h.left == nil {
				panic("cannot walk left when hasLeft() == false")
			}
			action, cparent, cnode, clp, crp, caction = w_child, w_rparent, h.left, lp, h, w_node

			// Visit the right child
		case walker.RIGHT:
			if h.right == nil {
				panic("cannot walk right when hasRight() == false")
			}
			action, cparent, cnode, clp, crp, caction = w_child, w_lparent, h.right, h, rp, w_node

			// Visit the previous node
		case walker.PREV:
			// Do I have a lesser child?
			if h.left != nil {
				// The PREV node is max(me.left)
				action, cparent, cnode, clp, crp, caction = w_child, w_rparent, h.left, lp, h, w_max

				// Do I have a lesser parent?
			} else if lp != nil {
				action = w_lparent
			} else {
				panic("cannot walk prev when hasPrev() == false")
			}

			// Visit the next node
		case walker.NEXT:
			// Do I have a greater child?
			if h.right != nil {
				// The NEXT node is min(me.right)
				action, cparent, cnode, clp, crp, caction = w_child, w_lparent, h.right, h, rp, w_min

				// Do I have a greater parent?
			} else if rp != nil {
				action = w_rparent
			} else {
				panic("cannot walk next when hasNext() == false")
			}

			// Visit a parent node
		case walker.PARENT, w_lparent, w_rparent:
			// If I have no parents
			if lp == nil && rp == nil {
				panic("cannot walk parent when hasParent() == false")
			}
			return action

			// Return from walk
		case walker.RETURN:
			return walker.RETURN

			// Unknown walk action
		default:
			panic(fmt.Sprintf("illegal walk action '%s'", action))
		}
	}
}
