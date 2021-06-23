package turbo

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/google/go-cmp/cmp"
)

func TestTurboFrame(t *testing.T) {
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
		input    TurboFrameOptions
		expected string
	}{
		{
			name:     "TurboFrame: given default params should return skeleton frame",
			input:    TurboFrameOptions{},
			expected: `<turbo-frame id="" src="" loading="eager" target=""></turbo-frame>`,
		},
		{
			name: "TurboFrame: given an id should populate id",
			input: TurboFrameOptions{
				Id: "my-id",
			},
			expected: `<turbo-frame id="my-id" src="" loading="eager" target=""></turbo-frame>`,
		},
		{
			name: "TurboFrame: given a url should populate url",
			input: TurboFrameOptions{
				Src: "https://my.website.example/with/a/path?and=some&query=params",
			},
			expected: `<turbo-frame id="" src="https://my.website.example/with/a/path?and=some&amp;query=params" loading="eager" target=""></turbo-frame>`,
		},
		{
			name: "TurboFrame: given eager loading should render as eager",
			input: TurboFrameOptions{
				Loading: EagerLoading,
			},
			expected: `<turbo-frame id="" src="" loading="eager" target=""></turbo-frame>`,
		},
		{
			name: "TurboFrame: given lazy loading should render as lazy",
			input: TurboFrameOptions{
				Loading: LazyLoading,
			},
			expected: `<turbo-frame id="" src="" loading="lazy" target=""></turbo-frame>`,
		},
		{
			name: "TurboFrame: given target set should render target",
			input: TurboFrameOptions{
				Target: "a-target",
			},
			expected: `<turbo-frame id="" src="" loading="eager" target="a-target"></turbo-frame>`,
		},
		{
			name: "TurboFrame: given a child component should render child component",
			input: TurboFrameOptions{
				Contents: testChildComponent("Test Child"),
			},
			expected: `<turbo-frame id="" src="" loading="eager" target="">` +
				`Test Child` +
				`</turbo-frame>`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			w := new(strings.Builder)
			err := TurboFrame(tt.input).Render(context.Background(), w)
			if err != nil {
				t.Errorf("failed to render: %v", err)
			}
			if diff := cmp.Diff(tt.expected, w.String()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
