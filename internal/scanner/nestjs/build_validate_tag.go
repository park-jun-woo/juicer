//ff:func feature=scan type=convert control=selection topic=nestjs
//ff:what optional 여부와 validators로 validate 태그 문자열을 조합한다
package nestjs

import "strings"

// buildValidateTag builds the validate tag string from optional flag and validators.
func buildValidateTag(optional bool, validators []string) string {
	switch {
	case !optional && len(validators) > 0:
		return "required," + strings.Join(validators, ",")
	case !optional:
		return "required"
	case len(validators) > 0:
		return strings.Join(validators, ",")
	default:
		return ""
	}
}
