//ff:func feature=scan type=extract control=sequence
//ff:what AST struct 필드의 `json:"..."` 태그에서 json 이름을 추출한다
package echo

import (
	"go/ast"
	"reflect"
	"strconv"
	"strings"
)

func astFieldJSONName(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}
	raw, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return ""
	}
	tag := reflect.StructTag(raw).Get("json")
	if tag == "" {
		return ""
	}
	name := strings.Split(tag, ",")[0]
	if name == "-" {
		return ""
	}
	return name
}
