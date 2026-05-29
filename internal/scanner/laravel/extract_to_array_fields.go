//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 클래스의 toArray() 메서드를 찾아 반환 배열의 키를 필드로 추출한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractToArrayFields extracts field names from the toArray method's return array.
func extractToArrayFields(fi *fileInfo, className string) []scanner.Field {
	method := findClassMethod(fi, className, "toArray")
	if method == nil {
		return nil
	}
	return extractArrayKeys(method, fi.src)
}
