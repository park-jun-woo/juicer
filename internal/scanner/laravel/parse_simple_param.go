//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what simple_parameter 노드에서 파라미터 타입명과 변수명을 파싱한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseSimpleParam(sp *sitter.Node, src []byte) methodParam {
	mp := methodParam{
		typeName: paramTypeName(sp, src),
	}
	if varName := findChildByType(sp, "variable_name"); varName != nil {
		mp.name = strings.TrimPrefix(nodeText(varName, src), "$")
	}
	return mp
}
