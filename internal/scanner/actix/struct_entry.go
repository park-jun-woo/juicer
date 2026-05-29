//ff:type feature=scan type=model topic=actix
//ff:what struct 인덱스 엔트리(소속 파일과 struct 이름)
package actix

type structEntry struct {
	file       *fileInfo
	structName string
}
