//ff:func feature=scan type=test control=sequence
//ff:what groupCtx 테스트 헬퍼
package fiber

func groupCtx() *groupArgCtx {
	return &groupArgCtx{
		routers: map[string]*routerInfo{
			"api":       {prefix: "/api"},
			"authGroup": {prefix: "/auth"},
		},
	}
}
