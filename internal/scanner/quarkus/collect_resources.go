//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 파싱된 Java 파일에서 @Path 리소스 클래스를 수집한다
package quarkus

func collectResources(files []*fileInfo) []resourceInfo {
	var result []resourceInfo
	for _, fi := range files {
		resources := extractResources(fi)
		result = append(result, resources...)
	}
	return result
}
