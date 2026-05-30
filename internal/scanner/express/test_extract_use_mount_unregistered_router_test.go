//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractUseMount_UnregisteredRouter 테스트
package express

import "testing"

func TestExtractUseMount_UnregisteredRouter(t *testing.T) {
	fi := mustParse(t, []byte(`other.use('/api', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
