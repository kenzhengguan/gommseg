package gommseg

import "testing"

func TestSegmentGetWord(t *testing.T) {
	text := "你好"
	_, ok := Ana.Get(text)

	if !ok {
		t.Errorf("shit happen\n")
	}
}

func TestSegmentWord(t *testing.T) {
	text := "你好"
	word, _ := Ana.Get(text)

	if word.Text != text {
		t.Errorf("expected %s, got %s", text, word.Text)
	}

	if word.Freq != 974 {
		t.Errorf("word freq should be 974, got %d", word.Freq)
	}
}

func TestSegmentMatchWordsLength(t *testing.T) {
	text := "学术危机"
	// text := []byte("你好吗")

	words := Ana.MatchWords(text)
	if len(words) != 2 {
		t.Errorf("words length should be 2, got %d, got %v", len(words), words)
	}
}

func TestSegmentMatchWords(t *testing.T) {
	text := "南京市长江大桥欢迎你"

	words := Ana.MatchWords(text)
	if words[0].Text != "南" {
		t.Errorf("match word error")
	}
}

func TestSegmentChunks(t *testing.T) {
	text := "南京市长江大桥欢迎你"
	chunks := Ana.Chunks(text)

	if len(chunks) != 16 {
		t.Errorf("expected 2, got %d", len(chunks))
	}
}

func TestSegmentChunksFilter(t *testing.T) {
	text := "南京市长江大桥欢迎你"
	chunks := Ana.Chunks(text)

	chunk := Ana.Filter(chunks)
	if chunk.Words[0].Text != "南京市" && chunk.Length() != 27 {
		t.Errorf("filter fail")
	}
}

func TestSegmentCut(t *testing.T) {
	text := "我们在野生动物园玩"
	Ana.Cut(text)
}

func BenchmarkSegment(b *testing.B) {
	text := "南京市长江大桥欢迎你"
	for n := 0; n < b.N; n++ {
		Ana.Cut(text)
	}
}
