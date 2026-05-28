//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 핸들러 본문에서 모든 res.json()/res.send()/res.sendStatus() 응답을 추출한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

// extractResponses — 핸들러 본문에서 모든 Response를 수집한다.
// 본문을 찾지 못하면 기본 응답 [{200, json}]을 반환한다.
func extractResponses(fi *fileInfo, ri routeInfo) []scanner.Response {
	body := findHandlerBody(fi, ri)
	if body == nil {
		return []scanner.Response{{Status: "200", Kind: "json"}}
	}

	calls := findAllByType(body, "call_expression")
	seen := map[string]bool{}
	var responses []scanner.Response

	for _, call := range calls {
		r := extractOneResponse(call, fi.Src)
		if r == nil {
			continue
		}
		key := r.Status + "/" + r.Kind
		if seen[key] {
			continue
		}
		seen[key] = true
		responses = append(responses, *r)
	}

	if len(responses) == 0 {
		return []scanner.Response{{Status: "200", Kind: "json"}}
	}
	return responses
}
