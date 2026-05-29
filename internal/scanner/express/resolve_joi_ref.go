//ff:func feature=scan type=extract control=sequence topic=express
//ff:what joiValidatorRef를 대상 파일의 const object로 해석하여 joi.RequestSchema를 반환한다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/joi"
)

// resolveJoiRef — validate(importName.member) 참조를 해석한다.
// 1) 라우트 파일의 import 매핑에서 importName → 대상 파일 경로.
// 2) 대상 파일을 ctx.parsed에서 찾거나 신규 파싱.
// 3) 그 파일의 top-level `const <member> = {...}` object를 Joi 요청 스키마로 파싱.
func resolveJoiRef(ref joiValidatorRef, fi *fileInfo, ctx *scanContext) joi.RequestSchema {
	imports := resolveImports(fi, ctx.absRoot, ctx.pathAliases)
	target := imports[ref.ImportName]
	if target == "" {
		return joi.RequestSchema{}
	}
	targetFi := loadParsedFile(ctx, target)
	if targetFi == nil {
		return joi.RequestSchema{}
	}
	objNode := findConstObject(targetFi.Root, targetFi.Src, ref.Member)
	if objNode == nil {
		return joi.RequestSchema{}
	}
	return joi.ParseRequestObject(objNode, targetFi.Src)
}
