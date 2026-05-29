//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 내 register_blueprint 호출에서 prefix 오버라이드를 적용한다
package flask

// applyRegisterBlueprintOverrides scans a file for register_blueprint calls
// and applies url_prefix overrides to the prefix map. The register argument name
// is resolved back to the canonical Blueprint variable via the file's import
// aliases, so app-factory patterns that import a Blueprint under a different
// local name (e.g. `from .auth import auth as auth_blueprint`) still key the
// override onto the same variable the route decorators reference.
func applyRegisterBlueprintOverrides(fi fileInfo, prefixes blueprintPrefix) {
	aliases := collectImportAliases(fi.root, fi.src)
	calls := findAllByType(fi.root, "call")
	for _, call := range calls {
		varName, overridePrefix := tryParseRegisterBlueprint(call, fi.src)
		if varName == "" {
			continue
		}
		if overridePrefix != "" {
			prefixes[canonicalBlueprintName(varName, aliases)] = overridePrefix
		}
	}
}
