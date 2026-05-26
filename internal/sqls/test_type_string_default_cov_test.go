//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_DefaultCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_DefaultCov(t *testing.T) {
	typeString(&ast.CompositeLit{})
}
