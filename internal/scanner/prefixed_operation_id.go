//ff:func feature=scan type=extract control=sequence
//ff:what 중복 operationId에 경로 또는 method 접두를 선택해 붙인다 (doubling 방지)
package scanner

import "strings"

func prefixedOperationID(ep Endpoint, id string) string {
	prefix := firstPathSegment(ep.Path)
	// 경로 세그먼트가 id 선두 토큰과 같으면(doubling) method로 구분한다.
	if prefix == "" || prefix == leadingToken(id) {
		return methodPrefixedID(ep.Method, id)
	}
	return prefix + strings.ToUpper(id[:1]) + id[1:]
}
