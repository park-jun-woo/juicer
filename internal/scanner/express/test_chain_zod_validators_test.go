//ff:func feature=scan type=test topic=express control=iteration dimension=1
//ff:what chainZodValidators 체인 메서드 인자에서 zod 검증 스키마 추출 테스트
package express

import "testing"

func TestChainZodValidators(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').post(validateRequest({ body: z.object({ name: z.string() }) }), handler);`))
	calls := findAllByType(fi.Root, "call_expression")
	var target = calls[0]
	for _, c := range calls {
		args := findChildByType(c, "arguments")
		if args != nil && args.NamedChildCount() >= 2 {
			target = c
			break
		}
	}
	infos := chainZodValidators(target, fi.Src)
	if len(infos) == 0 {
		t.Fatalf("expected zod validator info, got %+v", infos)
	}
}
