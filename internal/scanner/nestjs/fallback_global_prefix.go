//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what setGlobalPrefix 문자열 추출 실패 시 .env.example 또는 config에서 접두사를 찾는다
package nestjs

// fallbackGlobalPrefix tries to resolve the global prefix from .env files
// or config source when setGlobalPrefix uses a non-literal argument.
func fallbackGlobalPrefix(root string) string {
	if v := readEnvPrefix(root, ".env.example"); v != "" {
		return v
	}
	if v := readEnvPrefix(root, ".env"); v != "" {
		return v
	}
	return readConfigDefault(root)
}
