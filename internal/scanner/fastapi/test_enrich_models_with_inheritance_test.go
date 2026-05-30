//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestEnrichModelsWithInheritance 테스트
package fastapi

import "testing"

func TestEnrichModelsWithInheritance(t *testing.T) {
	src := []byte(`class Base(BaseModel):
    id: int

class Child(Base):
    name: str

class GrandChild(Child):
    extra: str
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := fileInfo{
		absPath: "/m.py",
		src:     src,
		root:    root,
		models:  map[string][]pydanticField{},
	}

	globalModels := map[string]*fileInfo{"Base": &fi}

	files := []fileInfo{fi}
	enrichModelsWithInheritance(files, globalModels)

	if _, ok := globalModels["Child"]; !ok {
		t.Fatalf("Child not enriched: %v keys", len(globalModels))
	}
	if _, ok := globalModels["GrandChild"]; !ok {
		t.Fatalf("GrandChild not enriched (convergence loop): %v", globalModels)
	}
}
