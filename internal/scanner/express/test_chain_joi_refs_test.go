//ff:func feature=scan type=test topic=express control=iteration dimension=1
//ff:what chainJoiRefs 체인 메서드 인자에서 validate(import.member) 참조 수집 테스트
package express

import "testing"

func TestChainJoiRefs(t *testing.T) {
	// router.route('/x').post(validate(schemas.body), handler)
	fi := mustParse(t, []byte(`router.route('/x').post(validate(schemas.body), handler);`))
	calls := findAllByType(fi.Root, "call_expression")
	// the .post(...) call has 2 args (validate(...) + handler)
	var target = calls[0]
	for _, c := range calls {
		args := findChildByType(c, "arguments")
		if args != nil && args.NamedChildCount() >= 2 {
			target = c
			break
		}
	}
	refs := chainJoiRefs(target, fi.Src)
	if len(refs) != 1 || refs[0].ImportName != "schemas" || refs[0].Member != "body" {
		t.Fatalf("got %+v", refs)
	}
}
