//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestSubstituteOne_WordBoundary 단어 경계 기반 타입 치환 테스트
package nestjs

import "testing"

func TestSubstituteOne_WordBoundary(t *testing.T) {
	typeMap := map[string]string{"D": "Category"}

	// Should replace standalone "D"
	got := substituteOne("D", typeMap)
	if got != "Category" {
		t.Errorf("substituteOne(D) = %q, want Category", got)
	}

	// Should NOT replace "D" inside "UpdateDto"
	got = substituteOne("UpdateDto", typeMap)
	if got != "UpdateDto" {
		t.Errorf("substituteOne(UpdateDto) = %q, want UpdateDto (no change)", got)
	}

	// Should replace "D" in generic wrapper
	got = substituteOne("ResponseDto<D>", typeMap)
	if got != "ResponseDto<Category>" {
		t.Errorf("substituteOne(ResponseDto<D>) = %q, want ResponseDto<Category>", got)
	}
}
