package quote

import "testing"

func TestSay2(t *testing.T) {
	want := "Hello"

	if got := Say2(); got != want {
		t.Errorf("Say2() = %q , want : %q", got, want)
	}
}

func TestSpeak2(t *testing.T) {
	want := "Hi, Mate"

	if got := Speak2(); got != want {
		t.Errorf("Speak() = %q , want : %q", got, want)
	}
}
