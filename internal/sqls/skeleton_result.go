//ff:type feature=sql type=model
//ff:what SkeletonResult 데이터 구조
package sqls

// SkeletonResult holds all extracted method skeletons.
//
//ff:func SkeletonResult
//ff:what SQL 스켈레톤 추출 결과 컨테이너
type SkeletonResult struct {
	Methods []MethodSkeleton `yaml:"methods" json:"methods"`
}
