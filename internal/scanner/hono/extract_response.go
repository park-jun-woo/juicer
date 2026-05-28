//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what c.json(), c.text(), c.body() 응답 호출을 분석하여 Response 슬라이스를 생성한다
package hono

import "github.com/park-jun-woo/codistill/internal/scanner"

func extractResponses(fi *fileInfo, routeLine int) []scanner.Response {
	var responses []scanner.Response
	calls := findAllByType(fi.Root, "call_expression")
	for _, call := range calls {
		r := extractOneResponse(call, fi.Src)
		if r != nil {
			responses = append(responses, *r)
		}
	}
	return responses
}
