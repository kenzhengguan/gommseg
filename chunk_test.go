package gommseg

import "testing"

func TestNewChunk(t *testing.T) {
	words := []*Word{
		{"test", 4},
		{"word", 5},
	}
	chunk := NewChunk(words)

	if chunk.Length() != 8 {
		t.Errorf("expected 8, got %d", chunk.Length())
	}
	if chunk.AverageLength() != 4 {
		t.Errorf("expected 4, got %d", chunk.AverageLength())
	}
	if chunk.Variance() != 0 {
		t.Errorf("expected 0, got %d", chunk.Variance())
	}
	if chunk.Freq() != 9 {
		t.Errorf("expected 9, got %d", chunk.Freq())
	}
}
