//ff:func feature=scan type=extract control=sequence dimension=1 topic=express
//ff:what 부모 파일의 prefix가 있으면 현재 엔트리의 prefix와 결합한다
package express

func applyParentPrefix(prefixMap map[string]string, m mountEntry) string {
	parentPrefix := prefixMap[m.sourceFile]
	if parentPrefix == "" {
		return ""
	}
	return joinExpressPath(parentPrefix, m.prefix)
}
