//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 타입이 Java 컬렉션 타입인지 확인한다
package spring

import "strings"

func isCollectionType(jtype string) bool {
	for _, prefix := range []string{"List<", "Set<", "Collection<", "ArrayList<", "LinkedList<", "HashSet<"} {
		if strings.HasPrefix(jtype, prefix) && strings.HasSuffix(jtype, ">") {
			return true
		}
	}
	return false
}
