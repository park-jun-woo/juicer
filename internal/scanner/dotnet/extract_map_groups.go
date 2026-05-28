//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what app.MapGroup("/prefix") 호출에서 변수명과 prefix를 수집한다
package dotnet

func extractMapGroups(files []*fileInfo) map[string]string {
	groups := make(map[string]string)
	for _, fi := range files {
		extractMapGroupsFromFile(fi, groups)
	}
	return groups
}
