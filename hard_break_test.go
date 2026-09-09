package header

import "testing"

func TestBlockHardBreakRendering(t *testing.T) {
	text := func(s string) *Block { return &Block{Type: "text", Text: s} }
	br := func() *Block { return &Block{Type: "hard_break"} }
	for _, tc := range []struct {
		name  string
		block *Block
		want  string
	}{
		{"between text", &Block{Type: "div", Content: []*Block{text("111"), br(), text("222")}}, "111\n222"},
		{"consecutive breaks", &Block{Type: "div", Content: []*Block{text("111"), br(), br(), text("222")}}, "111\n\n222"},
		{"trim outer breaks", &Block{Type: "div", Content: []*Block{br(), text("111"), br(), text("222"), br()}}, "111\n222"},
		{"break only", &Block{Type: "div", Content: []*Block{br()}}, ""},
		{"list item", &Block{Type: "div", Content: []*Block{{Type: "bullet_list", Content: []*Block{{Type: "list_item", Content: []*Block{text("111"), br(), text("222")}}}}}}, "* 111\n222"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := BlockToText(tc.block); got != tc.want {
				t.Errorf("BlockToText=%q, want %q", got, tc.want)
			}
			if got, _ := BlockToTextF(tc.block, nil); got != tc.want {
				t.Errorf("BlockToTextF=%q, want %q", got, tc.want)
			}
			messages := BlockToTextMessages(tc.block)
			if tc.want == "" {
				if len(messages) != 0 {
					t.Errorf("blank block produced %v", messages)
				}
				return
			}
			if len(messages) != 1 || messages[0].Text != tc.want {
				t.Errorf("BlockToTextMessages=%v, want one message %q", messages, tc.want)
			}
		})
	}
}

func TestBlockHardBreakSpanOffset(t *testing.T) {
	marked := &Block{Type: "text", Text: "222"}
	block := &Block{Type: "div", Content: []*Block{{Type: "text", Text: "111"}, {Type: "hard_break"}, marked}}
	text, spans := BlockToTextF(block, func(b *Block) (string, bool) {
		if b == marked {
			return b.Text, true
		}
		return "", false
	})
	if text != "111\n222" || len(spans) != 1 {
		t.Fatalf("text=%q spans=%v", text, spans)
	}
	if spans[0].Start != 4 || spans[0].End != 7 {
		t.Fatalf("span=%+v, want [4,7)", spans[0])
	}
}
