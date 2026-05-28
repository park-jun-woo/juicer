//ff:func feature=ddl type=test control=sequence
//ff:what applyStatement가 스키마 한정/따옴표 테이블명을 정상 처리하는지 테스트
package ddl

import "testing"

func TestApplyStatement_SchemaQualified(t *testing.T) {
	tables := make(map[string]*Table)

	// Schema-qualified table name
	applyStatement(tables, "CREATE TABLE public.profiles (id UUID PRIMARY KEY, name TEXT)")
	if tables["profiles"] == nil {
		t.Fatal("expected profiles table from public.profiles")
	}

	// Quoted table name
	applyStatement(tables, `CREATE TABLE "UserEvents" (id INT PRIMARY KEY)`)
	if tables["userevents"] == nil {
		t.Fatal("expected userevents table from quoted name")
	}

	// Schema + quoted
	applyStatement(tables, `CREATE TABLE public."Payments" (id INT PRIMARY KEY)`)
	if tables["payments"] == nil {
		t.Fatal("expected payments table from public.\"Payments\"")
	}

	// ALTER TABLE with schema prefix
	applyStatement(tables, "ALTER TABLE public.profiles ADD COLUMN email TEXT")
	if len(tables["profiles"].Columns) != 3 {
		t.Fatalf("expected 3 columns after alter, got %d", len(tables["profiles"].Columns))
	}

	// DROP TABLE with schema prefix
	applyStatement(tables, "DROP TABLE public.profiles")
	if tables["profiles"] != nil {
		t.Fatal("expected profiles to be dropped")
	}
}
