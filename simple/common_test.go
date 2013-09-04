package simple

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

import (
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Init
 **********************************************************************/

// init
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

/**********************************************************************
 ** Test Data
 **********************************************************************/

const key0 = 0
const key1 = 1
const key2 = 2
const key3 = 3
const key4 = 4
const key5 = 5
const key6 = 6
const key7 = 7
const key8 = 8
const key9 = 9

/**********************************************************************
 ** Assert Functions
 **********************************************************************/

// assertEmpty
func assertEmpty(r T, empty bool, t *testing.T) {
	if empty_ := r.Empty(); empty_ != empty {
		t.Fatalf("Empty() returned %v instead of %v", empty_, empty)
	}
}

// assertSize
func assertSize(r T, size int, t *testing.T) {
	if size_ := r.Size(); size_ != size {
		t.Fatalf("Size() returned %v instead of %v", size_, size)
	}
}

// assertReplaceOrInsert
func assertReplaceOrInsert(r T, key int, value interface{}, replaced bool, t *testing.T) {
	if replaced_ := r.ReplaceOrInsert(key, value); replaced_ != replaced {
		t.Fatalf("ReplaceOrInsert(%v) returned %v instead of %v", key, replaced_, replaced)
	}
}

// assertGet
func assertGet(r T, key int, value_ interface{}, found bool, t *testing.T) {
	v_, found_ := r.Get(key)
	if found_ != found {
		t.Fatalf("Get() returned %v", found_)
	}
	if found == true {
		if v_ != value_ {
			t.Fatalf("Get() returned value '%v' instead of '%v'", v_, value_)
		}
	}
}

// AssertRemove
func assertRemove(r T, key interface{}, removed bool, t *testing.T) {
	if removed_ := r.Remove(key); removed_ != removed {
		t.Fatalf("Remove(%v) returned %v instead of %v", key, removed_, removed)
	}
}

// assertKVF calls a func of type func()(key,value,found) and confirms the results
func assertKVF(key int, value_ interface{}, found bool, f func() (interface{}, interface{}, bool), t *testing.T) {
	k_, v_, found_ := f()
	if found_ != found {
		t.Fatalf("func returned %v", found_)
	}
	if found == true {
		k, ok := k_.(int)
		if ok == false {
			t.Fatal("func did not return key of type int")
		}
		if k != key {
			t.Fatalf("func returned key '%d' instead of '%d'", k, key)
		}
		if v_ != value_ {
			t.Fatalf("func returned value '%v' instead of '%v'", v_, value_)
		}
	}
}

// assertPanic
func assertPanic(t *testing.T, msg string, f func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("assertPanic: did not generate panic()")
		} else if r != msg {
			t.Fatalf("assertPanic: recover() recieved message '%s' instead of '%s'", r, msg)
		}
	}()
	f()
}

// assertBST
func assertBST(r T, t *testing.T) {
	if !r.(*tree).IsBST(t) {
		if r.Size() <= 100 {
			r.(*tree).Dump()
		}
		t.Fatalf("tree is not a BST")
	}
}

// assertCompare
func assertCompare(r T, data []int, t *testing.T) {
	if !r.(*tree).CompareArray(data, t) {
		r.(*tree).DumpArray()
		fmt.Println("--------- ")
		r.(*tree).Dump()
		t.Fatalf("tree is not correct")
	}
}

/**********************************************************************
 ** Helper Functions
 **********************************************************************/

// randomTree
func randomTree(n int) T {
	r := New(cmp.F_int)
	for _, i := range rand.Perm(n) {
		r.ReplaceOrInsert(i, i)
	}
	return r
}

// randomTreeDouble
func randomTreeDouble(n int) T {
	r := New(cmp.F_int)
	for _, i_ := range rand.Perm(n) {
		i := i_ + i_
		r.ReplaceOrInsert(i, i)
	}
	return r
}

/**********************************************************************
 ** Debug Functions
 **********************************************************************/

// Dump
func (r *tree) Dump() {
	dump(r.root)
}

// dump
func dump(x *node) {
	if x == nil {
		return
	}
	fmt.Printf("%3v <- %3v -> %3v\n", key(x.left), key(x), key(x.right))
	if x.left != nil && cmp.F_int(x.key, x.left.key) == cmp.GT {
		dump(x.left)
	}
	if x.right != nil && cmp.F_int(x.key, x.right.key) == cmp.LT {
		dump(x.right)
	}
}

// key
func key(x *node) int {
	if x == nil {
		return -1
	}
	return x.key.(int)
}

/**********************************************************************
 ** Integritry Functions
 **********************************************************************/

// IsEmpty
func (r *tree) IsEmpty(t *testing.T) bool {
	return r.root == nil
}

// IsBST
func (r *tree) IsBST(t *testing.T) (b bool) {
	b = false
	if r.root == nil {
		t.Error("Tree is empty")
	} else if !isBST(r.root, min(r.root).key, max(r.root).key) {
		t.Error("Tree is not a BST")
	} else {
		b = true
	}
	return
}

// isBST - Are all the values in the BST rooted at x between min and max,
// and does the same property hold for both subtrees?
func isBST(x *node, min interface{}, max interface{}) bool {
	if x == nil {
		return true
	}
	if cmp.F_int(x.key, min) == cmp.LT || cmp.F_int(max, x.key) == cmp.LT {
		return false
	}
	return isBST(x.left, min, x.key) && isBST(x.right, x.key, max)
}

// DumpArray
func (r *tree) DumpArray() {
	a := r.Array()
	fmt.Print("[")
	for i, n := range a {
		if i%3 == 0 {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(" [ ")
		} else if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(n)
		if (i+1)%3 == 0 {
			fmt.Print(" ]")
		}
	}
	fmt.Println(" ]")
}

// Array
func (r *tree) Array() []int {
	return array(r.root, make([]int, 0))
}

// array
func array(h *node, a []int) []int {
	if h == nil {
		return a
	}
	if h.left != nil && cmp.F_int(h.key, h.left.key) == cmp.GT {
		a = append(a, h.left.key.(int))
	} else {
		a = append(a, -1)
	}
	a = append(a, h.key.(int))
	if h.right != nil && cmp.F_int(h.key, h.right.key) == cmp.LT {
		a = append(a, h.right.key.(int))
	} else {
		a = append(a, -1)
	}
	a = array(h.left, a)
	a = array(h.right, a)
	return a
}

// CompareArray
func (r *tree) CompareArray(a2 []int, t *testing.T) bool {
	a1 := r.Array()
	if len(a1) != len(a2) {
		t.Errorf("Tree has size %v instead of %v", len(a1), len(a2))
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			t.Errorf("Tree[%v] has value %v instead of %v", i, a1[i], a2[i])
			return false
		}
	}
	return true
}
