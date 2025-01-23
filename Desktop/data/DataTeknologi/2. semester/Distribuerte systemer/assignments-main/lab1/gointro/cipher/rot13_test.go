package cipher

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRot13(t *testing.T) {
	rot13Tests := []struct {
		in, want string
	}{
		{
			in:   "Go programming is fun.",
			want: "Tb cebtenzzvat vf sha.",
		},
		{
			in:   "Tb cebtenzzvat vf sha.",
			want: "Go programming is fun.",
		},
		{
			in:   "There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.",
			want: "Gurer ner gjb uneq guvatf va pbzchgre fpvrapr: pnpur vainyvqngvba, anzvat guvatf, naq bss-ol-bar reebef.",
		},
		{
			in:   "Gurer ner gjb uneq guvatf va pbzchgre fpvrapr: pnpur vainyvqngvba, anzvat guvatf, naq bss-ol-bar reebef.",
			want: "There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.",
		},
	}
	b := make([]byte, 1024)
	for _, test := range rot13Tests {
		r := rot13Reader{strings.NewReader(test.in)}
		n, err := r.Read(b)
		if err != nil {
			t.Errorf("rot13.Read(%q): got %v, expected EOF", test.in, err)
		}
		got := string(b[:n])
		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("rot13.Read(%q): (-want +got):\n%s", test.in, diff)
		}
	}
}
