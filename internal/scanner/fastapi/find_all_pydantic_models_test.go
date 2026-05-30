//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findAllPydanticModels 테스트
package fastapi

import "testing"

func TestFindAllPydanticModels(t *testing.T) {
	src := []byte("class User(BaseModel):\n    name: str\nclass Order(BaseModel):\n    total: int\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	models := findAllPydanticModels(root, src)
	if len(models) != 2 {
		t.Fatalf("expected 2 models, got %d", len(models))
	}
	if _, ok := models["User"]; !ok {
		t.Fatal("missing User")
	}
	if _, ok := models["Order"]; !ok {
		t.Fatal("missing Order")
	}
}

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
