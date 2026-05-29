//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 경로가 주어진 prefix로 시작하는지 판별한다("" 또는 "/"는 항상 참)
package actix

func hasPrefix(path, prefix string) bool {
	if prefix == "" || prefix == "/" {
		return true
	}
	return len(path) >= len(prefix) && path[:len(prefix)] == prefix
}
