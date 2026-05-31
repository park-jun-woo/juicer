//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what dbNativeType @db.<Type> → 소문자 SQL 타입 변환 테스트
package prisma

import "testing"

func TestDBNativeType(t *testing.T) {
	v, ok := dbNativeType("@db.VarChar(255)")
	if !ok || v != "varchar(255)" {
		t.Errorf("varchar: got (%q,%v)", v, ok)
	}
	v, ok = dbNativeType("@db.Uuid")
	if !ok || v != "uuid" {
		t.Errorf("uuid: got (%q,%v)", v, ok)
	}
	if _, ok := dbNativeType("@id"); ok {
		t.Error("non-db prefix must be false")
	}
	if _, ok := dbNativeType("@db."); ok {
		t.Error("empty type must be false")
	}
}
