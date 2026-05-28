//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 타입이 Java 컬렉션 타입인지 확인한다
package quarkus

import "strings"

func isCollectionType(jtype string) bool {
	for _, prefix := range []string{"List<", "Set<", "Collection<", "ArrayList<", "LinkedList<", "HashSet<"} {
		if strings.HasPrefix(jtype, prefix) && strings.HasSuffix(jtype, ">") {
			return true
		}
	}
	return false
}
