//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 확장자 없는 경로에 TS/JS 확장자(.ts/.tsx/.js/.jsx/.mjs/.cjs) 또는 /index.<ext>를 붙여 실제 파일 경로를 반환한다
package express

func resolveExtension(base string) string {
	return resolveSourceBase(base)
}
