//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 경로 세그먼트를 결합한다
package quarkus

import "strings"

func joinPath(parts ...string) string {
	var segs []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		p = strings.Trim(p, "/")
		if p != "" {
			segs = append(segs, p)
		}
	}
	result := "/" + strings.Join(segs, "/")
	if result != "/" {
		result = strings.TrimRight(result, "/")
	}
	return result
}
