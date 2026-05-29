//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 한 파일의 struct 정의들을 타입명 인덱스에 등록한다
package actix

func indexFileStructs(fi *fileInfo, idx structIndex) {
	for _, structNode := range findAllByType(fi.root, "struct_item") {
		nameNode := findChildByType(structNode, "type_identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, fi.src)
		idx[name] = &structEntry{
			file:       fi,
			structName: name,
		}
	}
}
