//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what validate 문자열에 값을 추가한다
package dotnet

func appendValidate(existing, add string) string {
	if existing == "" {
		return add
	}
	return existing + "," + add
}
