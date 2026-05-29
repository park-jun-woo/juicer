//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what path parameter 슬라이스에서 이름으로 인덱스를 찾는다
package zod

import "github.com/park-jun-woo/codistill/internal/scanner"

func findPathParamIndex(params []scanner.Param, name string) int {
	for i, p := range params {
		if p.Name == name {
			return i
		}
	}
	return -1
}
