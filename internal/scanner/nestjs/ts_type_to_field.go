//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what TypeScript 타입을 scanner.Field로 변환한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// tsTypeToField converts a TypeScript type to a scanner.Field.
func tsTypeToField(name, tsType string, optional bool) scanner.Field {
	oa := tsTypeToOpenAPI(tsType)
	f := scanner.Field{
		Name: name,
		Type: oa.Type,
	}
	if oa.Format != "" {
		f.Type = oa.Type + ":" + oa.Format
	}
	if optional {
		f.Validate = "optional"
	}
	return f
}
