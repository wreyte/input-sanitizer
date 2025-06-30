package sanitize

import (
	"github.com/microcosm-cc/bluemonday"
)

type Sanitize struct{}

func (s Sanitize) SanitizeContent(content string) string {
	policy := bluemonday.UGCPolicy()

	policy.AllowElements("p", "br", "ul", "ol", "li", "strong", "em", "b", "i", "u", "a")
	policy.AllowAttrs("href").OnElements("a")

	sanitized := policy.Sanitize(content)
	return sanitized
}
