//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what ROLE_ 프리픽스를 제거하여 역할명을 정규화한다
package spring

import "strings"

func normalizeRole(role string) string {
	return strings.TrimPrefix(role, "ROLE_")
}
