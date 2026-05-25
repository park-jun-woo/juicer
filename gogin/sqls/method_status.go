//ff:type feature=sql type=model
//ff:what MethodStatus 데이터 구조
package sqls

// MethodStatus tracks one method's ratchet state.
//
//ff:func MethodStatus
//ff:what 개별 메서드의 ratchet 상태 (TODO/DONE/SKIP)
type MethodStatus struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	QueryName string `json:"query_name,omitempty"`
}
