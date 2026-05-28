//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 이름으로 APIView를 검색한다
package django

// findAPIView finds an APIView by name.
func findAPIView(apiviews []apiviewInfo, name string) *apiviewInfo {
	for i := range apiviews {
		if apiviews[i].name == name {
			return &apiviews[i]
		}
	}
	return nil
}
