//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what 타입명에서 nullable 접미사를 제거하고 nullable 여부를 반환한다
package dotnet

func stripNullable(typeName string) (string, bool) {
	if len(typeName) > 0 && typeName[len(typeName)-1] == '?' {
		return typeName[:len(typeName)-1], true
	}
	return typeName, false
}
