//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what enrichModelsWithInheritance: 부모가 known 모델인 클래스 추가 등록(수렴)
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
	// seed globalModels with Base only; enrich should add Child then GrandChild
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

func TestEnrichModelsWithInheritance_NoChange(t *testing.T) {
	// no parents known -> loop exits immediately
	src := []byte(`class Standalone:
    pass
`)
	root, _ := parsePython(src)
	fi := fileInfo{absPath: "/s.py", src: src, root: root, models: map[string][]pydanticField{}}
	globalModels := map[string]*fileInfo{}
	enrichModelsWithInheritance([]fileInfo{fi}, globalModels)
	if len(globalModels) != 0 {
		t.Fatalf("expected no enrichment, got %v", globalModels)
	}
}
