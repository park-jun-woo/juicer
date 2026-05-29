//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 모든 파일의 struct 정의를 타입명 인덱스로 구축한다
package actix

func buildStructIndex(files []*fileInfo) structIndex {
	idx := make(structIndex)
	for _, fi := range files {
		indexFileStructs(fi, idx)
	}
	return idx
}
