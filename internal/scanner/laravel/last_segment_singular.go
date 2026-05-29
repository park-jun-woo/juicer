//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 경로형 리소스명("{product}/reviews")의 마지막 세그먼트만 단수화한다
package laravel

import "strings"

// lastSegmentSingular singularizes only the final path segment of a resource
// name. "{product}/reviews" -> "review" so nested resource parameters stay
// usable as a single path parameter name.
func lastSegmentSingular(s string) string {
	if i := strings.LastIndexByte(s, '/'); i >= 0 {
		s = s[i+1:]
	}
	return singularize(s)
}
