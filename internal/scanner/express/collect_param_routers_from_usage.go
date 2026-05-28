//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what HTTP 메서드 호출 패턴에서 추정 라우터 변수명을 수집한다 (타입 어노테이션 폴백)
package express

func collectParamRoutersFromUsage(fi *fileInfo, routers map[string]bool) {
	httpMethods := map[string]bool{
		"get": true, "post": true, "put": true, "delete": true,
		"patch": true, "options": true, "head": true, "all": true,
		"use": true,
	}
	for _, call := range findAllByType(fi.Root, "call_expression") {
		mem := findChildByType(call, "member_expression")
		if mem == nil {
			continue
		}
		prop := mem.ChildByFieldName("property")
		if prop == nil {
			continue
		}
		methodName := nodeText(prop, fi.Src)
		if !httpMethods[methodName] {
			continue
		}
		obj := findChildByType(mem, "identifier")
		if obj == nil {
			continue
		}
		varName := nodeText(obj, fi.Src)
		if varName == "express" || varName == "module" || varName == "exports" {
			continue
		}
		args := findChildByType(call, "arguments")
		if args == nil || args.NamedChildCount() == 0 {
			continue
		}
		firstArg := args.NamedChild(0)
		if firstArg != nil && firstArg.Type() == "string" {
			routers[varName] = true
		}
	}
}
