//ff:func feature=sql type=parse control=sequence
//ff:what EXEC/SELECT CRUD를 SQL 내용에 따라 세분화한다
package sqls

import "go/ast"

// refineCRUDIfNeeded refines CRUD based on SQL fragment content and AST.
// Returns empty string if the method should be skipped.
func refineCRUDIfNeeded(crud string, fragments []string, body *ast.BlockStmt) string {
	if crud == "EXEC" {
		return refineExecCRUD(fragments, body)
	}

	if crud == "SELECT" {
		refined := refineCRUD(fragments)
		if refined == "INSERT" || refined == "UPDATE" || refined == "DELETE" {
			return refined
		}
	}

	return crud
}
