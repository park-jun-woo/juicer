//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryInheritClass 테스트
package fastapi

import "testing"

func TestTryInheritClass(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n\nclass Child(Base):\n    name: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root, models: map[string][]pydanticField{}}
	globalModels := map[string]*fileInfo{"Base": fi}

	child := classByName(root, src, "Child")
	if !tryInheritClass(child, fi, globalModels) {
		t.Fatal("Child should inherit from known Base")
	}
	if _, ok := fi.models["Child"]; !ok {
		t.Fatal("Child not registered in fi.models")
	}
	if _, ok := globalModels["Child"]; !ok {
		t.Fatal("Child not registered in globalModels")
	}

	// already registered -> false
	if tryInheritClass(child, fi, globalModels) {
		t.Fatal("already-registered class should return false")
	}

	// parent not known -> false
	base := classByName(root, src, "Base")
	delete(fi.models, "Base")
	if tryInheritClass(base, fi, globalModels) {
		t.Fatal("Base has no known parent (BaseModel) -> false")
	}
}
