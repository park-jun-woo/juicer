//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일에서 수집한 Zod 스키마를 scanContext에 병합한다
package express

func mergeFileSchemas(ctx *scanContext, fi *fileInfo) {
	schemas := collectZodSchemas(fi)
	for name, node := range schemas {
		ctx.schemas[name] = node
		ctx.schemaSrc[name] = fi.Src
	}
}
