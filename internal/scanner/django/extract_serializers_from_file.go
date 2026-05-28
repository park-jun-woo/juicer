//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 Serializer 클래스를 수집한다
package django

// extractSerializersFromFile finds Serializer classes in a single file.
func extractSerializersFromFile(fi fileInfo) []serializerInfo {
	var serializers []serializerInfo
	for _, classNode := range findAllByType(fi.root, "class_definition") {
		si := parseSerializerClass(classNode, fi)
		if si != nil {
			serializers = append(serializers, *si)
		}
	}
	return serializers
}
