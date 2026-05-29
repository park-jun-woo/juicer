//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what register 변수명을 alias 맵으로 정식 Blueprint 변수명으로 역해석한다
package flask

// canonicalBlueprintName resolves a register_blueprint argument name back to the
// canonical Blueprint variable name using the file's import alias map.
// If the name is not an aliased import, it is returned unchanged.
func canonicalBlueprintName(varName string, aliases importAlias) string {
	if orig, ok := aliases[varName]; ok {
		return orig
	}
	return varName
}
