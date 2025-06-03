package sanitize

import (
	"strings"
	"testing"
)

func TestSanitizeContent(t *testing.T) {
	s := Sanitize{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"RemoveTheScriptTag",
			`<script>alert("Hello!")</script>Safe content`,
			`Safe content`,
		},
		{
			"RemoveTheStyleTag",
			`<style>body { background-color: red; }</style>Safe content`,
			`Safe content`,
		},
		{
			"RemoveTheIframeTag",
			`<iframe src="https://example.com"></iframe>Safe content`,
			`Safe content`,
		},
		{
			"RemoveTheObjectTag",
			`<object data="http://example.com"></object>Safe content`,
			`Safe content`,
		},
		{
			"RemoveTheEmbedTag",
			`<embed src="http://example.com" />Safe content`,
			`Safe content`,
		},
		{
			"RemoveTheBaseTag",
			`<base href="https://example.com/" />Safe content`,
			`Safe content`,
		},
		{
			"PreserveTheGoCode",
			`package main\n\nimport "fmt"\n\nfunc main() {\n\tfmt.Println("Hello!")\n}`,
			`package main\n\nimport "fmt"\n\nfunc main() {\n\tfmt.Println("Hello!")\n}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.SanitizeContent(tt.input)

			if output != tt.expected {
				t.Errorf("Test %s failed. Expected: %s, Got: %s", tt.name, tt.expected, output)
			}

			for _, tag := range []string{"script", "style", "iframe", "object", "embed", "base"} {
				if strings.Contains(output, "<"+tag) || strings.Contains(output, "</"+tag) {
					t.Errorf("Test %s failed. Sanitized content should not contain <%s> tag",
						tt.name, tag)
				}
			}
		})
	}
}
