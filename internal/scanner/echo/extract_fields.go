//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what types.Struct에서 필드 목록을 추출한다 (임베딩은 재귀 전개)
package echo

import (
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractFields(st *types.Struct, visited map[string]bool) []scanner.Field {
	var fields []scanner.Field

	for i := 0; i < st.NumFields(); i++ {
		f := st.Field(i)
		tag := st.Tag(i)

		if f.Embedded() {
			embeddedFields := resolveEmbedded(f.Type(), visited)
			fields = append(fields, embeddedFields...)
			continue
		}

		if field := buildField(f, tag, visited); field != nil {
			fields = append(fields, *field)
		}
	}

	return fields
}
