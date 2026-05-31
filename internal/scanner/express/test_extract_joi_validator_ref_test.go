//ff:func feature=scan type=test topic=express control=sequence
//ff:what extractJoiValidatorRef validate(import.member) → joiValidatorRef 추출 테스트
package express

import "testing"

func TestExtractJoiValidatorRef(t *testing.T) {
	// valid: validate(schemas.userBody)
	fi := mustParse(t, []byte(`validate(schemas.userBody);`))
	call := findAllByType(fi.Root, "call_expression")[0]
	ref := extractJoiValidatorRef(call, fi.Src)
	if ref == nil || ref.ImportName != "schemas" || ref.Member != "userBody" {
		t.Fatalf("got %+v", ref)
	}
	// not a validate function
	fi = mustParse(t, []byte(`other(schemas.userBody);`))
	call = findAllByType(fi.Root, "call_expression")[0]
	if extractJoiValidatorRef(call, fi.Src) != nil {
		t.Error("non-validate should be nil")
	}
	// arg not a member_expression
	fi = mustParse(t, []byte(`validate(plain);`))
	call = findAllByType(fi.Root, "call_expression")[0]
	if extractJoiValidatorRef(call, fi.Src) != nil {
		t.Error("non-member arg should be nil")
	}
}
