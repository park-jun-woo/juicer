//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindAllPydanticModels_SkipsNonModel 테스트
package fastapi

import "testing"

func TestFindAllPydanticModels_SkipsNonModel(t *testing.T) {
	src := []byte("class User(BaseModel):\n    id: int\nclass Plain:\n    x = 1\n")
	root, _ := parsePython(src)
	models := findAllPydanticModels(root, src)
	if _, ok := models["Plain"]; ok {
		t.Fatalf("Plain should be excluded: %v", models)
	}
	if _, ok := models["User"]; !ok {
		t.Fatalf("User should be present: %v", models)
	}
}
