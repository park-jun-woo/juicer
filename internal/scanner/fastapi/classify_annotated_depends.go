//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Annotated Depends 타입을 form body 또는 미들웨어로 분류한다
package fastapi

// formDataClasses maps FastAPI form-data binding class names to true.
// When Depends() is called with no arg, the Annotated first type becomes
// the dependency class. If it is a form-data class, classify as body(form).
var formDataClasses = map[string]bool{
	"OAuth2PasswordRequestForm": true,
}

// classifyAnnotatedDepends resolves an Annotated Depends type and classifies
// it as either body(form) or middleware depending on the resolved name.
func classifyAnnotatedDepends(name, typeName string, aliasMap map[string]string, ri *routeInfo) {
	if aliasMap != nil {
		if fn, ok := aliasMap[typeName]; ok {
			ri.middleware = append(ri.middleware, fn)
			return
		}
	}
	fn := extractDependsFromAnnotated(typeName)
	if fn == "" {
		return
	}
	if formDataClasses[fn] {
		ri.bodyType = fn
		ri.bodyVarName = name
		return
	}
	ri.middleware = append(ri.middleware, fn)
}
