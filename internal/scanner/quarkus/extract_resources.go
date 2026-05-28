//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what fileInfo에서 @Path 리소스 클래스를 추출한다
package quarkus

func extractResources(fi *fileInfo) []resourceInfo {
	var result []resourceInfo
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		if !hasAnnotation(cls, fi.src, AnnPath) {
			continue
		}
		ri := buildResourceInfo(cls, fi)
		result = append(result, ri)
	}
	return result
}
