//ff:func feature=prisma type=convert control=selection topic=prisma
//ff:what autoincrement(빈 DEFAULT) 정수 타입을 serial/bigserial로 승격
package prisma

// promoteSerial upgrades integer/bigint columns with an empty (autoincrement)
// default to serial/bigserial. Other types are returned unchanged.
func promoteSerial(sqlType, def string, hasDefault bool) string {
	if !hasDefault || def != "" {
		return sqlType
	}
	switch sqlType {
	case "integer":
		return "serial"
	case "bigint":
		return "bigserial"
	default:
		return sqlType
	}
}
