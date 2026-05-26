//ff:func feature=scan type=convert control=selection topic=nestjs
//ff:what TypeScript 타입을 OpenAPI 타입으로 변환한다
package nestjs

import "strings"

// tsTypeToOpenAPI converts a TypeScript type string to an openAPIType.
func tsTypeToOpenAPI(ts string) openAPIType {
	ts = strings.TrimSpace(ts)
	ts = unwrapPromise(ts)

	switch ts {
	case "string":
		return openAPIType{Type: "string"}
	case "number":
		return openAPIType{Type: "number"}
	case "boolean":
		return openAPIType{Type: "boolean"}
	case "Date":
		return openAPIType{Type: "string", Format: "date-time"}
	case "Uuid":
		return openAPIType{Type: "string", Format: "uuid"}
	case "ObjectId":
		return openAPIType{Type: "string"}
	case "any":
		return openAPIType{Type: "object"}
	case "void", "":
		return openAPIType{}
	}

	if strings.HasPrefix(ts, "Array<") && strings.HasSuffix(ts, ">") {
		inner := strings.TrimSpace(ts[6 : len(ts)-1])
		return openAPIType{Type: "array", Items: inner}
	}
	if strings.HasSuffix(ts, "[]") {
		inner := strings.TrimSpace(ts[:len(ts)-2])
		return openAPIType{Type: "array", Items: inner}
	}
	if strings.HasPrefix(ts, "Record<") && strings.HasSuffix(ts, ">") {
		inner := ts[7 : len(ts)-1]
		parts := strings.SplitN(inner, ",", 2)
		if len(parts) == 2 {
			return openAPIType{Type: "object", Items: strings.TrimSpace(parts[1])}
		}
		return openAPIType{Type: "object"}
	}
	return openAPIType{Type: "object"}
}
