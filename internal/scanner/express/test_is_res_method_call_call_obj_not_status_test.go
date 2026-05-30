//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_CallObjNotStatus 테스트
package express

import "testing"

func TestIsResMethodCall_CallObjNotStatus(t *testing.T) {

	fi := mustParse(t, []byte(`foo().json({});`))
	if _, ok := isResMethodCall(outermostCall(fi), fi.Src); ok {
		t.Fatal("expected false")
	}
}
