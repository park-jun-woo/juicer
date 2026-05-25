//ff:type feature=ddl type=model
//ff:what 테이블 상태 모델 (컬럼 + 제약 + 인덱스)
package ddl

// Table represents the current state of a table after applying migrations.
type Table struct {
	Name        string
	Columns     []Column
	Constraints []string
	Indexes     []string
}
