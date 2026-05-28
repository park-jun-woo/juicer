//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what register_blueprint 호출에서 Blueprint prefix 전파를 해석한다
package flask

// resolveBlueprintPrefixes builds a map from blueprint variable name to resolved prefix.
// It checks register_blueprint calls for url_prefix overrides.
func resolveBlueprintPrefixes(files []fileInfo) blueprintPrefix {
	prefixes := make(blueprintPrefix)
	// Pass 1: collect all blueprint definitions
	for _, fi := range files {
		bps := collectBlueprints(fi.root, fi.src)
		for _, bp := range bps {
			prefixes[bp.varName] = bp.urlPrefix
		}
	}
	// Pass 2: check register_blueprint calls for prefix overrides
	for _, fi := range files {
		applyRegisterBlueprintOverrides(fi, prefixes)
	}
	return prefixes
}
