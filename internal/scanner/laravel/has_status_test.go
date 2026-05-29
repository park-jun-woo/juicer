//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what hasStatus — 응답 슬라이스에 지정 상태 코드가 존재하는지 검사한다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

func hasStatus(responses []scanner.Response, status string) bool {
	for _, r := range responses {
		if r.Status == status {
			return true
		}
	}
	return false
}
