//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 상속된 엔드포인트와 직접 정의된 엔드포인트를 병합한다
package nestjs

// mergeEndpoints combines inherited (parent) and direct (child) endpoints.
// If a child endpoint has the same handler name as a parent endpoint, the
// child's version takes precedence (override).
func mergeEndpoints(inherited, direct []endpointInfo) []endpointInfo {
	if len(inherited) == 0 {
		return direct
	}
	overrides := make(map[string]struct{}, len(direct))
	for _, ep := range direct {
		overrides[ep.handler] = struct{}{}
	}
	var merged []endpointInfo
	for _, ep := range inherited {
		if _, ok := overrides[ep.handler]; !ok {
			merged = append(merged, ep)
		}
	}
	merged = append(merged, direct...)
	return merged
}
