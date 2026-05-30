//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestDeclaratorMatchesName_Round5 테스트
package express

import "testing"

func TestDeclaratorMatchesName_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const router = express.Router();`))
	vd := exFirst(t, fi, "variable_declarator")
	if !declaratorMatchesName(vd, fi.Src, "router") {
		t.Fatal("should match router")
	}
	if declaratorMatchesName(vd, fi.Src, "other") {
		t.Fatal("should not match other")
	}
}
