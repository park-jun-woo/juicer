//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 배열 타입 필드에 items 하위 필드를 부착한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func attachArrayItems(f *scanner.Field, oaType openAPIType) {
	if oaType.Type != "array" || oaType.Items == "" {
		return
	}
	itemType := rustTypeToOpenAPI(oaType.Items)
	f.Fields = []scanner.Field{{
		Name: "items",
		Type: itemType.Type,
	}}
}
