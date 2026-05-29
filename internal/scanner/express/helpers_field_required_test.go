//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 테스트 헬퍼: 필드 슬라이스에서 name 필드가 required로 표시됐는지 확인한다
package express

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func fieldRequired(fields []scanner.Field, name string) bool {
	for _, f := range fields {
		if f.Name == name && strings.Contains(","+f.Validate+",", ",required,") {
			return true
		}
	}
	return false
}
