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
	for id, indices := range ids {
		if len(indices) == 1 {
			result[indices[0]] = id
			seen[id] = true
			continue
		}
		processDuplicateGroup(id, indices, endpoints, result, seen)
	}
	return result
}
