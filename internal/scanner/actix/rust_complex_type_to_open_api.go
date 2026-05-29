//ff:func feature=scan type=convert control=sequence topic=actix
//ff:what Vec/HashMap/Option 등 제네릭 Rust 타입을 OpenAPI 타입으로 변환한다
package actix

import "strings"

func rustComplexTypeToOpenAPI(rtype string) openAPIType {
	if strings.HasPrefix(rtype, "Vec<") {
		return openAPIType{Type: "array", Items: extractGenericInner(rtype)}
	}
	if strings.HasPrefix(rtype, "HashMap<") || strings.HasPrefix(rtype, "BTreeMap<") {
		return openAPIType{Type: "object"}
	}
	if strings.HasPrefix(rtype, "Option<") {
		return rustTypeToOpenAPI(extractGenericInner(rtype))
	}
	return openAPIType{Type: "object"}
}
