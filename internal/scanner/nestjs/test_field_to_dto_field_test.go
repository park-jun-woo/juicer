//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFieldToDTOField 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestFieldToDTOField(t *testing.T) {
	f := scanner.Field{Name: "email", Type: "string", Validate: "required,email"}
	df := fieldToDTOField(f)
	if df.name != "email" || df.tsType != "string" || df.optional {
		t.Fatalf("got %+v", df)
	}

	df2 := fieldToDTOField(scanner.Field{Name: "x", Type: ""})
	if !df2.optional || df2.tsType != "string" {
		t.Fatalf("got %+v", df2)
	}
}
