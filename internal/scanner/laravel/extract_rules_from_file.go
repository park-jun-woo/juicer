//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 클래스의 rules() 메서드를 찾아 필드 정의를 추출한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractRulesFromFile finds the rules() method in the class and extracts field definitions.
func extractRulesFromFile(fi *fileInfo, className string) []scanner.Field {
	method := findClassMethod(fi, className, "rules")
	if method == nil {
		return nil
	}
	return extractFieldsFromRulesMethod(method, fi.src)
}
