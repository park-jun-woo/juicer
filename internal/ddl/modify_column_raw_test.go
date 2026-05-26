//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what modifyColumnRaw의 SET/DROP NOT NULL, SET/DROP DEFAULT, TYPE 치환 테스트
package ddl

import "testing"

func TestModifyColumnRaw(t *testing.T) {
	tests := []struct {
		name   string
		raw    string
		action string
		want   string
	}{
		// SET NOT NULL
		{
			name:   "set not null on plain column",
			raw:    "group_id UUID",
			action: "SET NOT NULL",
			want:   "group_id UUID NOT NULL",
		},
		{
			name:   "set not null already present",
			raw:    "group_id UUID NOT NULL",
			action: "SET NOT NULL",
			want:   "group_id UUID NOT NULL",
		},
		// DROP NOT NULL
		{
			name:   "drop not null",
			raw:    "email TEXT NOT NULL DEFAULT ''",
			action: "DROP NOT NULL",
			want:   "email TEXT DEFAULT ''",
		},
		{
			name:   "drop not null when absent",
			raw:    "email TEXT DEFAULT ''",
			action: "DROP NOT NULL",
			want:   "email TEXT DEFAULT ''",
		},
		// SET DEFAULT
		{
			name:   "set default on column without default",
			raw:    "role TEXT NOT NULL",
			action: "SET DEFAULT 'admin'",
			want:   "role TEXT NOT NULL DEFAULT 'admin'",
		},
		{
			name:   "set default replaces existing default",
			raw:    "role TEXT NOT NULL DEFAULT 'user'",
			action: "SET DEFAULT 'admin'",
			want:   "role TEXT NOT NULL DEFAULT 'admin'",
		},
		{
			name:   "set default unquoted value",
			raw:    "payment_type TEXT NOT NULL DEFAULT 'postpaid'",
			action: "SET DEFAULT 'prepaid'",
			want:   "payment_type TEXT NOT NULL DEFAULT 'prepaid'",
		},
		// DROP DEFAULT
		{
			name:   "drop default",
			raw:    "email TEXT NOT NULL DEFAULT ''",
			action: "DROP DEFAULT",
			want:   "email TEXT NOT NULL",
		},
		{
			name:   "drop default when absent",
			raw:    "email TEXT NOT NULL",
			action: "DROP DEFAULT",
			want:   "email TEXT NOT NULL",
		},
		// TYPE
		{
			name:   "change type",
			raw:    "amount INT NOT NULL",
			action: "TYPE BIGINT",
			want:   "amount BIGINT NOT NULL",
		},
		// Unknown action
		{
			name:   "unknown action is no-op",
			raw:    "id INT",
			action: "UNKNOWN STUFF",
			want:   "id INT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modifyColumnRaw(tt.raw, tt.action)
			if got != tt.want {
				t.Errorf("modifyColumnRaw(%q, %q) = %q, want %q", tt.raw, tt.action, got, tt.want)
			}
		})
	}
}
