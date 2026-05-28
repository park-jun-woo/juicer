//ff:type feature=scan type=model topic=express
//ff:what app.use() 마운트 정보 구조체
package express

type useMount struct {
	Prefix       string
	VarName      string
	FilePath     string
	SourceRouter string
}
