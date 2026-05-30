//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModelFromSource_Found 테스트
package fastapi

import "testing"

func TestExtractPydanticModelFromSource_Found(t *testing.T) {
	src := []byte("class User(BaseModel):\n    id: int\n")
	fields, err := extractPydanticModelFromSource(src, "User")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].Name != "id" {
		t.Fatalf("got %+v", fields)
	}
}
