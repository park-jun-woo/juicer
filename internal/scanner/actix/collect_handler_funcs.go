//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what Rust 파일에서 모든 함수 정의를 수집하여 핸들러 인덱스에 등록한다
package actix

func collectHandlerFuncs(fi *fileInfo, index map[string]*handlerInfo) {
	for _, funcNode := range findAllByType(fi.root, "function_item") {
		nameNode := findChildByType(funcNode, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, fi.src)
		index[name] = &handlerInfo{
			funcNode: funcNode,
			file:     fi,
		}
	}
}
