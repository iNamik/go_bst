package visitor_test

//import . "github.com/iNamik/go_pkg/debug/assert"
//import . "github.com/iNamik/go_pkg/debug/ping"

import (
	"fmt"
	"testing"
)

import (
	"github.com/iNamik/go_bst/visitor"
)

/**********************************************************************
 ** Types
 **********************************************************************/

// f_visit allows us to attach visitor.I onto a closure
type f_visit func(interface{}, visitor.F) (interface{}, visitor.Result)

// f_visit::Visit
func (v f_visit) Visit(key interface{}, f visitor.F) (interface{}, visitor.Result) { return v(key, f) }

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

// assertVisit
func assertVisit(f visitor.F, value interface{}, found bool, newValue interface{}, action visitor.Action, t *testing.T) {
	nv, a := f(value, found)
	if a != action {
		t.Fatalf("visit returned action '%s' instead of '%s'", a, action)
	}
	if nv != newValue {
		t.Fatalf("visit returned newValue '%v' instead of '%v'", nv, newValue)
	}
}

// assertActionString
func assertActionString(action visitor.Action, value string, t *testing.T) {
	action_string := fmt.Sprint(action)
	if action_string != value {
		t.Fatalf("action code '%d' has String() value '%s' instead of '%s'", action, action_string, value)
	}
}

// assertResultString
func assertResultString(result visitor.Result, value string, t *testing.T) {
	result_string := fmt.Sprint(result)
	if result_string != value {
		t.Fatalf("result code '%d' has String() value '%s' instead of '%s'", result, result_string, value)
	}
}

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Get_Found
func Test_Get_Found(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, VALUE, true, nil, visitor.GET, t)
		return VALUE, visitor.FOUND
	}))
	value_, found := w.Get(VALUE)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("visit returned '%v' instead of '%d'", value, VALUE)
	}
}

// Test_Get_NotFound
func Test_Get_NotFound(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, nil, visitor.GET, t)
		return nil, visitor.NOT_FOUND
	}))
	_, found := w.Get(VALUE)
	if found == true {
		t.Fatalf("visit returned true")
	}
}

// Test_GetOrInsert_Get
func Test_GetOrInsert(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, VALUE, true, nil, visitor.GET, t)
		return VALUE, visitor.FOUND
	}))
	value_, found := w.Get(VALUE)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("visit returned '%v' instead of '%d'", value, VALUE)
	}
}

// Test_GetOrInsert_Found
func Test_GetOrInsert_Found(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, VALUE, true, nil, visitor.GET, t)
		return VALUE, visitor.FOUND
	}))
	value_, found := w.GetOrInsert(VALUE, VALUE)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("visit returned '%v' instead of '%d'", value, VALUE)
	}
}

// Test_GetOrInsert_NotFound
func Test_GetOrInsert_NotFound(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, VALUE, visitor.INSERT, t)
		return VALUE, visitor.INSERTED
	}))
	value_, found := w.GetOrInsert(VALUE, VALUE)
	if found == true {
		t.Fatalf("visit returned true")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("visit returned '%v' instead of '%d'", value, VALUE)
	}
}

// Test_GetAndReplace_Found
func Test_GetAndReplace_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, NEW, visitor.REPLACE, t)
		return NEW, visitor.FOUND
	}))
	value_, found := w.GetAndReplace(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != OLD {
		t.Fatalf("visit returned '%v' instead of '%d'", value, OLD)
	}
}

// Test_GetAndReplace_NotFound
func Test_GetAndReplace_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, nil, visitor.GET, t)
		return nil, visitor.NOT_FOUND
	}))
	value_, found := w.GetAndReplace(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
	if value_ != nil {
		t.Fatalf("visit returned '%v' instead of nil", value_)
	}
}

// Test_GetAndReplaceOrInsert_Found
func Test_GetAndReplaceOrInsert_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, NEW, visitor.REPLACE, t)
		return NEW, visitor.REPLACED
	}))
	value_, found := w.GetAndReplaceOrInsert(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != OLD {
		t.Fatalf("visit returned '%v' instead of '%d'", value, OLD)
	}
}

// Test_GetAndReplaceOrInsert_NotFound
func Test_GetAndReplaceOrInsert_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, NEW, visitor.INSERT, t)
		return NEW, visitor.INSERTED
	}))
	value_, found := w.GetAndReplaceOrInsert(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
	value := value_.(int)
	if value != NEW {
		t.Fatalf("visit returned '%v' instead of '%d'", value, NEW)
	}
}

// Test_GetAndRemove_Found
func Test_GetAndRemove_Found(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, VALUE, true, nil, visitor.REMOVE, t)
		return VALUE, visitor.REMOVED
	}))
	value_, found := w.GetAndRemove(VALUE)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != VALUE {
		t.Fatalf("visit returned '%v' instead of '%d'", value, VALUE)
	}
}

// Test_GetAndRemove_NotFound
func Test_GetAndRemove_NotFound(t *testing.T) {
	const VALUE = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, nil, visitor.GET, t)
		return nil, visitor.NOT_FOUND
	}))
	value_, found := w.GetAndRemove(VALUE)
	if found == true {
		t.Fatalf("visit returned true")
	}
	if value_ != nil {
		t.Fatalf("visit returned '%v' instead of nil", value_)
	}
}

// Test_GetAndRemoveOrInsert_Found
func Test_GetAndRemoveOrInsert_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, nil, visitor.REMOVE, t)
		return OLD, visitor.REMOVED
	}))
	value_, found := w.GetAndRemoveOrInsert(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
	value := value_.(int)
	if value != OLD {
		t.Fatalf("visit returned '%v' instead of '%d'", value, OLD)
	}
}

// Test_GetAndRemoveOrInsert_NotFound
func Test_GetAndRemoveOrInsert_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, NEW, visitor.INSERT, t)
		return NEW, visitor.INSERTED
	}))
	value_, found := w.GetAndRemoveOrInsert(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
	value := value_.(int)
	if value != NEW {
		t.Fatalf("visit returned '%v' instead of '%d'", value, NEW)
	}
}

// Test_Replace_Found
func Test_Replace_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, NEW, visitor.REPLACE, t)
		return NEW, visitor.REPLACED
	}))
	found := w.Replace(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
}

// Test_Replace_NotFound
func Test_Replace_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, nil, visitor.GET, t)
		return nil, visitor.NOT_FOUND
	}))
	found := w.Replace(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
}

// Test_ReplaceOrInsert_Found
func Test_ReplaceOrInsert_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, NEW, visitor.REPLACE, t)
		return NEW, visitor.REPLACED
	}))
	found := w.ReplaceOrInsert(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
}

// Test_ReplaceOrInsert_NotFound
func Test_ReplaceOrInsert_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, NEW, visitor.INSERT, t)
		return NEW, visitor.INSERTED
	}))
	found := w.ReplaceOrInsert(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
}

// Test_Remove_Found
func Test_Remove_Found(t *testing.T) {
	const VALUE = 0
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, VALUE, true, nil, visitor.REMOVE, t)
		return VALUE, visitor.REMOVED
	}))
	found := w.Remove(VALUE)
	if found == false {
		t.Fatalf("visit returned false")
	}
}

// Test_Remove_NotFound
func Test_Remove_NotFound(t *testing.T) {
	const VALUE = 0
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, nil, visitor.GET, t)
		return nil, visitor.NOT_FOUND
	}))
	found := w.Remove(VALUE)
	if found == true {
		t.Fatalf("visit returned true")
	}
}

// Test_RemoveeOrInsert_Found
func Test_RemoveOrInsert_Found(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, OLD, true, nil, visitor.REMOVE, t)
		return OLD, visitor.REMOVED
	}))
	found := w.RemoveOrInsert(OLD, NEW)
	if found == false {
		t.Fatalf("visit returned false")
	}
}

// Test_RemoveOrInsert_NotFound
func Test_RemoveOrInsert_NotFound(t *testing.T) {
	const OLD = 0
	const NEW = 1
	w := visitor.New((f_visit)(func(key interface{}, f visitor.F) (interface{}, visitor.Result) {
		assertVisit(f, nil, false, NEW, visitor.INSERT, t)
		return NEW, visitor.INSERTED
	}))
	found := w.RemoveOrInsert(OLD, NEW)
	if found == true {
		t.Fatalf("visit returned true")
	}
}

// Test_Action_String
func Test_Action_String(t *testing.T) {
	assertActionString(visitor.INSERT /***/, "INSERT" /***/, t)
	assertActionString(visitor.GET /******/, "GET" /******/, t)
	assertActionString(visitor.REPLACE /**/, "REPLACE" /**/, t)
	assertActionString(visitor.REMOVE /***/, "REMOVE" /***/, t)

	assertActionString(visitor.Action(-1), "visitor.Action(-1)", t)
}

// Test_Result_String
func Test_Result_String(t *testing.T) {
	assertResultString(visitor.INSERTED /***/, "INSERTED" /***/, t)
	assertResultString(visitor.NOT_FOUND /**/, "NOT_FOUND" /**/, t)
	assertResultString(visitor.FOUND /******/, "FOUND" /******/, t)
	assertResultString(visitor.REPLACED /***/, "REPLACED" /***/, t)
	assertResultString(visitor.REMOVED /****/, "REMOVED" /****/, t)

	assertResultString(visitor.Result(-1), "visitor.Result(-1)", t)
}
