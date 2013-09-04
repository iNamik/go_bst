package simple

import "fmt"

import (
	"github.com/iNamik/go_bst/finder"
	"github.com/iNamik/go_cmp"
)

// rnode
type rnode struct {
	n    *node
	fcmp cmp.F
}

// rnode::Key
func (r *rnode) Key() interface{} {
	return r.n.key
}

// rnode::Value
func (r *rnode) Value() interface{} {
	return r.n.value
}

// rnode::HasLeft
func (r *rnode) HasLeft() bool {
	return r.n.left != nil
}

// rnode::HasRight
func (r *rnode) HasRight() bool {
	return r.n.right != nil
}

// rnode::Cmp
func (r *rnode) Cmp(a interface{}, b interface{}) int {
	return r.fcmp(a, b)
}

// Find
func (t *tree) Find(f finder.F) (key interface{}, value interface{}, found bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for h := t.root; h != nil; {
		switch action := f(&rnode{n: h, fcmp: t.fcmp}); action {
		case finder.LEFT:
			h = h.left
		case finder.RIGHT:
			h = h.right
		case finder.FOUND:
			return h.key, h.value, true
		case finder.NOT_FOUND:
			return nil, nil, false
		default:
			panic(fmt.Sprintf("illegal find action '%s'", action))

		}
	}
	return nil, nil, false
}
