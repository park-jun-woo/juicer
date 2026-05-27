//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what main.ts 우선, src/*.ts 순회로 setGlobalPrefix 호출을 찾아 전역 접두사를 반환한다
package nestjs

// detectGlobalPrefix searches for setGlobalPrefix('prefix') in main.ts first,
// then falls back to other src/*.ts files. When the argument is not a string
// literal (e.g. configService.getOrThrow(...)), it falls back to .env.example
// or config file defaults.
func detectGlobalPrefix(root string) string {
	candidates := collectPrefixCandidates(root)

	// 순회: 리터럴 발견 시 즉시 반환, 비리터럴 발견 시 fallback 후보 기록
	needFallback := false
	for _, path := range candidates {
		prefix, found := detectGlobalPrefixInFile(path)
		if prefix != "" {
			return prefix
		}
		if found {
			needFallback = true
		}
	}
	if needFallback {
		return fallbackGlobalPrefix(root)
	}
	return ""
}
