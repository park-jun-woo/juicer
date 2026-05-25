//ff:func feature=sql type=parse control=sequence
//ff:what EXEC CRUD를 fragments와 AST에서 세분화한다
package sqls

import "go/ast"

// refineExecCRUD refines an EXEC crud from fragments and AST.
// Returns empty string if the method should be skipped.
func refineExecCRUD(fragments []string, body *ast.BlockStmt) string {
	crud := refineCRUD(fragments)
	if crud != "EXEC" {
		return crud
	}
	crud = refineCRUDFromAST(body)
	if crud == "EXEC" {
		return ""
	}
	return crud
}
