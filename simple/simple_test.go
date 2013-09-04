package simple

import (
	"math/rand"
	"testing"
)

import (
	"github.com/iNamik/go_cmp"
)

/**********************************************************************
 ** Test Functions
 **********************************************************************/

// Test_Empty_True
func Test_Empty_True(t *testing.T) {
	r := New(cmp.F_int)
	assertEmpty(r, true, t)
}

// Test_Size_Empty
func Test_Size_Empty(t *testing.T) {
	r := New(cmp.F_int)
	assertSize(r, 0, t)
}

// Test_Insert
func Test_Insert(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
}

// Test_Empty_False
func Test_Empty_False(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertEmpty(r, false, t)
}

// Test_Size_One
func Test_Size_One(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertSize(r, 1, t)
}

// Test_Replace
func Test_Replace(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key1, key1, true, t)
}

// Test_Get_Empty
func Test_Get_Empty(t *testing.T) {
	r := New(cmp.F_int)
	assertGet(r, key1, nil, false, t)
}

// Test_Get_Found
func Test_Get_Found(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertGet(r, key1, key1, true, t)
}

// Test_Get_NotFound
func Test_Get_NotFound(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertGet(r, key2, nil, false, t)
}

// Test_Get_Replace
func Test_Get_Replace(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertGet(r, key1, key1, true, t)
	assertReplaceOrInsert(r, key1, key2, true, t)
	assertGet(r, key1, key2, true, t)
}

// Test_Size_Large_Insert
func Test_Size_Large_Insert(t *testing.T) {
	const SIZE = 1000
	r := randomTree(SIZE)
	assertSize(r, SIZE, t)
}

// Test_Remove_Empty
func Test_Remove_Empty(t *testing.T) {
	r := New(cmp.F_int)
	assertRemove(r, key1, false, t)
}

// Test_Remove_Found
func Test_Remove_Found(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertRemove(r, key1, true, t)
}

// Test_Remove_NotFound_Left
func Test_Remove_NotFound_Left(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertRemove(r, key1, false, t)
}

// Test_Remove_NotFound_Right
func Test_Remove_NotFound_Right(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertRemove(r, key2, false, t)
}

// Test_Size_Large_InsertRemove
func Test_Size_Large_InsertRemove(t *testing.T) {
	const SIZE = 1000
	r := randomTree(SIZE * 2)
	for _, i_ := range rand.Perm(SIZE) {
		i := i_ + i_
		assertRemove(r, i, true, t)
	}
	assertSize(r, SIZE, t)
}

// Test_Tree1
func Test_Tree1(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key5, key5, false, t)
	assertReplaceOrInsert(r, key4, key4, false, t)
	assertReplaceOrInsert(r, key3, key3, false, t)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key6, key6, false, t)
	assertReplaceOrInsert(r, key7, key7, false, t)
	assertReplaceOrInsert(r, key8, key8, false, t)
	assertReplaceOrInsert(r, key9, key9, false, t)
	assertBST(r, t)
}

// Test_Tree2
func Test_Tree2(t *testing.T) {
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key9, key9, false, t)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertReplaceOrInsert(r, key8, key8, false, t)
	assertReplaceOrInsert(r, key3, key3, false, t)
	assertReplaceOrInsert(r, key7, key7, false, t)
	assertReplaceOrInsert(r, key4, key4, false, t)
	assertReplaceOrInsert(r, key6, key6, false, t)
	assertReplaceOrInsert(r, key5, key5, false, t)
	assertBST(r, t)
}

// Test_Tree3
func Test_Tree3(t *testing.T) {
	var COMPARE = []int{-1, 1, -1}
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertRemove(r, key2, true, t)
	assertBST(r, t)
	assertCompare(r, COMPARE, t)
}

// Test_Tree4
func Test_Tree4(t *testing.T) {
	var COMPARE = []int{-1, 2, -1}
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertRemove(r, key1, true, t)
	assertBST(r, t)
	assertCompare(r, COMPARE, t)
}

// Test_Tree5
func Test_Tree5(t *testing.T) {
	var COMPARE = []int{-1, 1, 3, -1, 3, -1}
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key3, key3, false, t)
	assertRemove(r, key2, true, t)
	assertBST(r, t)
	assertCompare(r, COMPARE, t)
}

// Test_Tree6
func Test_Tree6(t *testing.T) {
	var COMPARE = []int{1, 3, -1, -1, 1, 2, -1, 2, -1}
	r := New(cmp.F_int)
	assertReplaceOrInsert(r, key4, key4, false, t)
	assertReplaceOrInsert(r, key1, key1, false, t)
	assertReplaceOrInsert(r, key3, key3, false, t)
	assertReplaceOrInsert(r, key2, key2, false, t)
	assertRemove(r, key4, true, t)
	assertBST(r, t)
	assertCompare(r, COMPARE, t)
}

// Test_Tree_Large_RandomInsertRemove
func Test_Tree_Large_RandomInsertRemove(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	const SIZE = 1000
	const COUNT = 1000000
	var array [SIZE]bool
	r := New(cmp.F_int)
	for i := 0; i < COUNT; i++ {
		n := rand.Intn(SIZE)
		_, exists := r.Get(n)
		// Should value already be in tree?
		if array[n] == false {
			if exists == true {
				t.Fatalf("(1)Tree contains %v when it shouldn't", n)
			}
			//fmt.Println("Add ", n)
			assertReplaceOrInsert(r, n, i, false, t)
			if _, exists = r.Get(n); exists == false {
				t.Fatalf("(1)Tree does not contain %v when it should", n)
			}
			array[n] = true
		} else {
			if exists == false {
				t.Fatalf("(2)Tree does not contain %v when it should", n)
			}
			//fmt.Println("Remove ", n)
			assertRemove(r, n, true, t)
			if _, exists = r.Get(n); exists == true {
				t.Fatalf("(2)Tree contains %v when it shouldn't", n)
			}
			array[n] = false
		}
	}
	assertBST(r, t)
}
