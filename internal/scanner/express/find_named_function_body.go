//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일 루트에서 이름이 name인 함수 정의를 찾아 본문을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func findNamedFunctionBody(fi *fileInfo, name string) *sitter.Node {
	root := fi.Root
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		if body := matchAnyDecl(child, fi.Src, name); body != nil {
			return body
		}
	}
	return nil
}
