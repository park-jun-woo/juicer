//ff:type feature=scan type=model topic=actix
//ff:what serde 어트리뷰트 파싱 결과(rename/default/skip)
package actix

type serdeAttr struct {
	rename     string
	hasDefault bool
	skip       bool
}
