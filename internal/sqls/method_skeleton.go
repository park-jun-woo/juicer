//ff:type feature=sql type=model
//ff:what MethodSkeleton 데이터 구조
package sqls

// MethodSkeleton describes one repository method's SQL usage.
//
//ff:func MethodSkeleton
//ff:what 레포지토리 메서드의 SQL 사용 정보
type MethodSkeleton struct {
	Repo         string   `yaml:"repo"                    json:"repo"`
	Method       string   `yaml:"method"                  json:"method"`
	CRUD         string   `yaml:"crud"                    json:"crud"`
	Tables       []string `yaml:"tables"                  json:"tables"`
	Params       []string `yaml:"params,omitempty"         json:"params,omitempty"`
	Returns      []string `yaml:"returns,omitempty"        json:"returns,omitempty"`
	SQLFragments []string `yaml:"sql_fragments,omitempty"  json:"sql_fragments,omitempty"`
	Dynamic      bool     `yaml:"dynamic,omitempty"        json:"dynamic,omitempty"`
}
