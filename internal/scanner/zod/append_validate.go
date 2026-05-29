//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what validate 문자열에 새 검증 규칙을 추가한다
package zod

func appendValidate(current, add string) string {
	if current == "" {
		return add
	}
	return current + "," + add
}
