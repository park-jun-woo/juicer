//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 경로 파라미터 이름 슬라이스를 scanner.Param 슬라이스로 변환한다
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildPathParams(params []string) []scanner.Param {
	var result []scanner.Param
	for _, p := range params {
		result = append(result, scanner.Param{Name: p, Type: "string"})
	}
	return result
}
