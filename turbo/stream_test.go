package turbo

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/google/go-cmp/cmp"
)

func TestStream(t *testing.T) {
	var testChildComponent = func(contents string) (t *templ.Component) {
		var c templ.Component

		c = templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			_, err = io.WriteString(w, contents)
			return err
		})

		return &c
	}

	var tests = []struct {
		name     string
		input    TurboStreamOptions
		expected string
	}{
		{
			name:     "TurboStream: given default params should return skeleton frame",
			input:    TurboStreamOptions{},
			expected: `<turbo-stream action="" target=""></turbo-stream>`,
		},
		{
			name: "TurboStream: given action should render action",
			input: TurboStreamOptions{
				Action: Replace,
			},
			expected: `<turbo-stream action="replace" target=""></turbo-stream>`,
		},
		{
			name: "TurboStream: given target should render target",
			input: TurboStreamOptions{
				Target: "target",
			},
			expected: `<turbo-stream action="" target="target"></turbo-stream>`,
		},
		{
			name: "TurboStream: given template should render template",
			input: TurboStreamOptions{
				Contents: testChildComponent("Some Template"),
			},
			expected: `<turbo-stream action="" target=""><template>Some Template</template></turbo-stream>`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			w := new(strings.Builder)
			err := TurboStream(tt.input).Render(context.Background(), w)
			if err != nil {
				t.Errorf("failed to render: %v", err)
			}
			if diff := cmp.Diff(tt.expected, w.String()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
