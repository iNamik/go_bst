package simple

import . "github.com/iNamik/go_pkg/debug/assert"

//import . "github.com/iNamik/go_pkg/debug/ping"

// Min
func (t *tree) Min() (interface{}, interface{}, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.root == nil {
		return nil, nil, false
	}
	h := min(t.root)
	return h.key, h.value, true
}

// Max
func (t *tree) Max() (interface{}, interface{}, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.root == nil {
		return nil, nil, false
	}
	h := max(t.root)
	return h.key, h.value, true
}

// min
func min(h *node) *node {
	Assert(h != nil)
	for h.left != nil {
		h = h.left
	}
	return h
}

// max
func max(h *node) *node {
	Assert(h != nil)
	for h.right != nil {
		h = h.right
	}
	return h
}
