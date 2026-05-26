//ff:func feature=scan type=extract control=sequence
//ff:what ConstToString 함수
package scanner

import (
	"fmt"
	"go/constant"
)

func ConstToString(v constant.Value) string {
	if v.Kind() == constant.Int {
		if i, ok := constant.Int64Val(v); ok {
			return fmt.Sprintf("%d", i)
		}
	}
	return v.ExactString()
}
