//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 단일 파일의 dotted identifier prefix를 해석한다
package fastapi

import "strings"

// resolveDottedPrefixesForFile resolves dotted identifier prefixes for a single
// file entry. Each prefix value that looks like "settings.ATTR" is resolved by
// looking up the attribute value across all parsed files.
func resolveDottedPrefixesForFile(absRoot string, files []fileInfo, idx int) {
	for varName, prefix := range files[idx].prefixes {
		trimmed := strings.TrimPrefix(prefix, "/")
		if !isDottedIdentifier(trimmed) {
			continue
		}
		resolved := resolveAttributeValue(absRoot, files, trimmed)
		if resolved != "" {
			files[idx].prefixes[varName] = resolved
		}
	}
}
