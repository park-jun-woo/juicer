//ff:func feature=scan type=extract control=iteration dimension=2 topic=express
//ff:what 라우트 목록에서 ZodValidator의 스키마 변수명을 중복 없이 수집한다
package express

func collectSchemaNames(routes []routeInfo, seen map[string]bool, names *[]string) {
	for _, r := range routes {
		for _, v := range r.ZodValidators {
			if v.SchemaName != "" && !seen[v.SchemaName] {
				*names = append(*names, v.SchemaName)
				seen[v.SchemaName] = true
			}
		}
	}
}
