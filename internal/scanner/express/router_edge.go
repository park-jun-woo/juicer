//ff:type feature=scan type=model topic=express
//ff:what 부모 라우터에서 자식 라우터로의 마운트(세그먼트) 엣지
package express

// routerEdge는 부모 라우터에서 자식 라우터로의 마운트(세그먼트)다.
type routerEdge struct {
	child routerKey
	seg   string
}
