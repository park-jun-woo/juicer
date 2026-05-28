//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 단일 파일에서 MapGroup 호출을 수집한다
package dotnet

func extractMapGroupsFromFile(fi *fileInfo, groups map[string]string) {
	stmts := findAllByType(fi.root, "local_declaration_statement")
	for _, stmt := range stmts {
		matchMapGroupDeclaration(stmt, fi, groups)
	}
}
