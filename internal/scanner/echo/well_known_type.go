//ff:func feature=scan type=extract control=selection
//ff:what Named 타입이 well-known 타입(time.Time, time.Duration, uuid.UUID)인지 확인한다
package echo

import "go/types"

func wellKnownType(named *types.Named) (string, bool) {
	pkg := named.Obj().Pkg()
	if pkg == nil {
		return "", false
	}
	fullName := pkg.Name() + "." + named.Obj().Name()
	switch fullName {
	case "time.Time", "time.Duration", "uuid.UUID":
		return fullName, true
	}
	return "", false
}
