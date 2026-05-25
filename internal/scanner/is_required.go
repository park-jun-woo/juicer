//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what validate/binding 태그에서 required 여부를 판별한다
package scanner

import (
	"strings"
)

func isRequired(f Field) bool {
	if f.Validate == "" {
		return false
	}
	for _, part := range strings.Split(f.Validate, ",") {
		if part == "required" {
			return true
		}
	}
	return false
}

