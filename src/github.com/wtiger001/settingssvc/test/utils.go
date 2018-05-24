package test

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertNilE ...
func AssertNilE(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Fatalf("Error Occured, %s\n", err.Error())
}

// AssertNil ...
func AssertNil(t *testing.T, item interface{}) {
	if item != nil {
		return
	}
	t.Fatalf("Item is not Nil")
}

// AssertEqual ...
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	// Simple Equality
	if a == b {
		return
	}
	// Deep Equal
	if reflect.DeepEqual(a, b) {
		return
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
