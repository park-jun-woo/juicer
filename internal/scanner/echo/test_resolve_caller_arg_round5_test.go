//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestResolveCallerArg_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestResolveCallerArg_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
type Ctx interface{ JSON() }
type Dto struct { Name string `+"`json:\"name\"`"+` }
var d Dto
var n int
var _ = d
var _ = n
`)
	// int-kind param -> status branch
	var intType, ifaceType types.Type
	for id, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if id.Name == "n" {
			intType = obj.Type()
		}
	}

	if intType == nil {
		intType = types.Typ[types.Int]
	}
	res := resolveCallerArg(intType, parseExpr(t, "200"), info)
	if res.status == "" && res.typeName == "" && !res.skip {

	}
	_ = ifaceType

	empty := types.NewInterfaceType(nil, nil).Complete()
	_ = resolveCallerArg(empty, parseExpr(t, "d"), info)
}
