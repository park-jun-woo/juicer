//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestEnrichModelsWithInheritance_NoChange 테스트
package fastapi

import "testing"

func TestEnrichModelsWithInheritance_NoChange(t *testing.T) {

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
