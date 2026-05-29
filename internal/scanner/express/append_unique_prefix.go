//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what key의 prefix 목록에 v가 없으면 추가하고 추가 여부를 반환한다
package express

// appendUniquePrefix는 key의 prefix 목록에 v가 없으면 추가하고 true를 반환한다.
func appendUniquePrefix(m map[routerKey][]string, key routerKey, v string) bool {
	for _, existing := range m[key] {
		if existing == v {
			return false
		}
	}
	m[key] = append(m[key], v)
	return true
}
