//ff:func feature=scan type=convert control=iteration dimension=1 topic=laravel
//ff:what 응답 슬라이스의 상태 코드 집합을 만든다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

func statusSet(responses []scanner.Response) map[string]bool {
	seen := make(map[string]bool, len(responses))
	for _, r := range responses {
		seen[r.Status] = true
	}
	return seen
}
