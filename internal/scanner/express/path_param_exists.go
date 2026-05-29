//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 이름이 name인 path param이 슬라이스에 이미 존재하는지 판별한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func pathParamExists(params []scanner.Param, name string) bool {
	for _, p := range params {
		if p.Name == name {
			return true
		}
	}
	return false
}
