//ff:func feature=scan type=extract control=sequence
//ff:what 타입정보 없는 응답에서 composite-literal body 타입명을 AST struct로 해석한다
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func fillFallbackResponseBody(resp *scanner.Response, idx *funcIndex) {
	if resp.Kind != "json" || resp.TypeName != "" || resp.Body == "" {
		return
	}
	typeName, base := compositeLitTypeName(resp.Body)
	if base == "" {
		return
	}
	st := idx.astStructs[base]
	if st == nil {
		return
	}
	fields := astStructFields(st)
	if len(fields) == 0 {
		return
	}
	resp.TypeName = typeName
	resp.Fields = fields
	resp.Confidence = "partial"
}
