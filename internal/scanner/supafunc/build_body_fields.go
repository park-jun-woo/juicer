//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what 필드 이름 목록을 scanner.Field 슬라이스로 변환한다
package supafunc

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildBodyFields(names []string) []scanner.Field {
	fields := make([]scanner.Field, len(names))
	for i, n := range names {
		fields[i] = scanner.Field{Name: n, Type: "unknown"}
	}
	return fields
}
