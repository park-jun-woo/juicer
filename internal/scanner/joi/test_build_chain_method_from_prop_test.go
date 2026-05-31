//ff:func feature=scan type=test control=sequence topic=joi
//ff:what buildChainMethodFromProp 프로퍼티명+문자열 인자로 ChainMethod 생성 테스트
package joi

import (
	"reflect"
	"testing"
)

func TestBuildChainMethodFromProp(t *testing.T) {
	root, src := parseJoiTS(t, `Joi.valid('A', 'B')`)
	call := firstOfType(root, "call_expression")
	fn := call.ChildByFieldName("function")
	prop := fn.ChildByFieldName("property")
	cm := buildChainMethodFromProp(call, prop, src)
	if cm.Name != "valid" {
		t.Errorf("name: %q", cm.Name)
	}
	if !reflect.DeepEqual(cm.Args, []string{"A", "B"}) {
		t.Errorf("args: %v", cm.Args)
	}
}
