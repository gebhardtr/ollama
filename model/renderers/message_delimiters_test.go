package renderers

import "testing"

func TestMessageDelimitersForRenderer(t *testing.T) {
	for _, name := range []string{"qwen3.5", "lfm2", "lfm2-thinking", "nemotron-3.5-nano"} {
		got := MessageDelimitersForRenderer(name)
		if len(got) != 2 || got[0].Role != "user" || got[0].Delimiter != "<|im_start|>user" ||
			got[1].Role != "assistant" || got[1].Delimiter != "<|im_start|>assistant" {
			t.Errorf("%s: got %v, want the ChatML delimiters", name, got)
		}
	}

	for _, name := range []string{"", "unknown", "gemma4"} {
		if got := MessageDelimitersForRenderer(name); got != nil {
			t.Errorf("%s: got %v, want nil", name, got)
		}
	}
}
