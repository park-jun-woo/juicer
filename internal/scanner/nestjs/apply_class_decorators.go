//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 클래스 레벨 데코레이터를 각 메서드 엔드포인트에 전파한다
package nestjs

// applyClassDecorators propagates class-level @UseGuards and @Roles
// to each endpoint in the controller. Class-level guards are prepended
// before method-level guards; class-level roles are prepended before
// method-level roles.
func applyClassDecorators(ci *controllerInfo) {
	if len(ci.classMiddleware) == 0 && len(ci.classRoles) == 0 {
		return
	}
	for i := range ci.endpoints {
		if len(ci.classMiddleware) > 0 {
			merged := make([]string, 0, len(ci.classMiddleware)+len(ci.endpoints[i].middleware))
			merged = append(merged, ci.classMiddleware...)
			merged = append(merged, ci.endpoints[i].middleware...)
			ci.endpoints[i].middleware = merged
		}
		if len(ci.classRoles) > 0 {
			merged := make([]string, 0, len(ci.classRoles)+len(ci.endpoints[i].roles))
			merged = append(merged, ci.classRoles...)
			merged = append(merged, ci.endpoints[i].roles...)
			ci.endpoints[i].roles = merged
		}
	}
}
