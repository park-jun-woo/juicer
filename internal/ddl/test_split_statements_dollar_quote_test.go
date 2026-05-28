//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what splitStatements가 달러 인용 내 세미콜론을 무시하는지 테스트
package ddl

import "testing"

func TestSplitStatements_DollarQuote(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "dollar-dollar quote",
			in:   "CREATE FUNCTION f() RETURNS void AS $$ BEGIN RAISE NOTICE 'x'; END; $$ LANGUAGE plpgsql; CREATE TABLE t (id INT)",
			want: 2,
		},
		{
			name: "tagged dollar quote",
			in:   "CREATE FUNCTION f() RETURNS void AS $fn$ BEGIN RAISE NOTICE 'x'; END; $fn$ LANGUAGE plpgsql; CREATE TABLE t (id INT)",
			want: 2,
		},
		{
			name: "unbalanced parens in dollar quote",
			in:   "CREATE FUNCTION f() RETURNS void AS $$ IF (x THEN y; END; $$ LANGUAGE plpgsql; SELECT 1",
			want: 2,
		},
		{
			name: "no dollar quote",
			in:   "CREATE TABLE a (id INT); CREATE TABLE b (id INT)",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts := splitStatements(tt.in)
			if len(stmts) != tt.want {
				t.Fatalf("expected %d statements, got %d: %v", tt.want, len(stmts), stmts)
			}
		})
	}
}
