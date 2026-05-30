//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what sampleFields 테스트 헬퍼
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

func sampleFields() []scanner.Field {
	return []scanner.Field{
		{Name: "name", Type: "string", Validate: "required"},
		{Name: "email", Type: "string", Validate: "required,email"},
		{Name: "age", Type: "number", Validate: "required"},
	}
}
