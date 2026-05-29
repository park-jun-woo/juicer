//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what serde 어트리뷰트 목록에 skip 표시가 있는지 확인한다
package actix

func serdeAttrsSkip(attrs []serdeAttr) bool {
	for _, a := range attrs {
		if a.skip {
			return true
		}
	}
	return false
}
