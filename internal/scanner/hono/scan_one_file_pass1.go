//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 단일 파일의 Pass 1: 파싱, Hono 인스턴스 수집, basePath, Zod 스키마, route 그룹 수집
package hono

import "github.com/park-jun-woo/codistill/internal/scanner/zod"

func scanOneFilePass1(path string) *pass1FileResult {
	fi, err := parseFile(path)
	if err != nil {
		return nil
	}
	vars := collectHonoVars(fi)
	bp := collectBasePaths(fi, vars)
	schemas := zod.CollectSchemas(fi.Root, fi.Src)
	groups := collectRouteGroups(fi, vars)
	for i := range groups {
		groups[i].SourceFile = path
	}
	return &pass1FileResult{fi: fi, vars: vars, bp: bp, schemas: schemas, groups: groups}
}
