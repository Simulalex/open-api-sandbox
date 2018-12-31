package path

import (
	"regexp"
)

type URLSubstitutor interface {
	ToRegex(swaggerPath string) string
}

type urlSubstitutor struct{}

// Replace all path strings with /pet/{petId}/uploadImage -> /pet/[^/\s]+/uploadImage
func (subs urlSubstitutor) ToRegex(swaggerPath string) string {
	re := regexp.MustCompile(`{(.*?)}`)
	return re.ReplaceAllString(swaggerPath, `[^/\s]+`)
}
