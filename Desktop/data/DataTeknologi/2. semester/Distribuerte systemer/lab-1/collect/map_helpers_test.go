package collect

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "c": 3, "b": 2}
	expected := []string{"a", "b", "c"}

	result := keys(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("keys() = %v; want %v", result, expected)
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "c": 3, "b": 2}
	expected := []int{1, 2, 3}

	result := values(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("values() = %v; want %v", result, expected)
	}
}
