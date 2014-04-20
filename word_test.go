package gommseg

import "testing"

func TestNewWord(t *testing.T) {
	text := "test word"
	freq := 5
	word := NewWord(text, freq)

	if word.Text != text {
		t.Errorf("expected %s, got %s", text, word.Text)
	}

	if word.Freq != freq {
		t.Errorf("expected %d, got %d", freq, word.Freq)
	}
}
