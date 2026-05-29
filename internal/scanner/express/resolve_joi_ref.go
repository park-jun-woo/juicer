//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what joiValidatorRef를 대상 파일의 const object로 해석하되 배럴 re-export 홉을 최대 5회 따라가 joi.RequestSchema를 반환한다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/joi"
)

// maxReexportHops — 배럴→배럴 체인을 따라가는 최대 홉 수 (사이클/무한루프 방지).
const maxReexportHops = 5

// resolveJoiRef — validate(importName.member) 참조를 해석한다.
// 1) 라우트 파일의 import 매핑에서 importName → 대상 파일 경로.
// 2) 대상 파일을 ctx.parsed에서 찾거나 신규 파싱.
// 3) 그 파일의 top-level `const <member> = {...}` object를 Joi 요청 스키마로 파싱.
// 3에서 const가 없으면, 배럴 re-export 한 홉을 따라가 대상 파일을 교체하고 재시도한다.
func resolveJoiRef(ref joiValidatorRef, fi *fileInfo, ctx *scanContext) joi.RequestSchema {
	imports := resolveImports(fi, ctx.absRoot, ctx.pathAliases)
	target := imports[ref.ImportName]
	if target == "" {
		return joi.RequestSchema{}
	}
	visited := map[string]bool{}
	for hop := 0; hop < maxReexportHops; hop++ {
		if target == "" || visited[target] {
			return joi.RequestSchema{}
		}
		visited[target] = true
		targetFi := loadParsedFile(ctx, target)
		if targetFi == nil {
			return joi.RequestSchema{}
		}
		if objNode := findConstObject(targetFi.Root, targetFi.Src, ref.Member); objNode != nil {
			return joi.ParseRequestObject(objNode, targetFi.Src)
		}
		target = resolveReexport(targetFi, ref.ImportName, ctx.absRoot, ctx.pathAliases)
	}
	return joi.RequestSchema{}
}
