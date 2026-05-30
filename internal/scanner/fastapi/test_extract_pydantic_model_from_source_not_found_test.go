//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModelFromSource_NotFound 테스트
package fastapi

import "testing"

func TestExtractPydanticModelFromSource_NotFound(t *testing.T) {
	src := []byte("class Other(BaseModel):\n    x: int\n")
	fields, err := extractPydanticModelFromSource(src, "Missing")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 0 {
		t.Fatalf("expected no fields for missing class, got %+v", fields)
	}

}
