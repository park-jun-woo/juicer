//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryInheritAllFiles 테스트 (round5)
package fastapi

import "testing"

func TestTryInheritAllFiles_Round5(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n\nclass Child(Base):\n    name: str\n")
	root, _ := parsePython(src)
	fi := fileInfo{src: src, root: root, models: map[string][]pydanticField{}}
	files := []fileInfo{fi}
	globalModels := map[string]*fileInfo{"Base": &files[0]}

	if !tryInheritAllFiles(files, globalModels) {
		t.Fatal("expected Child to be added on first pass")
	}
	// second pass: nothing new
	if tryInheritAllFiles(files, globalModels) {
		t.Fatal("expected no additions on second pass")
	}
}

func TestTryInheritAllFiles_Empty_Round5(t *testing.T) {
	if tryInheritAllFiles(nil, map[string]*fileInfo{}) {
		t.Fatal("empty files should return false")
	}
}
