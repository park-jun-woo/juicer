//ff:func feature=scan type=convert control=selection topic=flask
//ff:what 분류된 키를 bodyFields의 form/json 슬롯에 반영한다
package flask

// applyBodyKey records a classified body key into bf. For "json" it also marks
// hasJSON true even when the key itself is empty.
func applyBodyKey(bf bodyFields, kind, key string) bodyFields {
	switch kind {
	case "form":
		bf.formFields = appendUnique(bf.formFields, key)
	case "json":
		bf.hasJSON = true
		bf.jsonFields = appendUnique(bf.jsonFields, key)
	}
	return bf
}
