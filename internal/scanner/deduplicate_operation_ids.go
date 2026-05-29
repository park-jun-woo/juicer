//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 중복 operationId에 경로 prefix를 추가하여 유일성을 보장한다
package scanner

func deduplicateOperationIDs(endpoints []Endpoint) map[int]string {
	ids := map[string][]int{}
	for i, ep := range endpoints {
		id := generateOperationID(ep)
		ids[id] = append(ids[id], i)
	}
	result := map[int]string{}
	seen := map[string]bool{}
	processed := map[string]bool{}
	// ids는 map이므로 직접 range하면 비결정적이다.
	// endpoint 인덱스 순서로 순회하고, 중복 그룹은 첫 등장 시점에 한 번만 처리한다.
	for i, ep := range endpoints {
		id := generateOperationID(ep)
		indices := ids[id]
		if len(indices) == 1 {
			result[i] = id
			seen[id] = true
			continue
		}
		if processed[id] {
			continue
		}
		processed[id] = true
		processDuplicateGroup(id, indices, endpoints, result, seen)
	}
	return result
}
