//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 바깥 prefix와 안쪽 prefix를 결합한다
package laravel

// joinGroupPrefix joins outer prefix with inner prefix.
func joinGroupPrefix(outer, inner string) string {
	if outer == "" {
		return inner
	}
	if inner == "" {
		return outer
	}
	return outer + "/" + inner
}
