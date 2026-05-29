//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 핸들러명이 비었으면(클로저 등) method+path 기반 익명 id를 부여한다
package actix

// handlerOrAnon returns the resolved handler name, or a deterministic
// path-based anonymous id when the handler is anonymous (e.g. an inline
// closure). This preserves the route while preventing an empty handler from
// breaking extractor/scope-prefix matching downstream.
func handlerOrAnon(handler, method, path string) string {
	if handler != "" {
		return handler
	}
	return "anon:" + method + ":" + path
}
