//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what enumPathCtx로 단일 경로 인자의 enum 멤버표현식을 해석한다
package nestjs

// resolve returns the enum-resolved value of arg, or arg unchanged when it is
// not a resolvable member expression.
func (pc enumPathCtx) resolve(arg string) string {
	if resolved, ok := resolveEnumPathArg(arg, pc.root, pc.src, pc.absFile, pc.imports, pc.projectRoot); ok {
		return resolved
	}
	return arg
}
