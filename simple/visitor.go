package simple

import "fmt"

import (
	"github.com/iNamik/go_bst/visitor"
	"github.com/iNamik/go_cmp"
)

// tree::Visit
func (t *tree) Visit(key interface{}, f visitor.F) (value interface{}, result visitor.Result) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root, value, result, t.left = visit(t.root, key, t.fcmp, f, t.left)
	if result == visitor.INSERTED {
		t.size++
	} else if result == visitor.REMOVED {
		t.size--
	}
	return value, result
}

// visit
func visit(h *node, key interface{}, fcmp cmp.F, f visitor.F, left bool) (_ *node, value interface{}, result visitor.Result, _ bool) {
	if h == nil {
		var action visitor.Action
		value, action = f(nil, false)
		switch action {
		case visitor.INSERT:
			return &node{key: key, value: value}, value, visitor.INSERTED, left
		case visitor.GET:
			return nil, nil, visitor.NOT_FOUND, left
		default:
			panic(fmt.Sprintf("illegal action '%s' when visiting non-found key", action))
		}
	}
	switch fcmp(key, h.key) {
	case cmp.LT:
		h.left, value, result, left = visit(h.left, key, fcmp, f, left)
	case cmp.GT:
		h.right, value, result, left = visit(h.right, key, fcmp, f, left)
	default:
		var action visitor.Action
		value, action = f(h.value, true)
		switch action {
		case visitor.GET:
			value = h.value
			result = visitor.FOUND
		case visitor.REPLACE:
			h.value = value
			result = visitor.REPLACED
		case visitor.REMOVE:
			value = h.value
			h, left = removeNode(h, left)
			result = visitor.REMOVED
		default:
			panic(fmt.Sprintf("illegal action '%s' when visiting found key", action))
		}
	}
	return h, value, result, left
}
