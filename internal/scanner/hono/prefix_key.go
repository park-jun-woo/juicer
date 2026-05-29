//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what (파일경로, 변수명) 쌍을 prefixMap/basePaths 조회용 합성 키로 만든다
package hono

func prefixKey(file, varName string) string {
	return file + "\x00" + varName
}
