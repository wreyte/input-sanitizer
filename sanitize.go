package sanitize

import (
	"html"

	"github.com/microcosm-cc/bluemonday"
)

type Sanitize struct{}

func (s Sanitize) SanitizeContent(content string) string {
	policy := bluemonday.UGCPolicy()
	sanitized := policy.Sanitize(content)

	unescaped := html.UnescapeString(sanitized)

	return unescaped
}
