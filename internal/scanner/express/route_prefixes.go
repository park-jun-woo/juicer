//ff:func feature=scan type=extract control=sequence topic=express
//ff:what (file, routerVar) 라우터 인스턴스의 prefix 목록을 반환한다(없으면 "" 하나)
package express

// routePrefixes는 (file, routerVar) 라우터 인스턴스의 prefix 목록을 반환한다.
// 마운트가 없으면(루트 라우터/미해석) prefix "" 하나로 처리한다.
func routePrefixes(ctx *scanContext, file, routerVar string) []string {
	if ps := ctx.routerPrefixes[routerKey{file, routerVar}]; len(ps) > 0 {
		return ps
	}
	return []string{""}
}
