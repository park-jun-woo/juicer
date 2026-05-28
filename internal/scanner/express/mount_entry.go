//ff:type feature=scan type=model topic=express
//ff:what 라우터 마운트 엔트리 구조체
package express

type mountEntry struct {
	prefix     string
	varName    string
	filePath   string
	sourceFile string
}
