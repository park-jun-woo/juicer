//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 크로스파일 마운트에서 자식 파일의 어느 라우터 변수에 prefix를 붙일지 결정한다
package express

// resolveChildVar는 부모 파일에서 alias로 import해 마운트한 라우터가
// 자식 파일(childFile)의 어떤 라우터 변수인지 해석한다.
//   - 자식 파일에 라우터가 하나면 그 변수 (default export 단일 라우터 패턴)
//   - 여럿이면 alias와 같은 이름의 변수 (named export 패턴: alias == export명)
//   - 그 외(모호) → "" 를 반환하고 호출부가 파일 내 전체 라우터에 적용한다
func resolveChildVar(childFile, alias string, allRouters map[string]map[string]bool) string {
	rs := allRouters[childFile]
	if len(rs) == 0 {
		return ""
	}
	if len(rs) == 1 {
		for v := range rs {
			return v
		}
	}
	if rs[alias] {
		return alias
	}
	return ""
}
