package skipmap

import (
	"fmt"
	"testing"
)

func TestSkipMap(t *testing.T) {
	// Create a skip sm with int key.
	sm := New()

	// Add some values. Value can be anything.
	sm.Set(12, "hello world")
	sm.Set(33, 56)
	sm.Set(35, 56)
	sm.Set(78, 90.12)

	// Get element by index.
	n := sm.Get(33)     // Value is stored in n.Value.
	fmt.Println(n.elem) // Output: 56
	n.Val()
	//next := n.                  // Get next element.
	//prev := next.Prev()                 // Get previous element.
	//fmt.Println(next.Value, prev.Value) // Output: 90.12    56
	//
	//// Or, directly get value just like a map
	//val, ok := sm.GetValue(34)
	//fmt.Println(val, ok) // Output: 56  true
	//
	//// Find first elements with score greater or equal to key
	//foundElem := sm.Get(30)
	//fmt.Println(foundElem.Key(), foundElem.Value) // Output: 34 56

	// Remove an element for key.
	sm.Del(34)
}
