//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 Serializer 클래스를 수집한다
package django

// extractSerializers finds all Serializer classes in the parsed files.
func extractSerializers(files []fileInfo) map[string]serializerInfo {
	result := make(map[string]serializerInfo)
	for _, fi := range files {
		for _, si := range extractSerializersFromFile(fi) {
			result[si.name] = si
		}
	}
	return result
}
