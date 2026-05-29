//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 파일의 모든 클래스에서 지정 이름의 메서드를 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findAnyClassMethod(fi *fileInfo, methodName string) *sitter.Node {
	for _, cls := range findAllByType(fi.root, "class_declaration") {
		if m := findMethodInClass(cls, fi.src, methodName); m != nil {
			return m
		}
	}
	return nil
}
