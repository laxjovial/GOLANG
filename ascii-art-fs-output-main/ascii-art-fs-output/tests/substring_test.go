package tests

import (
	"reflect"
	"testing"

	"acad.learn2earn.ng/git/oelaho/ascii-art-color/master"
)

func TestSubstringIndexes_BasicMatch(t *testing.T) {
	result := master.SubstringIndexes("a king kitten have kit", "kit")
	expected := map[int]bool{7: true, 8: true, 9: true, 19: true, 20: true, 21: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SubstringIndexes() = %v, want %v", result, expected)
	}
}

func TestSubstringIndexes_NoMatch(t *testing.T) {
	result := master.SubstringIndexes("hello world", "xyz")
	if len(result) != 0 {
		t.Errorf("expected empty map, got %v", result)
	}
}

func TestSubstringIndexes_EmptySubstring(t *testing.T) {
	result := master.SubstringIndexes("hello", "")
	if len(result) != 0 {
		t.Errorf("expected empty map for empty substring, got %v", result)
	}
}

func TestSubstringIndexes_EmptyText(t *testing.T) {
	result := master.SubstringIndexes("", "kit")
	if len(result) != 0 {
		t.Errorf("expected empty map for empty text, got %v", result)
	}
}

func TestSubstringIndexes_BothEmpty(t *testing.T) {
	result := master.SubstringIndexes("", "")
	if len(result) != 0 {
		t.Errorf("expected empty map for both empty, got %v", result)
	}
}

func TestSubstringIndexes_FullMatch(t *testing.T) {
	result := master.SubstringIndexes("abc", "abc")
	expected := map[int]bool{0: true, 1: true, 2: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SubstringIndexes() = %v, want %v", result, expected)
	}
}

func TestSubstringIndexes_OverlappingMatches(t *testing.T) {
	result := master.SubstringIndexes("aaaa", "aa")
	expected := map[int]bool{0: true, 1: true, 2: true, 3: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SubstringIndexes() = %v, want %v", result, expected)
	}
}

func TestSubstringIndexes_SingleChar(t *testing.T) {
	result := master.SubstringIndexes("banana", "a")
	expected := map[int]bool{1: true, 3: true, 5: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SubstringIndexes() = %v, want %v", result, expected)
	}
}

func TestSubstringIndexes_SubstringLongerThanText(t *testing.T) {
	result := master.SubstringIndexes("hi", "hello")
	if len(result) != 0 {
		t.Errorf("expected empty map when substring longer than text, got %v", result)
	}
}

func TestAllIndexes_BasicLength(t *testing.T) {
	result := master.AllIndexes(5)
	expected := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AllIndexes(5) = %v, want %v", result, expected)
	}
}

func TestAllIndexes_ZeroLength(t *testing.T) {
	result := master.AllIndexes(0)
	if len(result) != 0 {
		t.Errorf("AllIndexes(0) expected empty map, got %v", result)
	}
}

func TestAllIndexes_LengthOne(t *testing.T) {
	result := master.AllIndexes(1)
	expected := map[int]bool{0: true}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AllIndexes(1) = %v, want %v", result, expected)
	}
}

func TestAllIndexes_ContainsAllKeys(t *testing.T) {
	n := 10
	result := master.AllIndexes(n)
	if len(result) != n {
		t.Errorf("AllIndexes(%d) returned map of length %d, want %d", n, len(result), n)
	}
	for i := 0; i < n; i++ {
		if !result[i] {
			t.Errorf("AllIndexes(%d) missing index %d", n, i)
		}
	}
}
