//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what 추출된 상태 코드 목록으로 Response 슬라이스를 구성한다
package supafunc

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildResponses(statuses []string) []scanner.Response {
	var responses []scanner.Response
	for _, s := range statuses {
		responses = append(responses, scanner.Response{Status: s, Kind: "json"})
	}
	return responses
}
