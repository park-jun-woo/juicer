//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 메서드 본문의 abort/abort_if/abort_unless 호출 상태 코드를 응답으로 누적한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func appendAbortResponses(cm *controllerMethod, responses []scanner.Response) []scanner.Response {
	if cm.methodNode == nil {
		return responses
	}
	seen := statusSet(responses)
	for _, call := range findAllByType(cm.methodNode, "function_call_expression") {
		status := abortCallStatus(call, cm.src)
		if status == "" || seen[status] {
			continue
		}
		seen[status] = true
		responses = append(responses, scanner.Response{Status: status, Kind: "json"})
	}
	return responses
}
