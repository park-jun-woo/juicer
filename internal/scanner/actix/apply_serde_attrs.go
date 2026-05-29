//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what serde 어트리뷰트(rename/default)를 적용해 JSON 이름과 nullable 여부를 결정한다
package actix

func applySerdeAttrs(attrs []serdeAttr, fieldName string, nullable bool) (string, bool) {
	jsonName := fieldName
	for _, a := range attrs {
		if a.rename != "" {
			jsonName = a.rename
		}
		if a.hasDefault {
			nullable = true
		}
	}
	return jsonName, nullable
}
