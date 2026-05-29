//ff:func feature=scan type=extract control=sequence
//ff:what HTTP method를 접두로 붙여 operationId를 만든다 (get + Login = getLogin)
package scanner

import "strings"

func methodPrefixedID(method, id string) string {
	prefix := strings.ToLower(method)
	if prefix == "" || id == "" {
		return prefix + id
	}
	return prefix + strings.ToUpper(id[:1]) + id[1:]
}
