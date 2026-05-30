//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what defTypeByName 테스트 헬퍼: types.Info.Defs에서 이름으로 타입 조회
package echo

import (
	"go/types"
)

// defTypeByName returns the type of the first definition matching name.
func defTypeByName(info *types.Info, name string) types.Type {
	for id, obj := range info.Defs {
		if obj != nil && id.Name == name {
			return obj.Type()
		}
	}
	return nil
}
