//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스와 메서드 역할을 병합한다
package quarkus

func mergeRoles(class, method []string) []string {
	if len(method) > 0 {
		return method
	}
	return class
}
