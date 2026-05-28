//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what hasAnyRole 표현식 내부의 역할 문자열을 추출한다
package spring

func extractRolesFromExpr(inner string) []string {
	var roles []string
	for _, rm := range roleStringRegexp.FindAllStringSubmatch(inner, -1) {
		if len(rm) > 1 {
			roles = append(roles, normalizeRole(rm[1]))
		}
	}
	return roles
}
