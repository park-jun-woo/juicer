//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 문자열이 Express 라우트 path 형태(`/`로 시작 또는 `*` catch-all)인지 판별한다
package express

func isRoutePathArg(s string) bool {
	return s == "*" || (len(s) > 0 && s[0] == '/')
}
