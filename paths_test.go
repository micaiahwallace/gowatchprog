package gowatchprog

import "testing"

func TestSafeName(t *testing.T) {

	// Test name with non-ascii characters
	t1input := "Test name!@here123"
	t1 := &Program{Name: t1input}
	t1actual := t1.safeName()
	t1wanted := "Test-name--here123"
	if t1actual != t1wanted {
		t.Errorf("safeName(%s) = \"%s\"; Wanted \"%s\"", t1input, t1actual, t1wanted)
	}

	// Test with a valid name
	t2input := "ok-name-here"
	t2 := &Program{Name: t2input}
	t2actual := t2.safeName()
	t2wanted := "ok-name-here"
	if t2actual != t2wanted {
		t.Errorf("safeName(%s) = \"%s\"; Wanted \"%s\"", t2input, t2actual, t2wanted)
	}
}
