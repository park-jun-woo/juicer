//ff:func feature=scan type=convert control=iteration dimension=1 topic=spring
//ff:what 필드 배열의 Java 타입을 scanner 타입으로 변환한다
package spring

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func convertFieldTypes(fields []scanner.Field) []scanner.Field {
	for i := range fields {
		if fields[i].Type != "" && !strings.Contains(fields[i].Type, ":") || isJavaType(fields[i].Type) {
			fields[i].Type = fieldTypeToScannerType(fields[i].Type)
		}
	}
	return fields
}
