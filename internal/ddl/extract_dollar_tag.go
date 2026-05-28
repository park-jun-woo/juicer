//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what 달러 인용 태그 추출 ($$ 또는 $tag$)
package ddl

// extractDollarTag checks if runes starting at position i form a dollar-quote
// tag (either $$ or $tag$). Returns the tag string if found, or "" otherwise.
// A valid tag is $$ or $identifier$ where identifier is [a-zA-Z_][a-zA-Z0-9_]*.
func extractDollarTag(runes []rune, i int) string {
	n := len(runes)
	if i >= n || runes[i] != '$' {
		return ""
	}

	// Check for $$ (empty tag)
	if i+1 < n && runes[i+1] == '$' {
		return "$$"
	}

	// Check for $identifier$
	if i+1 >= n {
		return ""
	}
	ch := runes[i+1]
	if !isDollarTagStart(ch) {
		return ""
	}
	j := i + 2
	for j < n && isDollarTagCont(runes[j]) {
		j++
	}
	if j < n && runes[j] == '$' {
		return string(runes[i : j+1])
	}
	return ""
}
