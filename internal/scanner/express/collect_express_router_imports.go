//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what import { Router } from 'express' 구문에서 Router alias 이름을 수집한다
package express

func collectExpressRouterImports(fi *fileInfo) map[string]bool {
	aliases := make(map[string]bool)
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		importPath := extractImportPath(stmt, fi.Src)
		if importPath != "express" {
			continue
		}
		collectRouterAliasesFromStmt(stmt, fi.Src, aliases)
	}
	return aliases
}
