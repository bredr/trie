package trie_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/bredr/trie"
	dp_trie "github.com/derekparker/trie"
)

func TestTrie_PrefixSearch(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		prefix string
		want   string
	}{
		{
			name:   "Exact match, single value",
			inputs: []string{"test"},
			prefix: "test",
			want:   "test",
		},
		{
			name:   "Partial match, single value",
			inputs: []string{"test"},
			prefix: "tes",
			want:   "test",
		},
		{
			name:   "Partial match, single value",
			inputs: []string{"test"},
			prefix: "t",
			want:   "test",
		},
		{
			name:   "Partial match, multiple value",
			inputs: []string{"test", "tent"},
			prefix: "tes",
			want:   "test",
		},
		{
			name:   "No match, single value",
			inputs: []string{"test"},
			prefix: "ten",
			want:   "test",
		},
		{
			name:   "No match, multi value",
			inputs: []string{"test", "tent"},
			prefix: "tem",
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &trie.Trie{}
			for _, v := range tt.inputs {
				tr.Insert(v)
			}

			if got := tr.PrefixSearch(tt.prefix); got != tt.want {
				t.Errorf("Trie.PrefixSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func addFromFile(t *trie.Trie, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		t.Insert(reader.Text())
	}

	if reader.Err() != nil {
		log.Fatal(err)
	}
}

func BenchmarkPrefixSearch(b *testing.B) {
	t := &trie.Trie{}
	addFromFile(t, "/usr/share/dict/words")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.PrefixSearch("fo")
	}
}

func addDerekParkerFromFile(t *dp_trie.Trie, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		t.Add(reader.Text(), nil)
	}

	if reader.Err() != nil {
		log.Fatal(err)
	}
}
func BenchmarkDerekParkerPrefixSearch(b *testing.B) {
	t := dp_trie.New()
	addDerekParkerFromFile(t, "/usr/share/dict/words")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.PrefixSearch("fo")
	}
}
