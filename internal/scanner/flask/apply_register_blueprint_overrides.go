//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 내 register_blueprint 호출에서 prefix 오버라이드를 적용한다
package flask

// applyRegisterBlueprintOverrides scans a file for register_blueprint calls
// and applies url_prefix overrides to the prefix map.
func applyRegisterBlueprintOverrides(fi fileInfo, prefixes blueprintPrefix) {
	calls := findAllByType(fi.root, "call")
	for _, call := range calls {
		varName, overridePrefix := tryParseRegisterBlueprint(call, fi.src)
		if varName == "" {
			continue
		}
		if overridePrefix != "" {
			prefixes[varName] = overridePrefix
		}
	}
}
