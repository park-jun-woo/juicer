//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what Prisma 스칼라/Unsupported/@db native 타입을 SQL 타입으로 매핑
package prisma

import "strings"

// scalarTypes maps Prisma scalar names to PostgreSQL types.
var scalarTypes = map[string]string{
	"String":   "text",
	"Int":      "integer",
	"BigInt":   "bigint",
	"Boolean":  "boolean",
	"DateTime": "timestamp(3)",
	"Float":    "double precision",
	"Decimal":  "numeric",
	"Json":     "jsonb",
	"Bytes":    "bytea",
}

// mapType resolves the SQL type for a field: @db native first, then
// Unsupported("...") raw text, then the scalar table, else the type as-is.
func mapType(f field) string {
	for _, a := range f.attrs {
		if native, ok := dbNativeType(a); ok {
			return native
		}
	}
	if strings.HasPrefix(f.baseType, "Unsupported(") {
		return unsupportedType(f.baseType)
	}
	if sqlType, ok := scalarTypes[f.baseType]; ok {
		return sqlType
	}
	return f.baseType
}
