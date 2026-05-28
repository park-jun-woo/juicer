//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 미해결 스키마 import를 추적하여 추가 파일을 파싱한다
package express

func resolveSchemaImports(ctx *scanContext) {
	neededSchemas := collectNeededSchemas(ctx)
	if len(neededSchemas) == 0 {
		return
	}
	unresolvedSet := buildUnresolvedSet(neededSchemas, ctx.schemas)
	if len(unresolvedSet) == 0 {
		return
	}
	for _, fi := range ctx.parsed {
		for _, stmt := range findAllByType(fi.Root, "import_statement") {
			resolveOneImportStmt(ctx, fi, stmt, unresolvedSet)
		}
	}
}
