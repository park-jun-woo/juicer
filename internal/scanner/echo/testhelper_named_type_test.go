//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what namedType 테스트 헬퍼
package echo

import (
	"go/types"
	"testing"
)

// namedType returns the *types.Named for the package-level var `name`.
func namedType(t *testing.T, info *types.Info, name string) types.Type {
	t.Helper()
	for id, obj := range info.Defs {
		if obj != nil && id.Name == name {
			return obj.Type()
		}
	}
	t.Fatalf("no type for %s", name)
	return nil
}
