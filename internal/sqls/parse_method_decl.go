//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 메서드 선언에서 SQL 스켈레톤을 추출한다
package sqls

import (
	"go/ast"
	"sort"
)

// parseMethodDecl extracts a MethodSkeleton from a single method declaration.
// Returns nil if the method is not a DB method.
func parseMethodDecl(fn *ast.FuncDecl) *MethodSkeleton {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return nil
	}

	repoName := receiverTypeName(fn.Recv.List[0].Type)
	if repoName == "" {
		return nil
	}

	crud := detectCRUD(fn.Body)
	if crud == "" {
		return nil
	}

	fragments := collectSQLFragments(fn.Body)

	crud = refineCRUDIfNeeded(crud, fragments, fn.Body)
	if crud == "" {
		return nil
	}

	tables := extractTables(fragments)
	inlineFragments := collectInlineSQLArgs(fn.Body)
	for _, frag := range inlineFragments {
		for _, t := range extractTablesFromSQL(frag) {
			tables = appendUnique(tables, t)
		}
	}
	if len(fragments) == 0 {
		fragments = inlineFragments
	}

	sort.Strings(tables)

	return &MethodSkeleton{
		Repo:         repoName,
		Method:       fn.Name.Name,
		CRUD:         crud,
		Tables:       tables,
		Params:       extractParams(fn.Type.Params),
		Returns:      extractReturns(fn.Type.Results),
		SQLFragments: fragments,
		Dynamic:      detectDynamic(fn.Body),
	}
}
