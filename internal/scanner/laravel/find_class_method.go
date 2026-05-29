//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 지정 클래스명에서 지정 메서드명의 method_declaration 노드를 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findClassMethod(fi *fileInfo, className, methodName string) *sitter.Node {
	for _, cls := range findAllByType(fi.root, "class_declaration") {
		if !classNameMatches(cls, fi.src, className) {
			continue
		}
		if m := findMethodInClass(cls, fi.src, methodName); m != nil {
			return m
		}
	}
	return nil
}
