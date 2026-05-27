//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 파일 간 dotted identifier prefix를 해석한다
package fastapi

// resolveDottedPrefixes iterates over all files' prefixes and resolves any
// dotted identifier values (e.g., "settings.API_V1_STR") by looking up the
// attribute value across all parsed files.
func resolveDottedPrefixes(absRoot string, files []fileInfo) {
	for i := range files {
		resolveDottedPrefixesForFile(absRoot, files, i)
	}
}
