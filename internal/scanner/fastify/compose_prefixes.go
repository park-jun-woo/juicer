//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 부모 prefix 집합과 마운트 prefix를 합성하여 자식 prefix 목록을 만든다
package fastify

func composePrefixes(parents []string, prefix string) []string {
	if len(parents) == 0 {
		return []string{prefix}
	}
	var out []string
	for _, p := range parents {
		out = append(out, joinFastifyPath(p, prefix))
	}
	return out
}
