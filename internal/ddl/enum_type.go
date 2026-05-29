//ff:type feature=ddl type=model
//ff:what enum 타입 상태 모델 (타입명 + 값 목록)
package ddl

// EnumType represents a Postgres enum type emitted as CREATE TYPE ... AS ENUM.
type EnumType struct {
	Name   string
	Values []string
}
