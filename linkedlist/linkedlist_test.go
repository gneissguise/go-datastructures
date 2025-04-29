package linkedlist_test

import (
	"testing"

	"github.com/gneissguise/go-datastructures/linkedlist"
)

func createPopulatedList(t *testing.T, vals ...int) *linkedlist.LinkedList[int] {
	// Mark this as a test helper function
	t.Helper()
	// Initialize new list and append vals
	list := linkedlist.New[int]()
	for _, v := range vals {
		list.Append(v)
	}
	// Return the new list with appended vals
	return list
}

func TestNew(t *testing.T) {
	t.Log("TestNew started.")
	// Initialize new list
	list := linkedlist.New[string]()
	if list == nil {
		t.Fatal("Expected a new list, got nil")
	}
	if list.Length() != 0 {
		t.Error("Expected new list length 0, got %d", list.Length())
	}
	_, ok := list.Head()
	if !ok {
		t.Error("Expected head to be nil, got a non-nil value")
	}
	_, ok = list.Tail()
	if !ok {
		t.Error("Expected tail to be nil, got a non-nil value")
	}
	t.Log("TestNew finished.")
}

func TestAppend(t *testing.T) {
	t.Log("TestAppend started.")
	t.Run("Append a single item to a list", func(t *testing.T) {
		// Initialize new empty list
		list := linkedlist.New[int]()
		// Append a new int value of 10 to the list
		list.Append(10)

		if list.Length() != 1 {
			t.Fatalf("Expected length 1, got %d", list.Length())
		}
		headVal, ok := list.Head()
		if !ok || headVal != 10 {
			t.Errorf("Expected Head() to be 10 (ok=true), got %v (ok=%v)", headVal, ok)
		}
		tailVal, ok := list.Tail()
		if !ok || tailVal != 10 {
			t.Errorf("Expected Tail() to be 10 (ok=true), got %v (ok=%v)", tailVal, ok)
		}
	})

	t.Run("Append multiple items to a list", func(t *testing.T) {
		list := linkedlist.New[string]()
		list.Append("a")
		list.Append("b")
		list.Append("c")

		if list.Length() != 3 {
			t.Fatalf("Expected length 3, got %d", list.Length())
		}
		headVal, ok := list.Head()
		if !ok || headVal != "a" {
			t.Errorf("Expected Head() to be 'a' (ok=true), got %v (ok=%v)", headVal, ok)
		}
		tailVal, ok := list.Tail()
		if !ok || tailVal != "c" {
			t.Errorf("Expected Tail() to be 'c' (ok=true), got %v (ok=%v)", tailVal, ok)
		}

	})
	t.Log("TestAppend finished.")
}

func TestGet(t *testing.T) {
	t.Log("TestGet started.")
	// use helper to create a populated list
	list := createPopulatedList(t, 10, 20, 30)

	tests := []struct {
		name     string
		index    int
		expected int
		ok       bool
	}{
		{"Get the first item", 0, 10, true},
		{"Get the middle item", 1, 20, true},
		{"Get the last item", 2, 30, true},
		{"Get an out of bounds item (negative)", -1, 0, false},
		{"Get an out of bounds item (too large)", 3, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := list.Get(tt.index)
			// when Get() does not return ok..
			if ok != tt.ok {
				t.Errorf("Expected Get(%d) ok = %v, got ok = %v", tt.index, tt.ok, ok)
			}
			// Check the value if ok is expected to be true
			if tt.ok && val != tt.expected {
				t.Errorf("Expected Get(%d) val = %v, got val = %v", tt.index, tt.expected, val)
			}
		})
	}

	t.Run("Get from an empty list", func(t *testing.T) {
		list := linkedlist.New[int]()
		_, ok := list.Get(0)
		if ok {
			t.Error("Expected Get(0) on an empty list to return ok = false, got ok = true")
		}
	})

	t.Log("TestGet finished.")
}

func TestPrepend(t *testing.T) {
	t.Log("TestPrepend started.")

	// Test Prepend() on a populated list
	t.Run("Prepend value to  a populated list", func(t *testing.T) {
		list := createPopulatedList(t, 5, 10, 15)
		list.Prepend(1)
		if list.Length() != 4 {
			t.Fatalf("Expected length 4, got %d", list.Length())
		}
		// get first val on the list
		val, ok := list.Get(0)
		if !ok || val != 1 {
			t.Errorf("Expected first value to be 1 (ok=true), got %v (ok=%v)", val, ok)
		}
	})

	// Test Prepend() on an empty list
	t.Run("Prepend value to an empty list", func(t *testing.T) {
		list := linkedlist.New[int]()
		list.Prepend(101)
		if list.Length() != 1 {
			t.Fatalf("Expected length 1, got %d", list.Length())
		}
		val, ok := list.Get(0)
		if !ok || val != 101 {
			t.Errorf("Expected value to be 101 (ok=true), got %v (ok=%v)", val, ok)
		}
	})

	t.Log("TestPrepend finished.")
}

func TestInsertAt(t *testing.T) {
	t.Log("TestInsertAt started.")

	t.Run("Insert a value to a populated list", func(t *testing.T) {

	})

	t.Run("Insert a value to an empty list", func(t *testing.T) {

	})

	t.Run("Insert a value out of bounds (negative)", func(t *testing.T) {

	})

	t.Run("Insert a value out of bounds (too large)", func(t *testing.T) {

	})

	t.Log("TestInsertAt finished.")
}

func TestRemoveAt(t *testing.T) {
	t.Log("TestRemoveAt started.")

	t.Run("Remove a value at a location in a populated list", func(t *testing.T) {

	})

	t.Run("Remove a value at a location in an empty list", func(t *testing.T) {

	})

	t.Run("Remove a value out of bounds (negative)", func(t *testing.T) {

	})

	t.Run("Remove a value out of bounds (too large)", func(t *testing.T) {

	})

	t.Log("TestRemoveAt finished.")
}

func TestRemoveValue(t *testing.T) {
	t.Log("TestRemoveValue started.")

	t.Run("Remove a value from a populated list", func(t *testing.T) {

	})

	t.Run("Remove a value from an empty list", func(t *testing.T) {

	})

	t.Log("TestRemoveValue finished.")
}

func TestSearch(t *testing.T) {
	t.Log("TestSearch started.")

	t.Run("Search for a value in a populated list", func(t *testing.T) {

	})

	t.Run("Search for a value in an empty list", func(t *testing.T) {

	})

	t.Log("TestSearch finished.")
}

func TestReverse(t *testing.T) {
	t.Log("TestReverse started.")

	t.Run("Reverse a populated list", func(t *testing.T) {

	})

	t.Run("Reverse an empty list", func(t *testing.T) {

	})

	t.Log("TestReverse finished.")
}
