//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what "a::b::c" 형태 scoped 식별자를 :: 기준으로 분할한다
package actix

func splitScoped(s string) []string {
	var parts []string
	current := ""
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == ':' && s[i+1] == ':' {
			parts = append(parts, current)
			current = ""
			i++
			continue
		}
		current += string(s[i])
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
