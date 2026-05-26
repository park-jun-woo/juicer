//ff:func feature=ddl type=test control=sequence
//ff:what ALTER TABLE ALTER COLUMN 및 ADD/DROP CONSTRAINT 통합 테스트
package ddl

import "testing"

func TestApplyAlterTable_AlterColumn(t *testing.T) {
	tables := map[string]*Table{
		"form_templates": {
			Name: "form_templates",
			Columns: []Column{
				{Name: "id", Raw: "id UUID"},
				{Name: "group_id", Raw: "group_id UUID"},
			},
		},
		"users": {
			Name: "users",
			Columns: []Column{
				{Name: "id", Raw: "id UUID"},
				{Name: "email", Raw: "email TEXT NOT NULL DEFAULT ''"},
				{Name: "password_hash", Raw: "password_hash TEXT NOT NULL DEFAULT ''"},
				{Name: "role", Raw: "role TEXT NOT NULL"},
			},
			Constraints: []string{
				"CONSTRAINT users_email_key UNIQUE (email)",
				"CONSTRAINT users_password_hash_key UNIQUE (password_hash)",
			},
		},
		"contracts": {
			Name: "contracts",
			Columns: []Column{
				{Name: "id", Raw: "id UUID"},
				{Name: "payment_type", Raw: "payment_type TEXT NOT NULL"},
				{Name: "unit_id", Raw: "unit_id UUID NOT NULL"},
			},
		},
		"tenant_billings": {
			Name: "tenant_billings",
			Columns: []Column{
				{Name: "id", Raw: "id UUID"},
				{Name: "contract_id", Raw: "contract_id UUID"},
			},
		},
	}

	// 000007: ALTER COLUMN group_id SET NOT NULL
	applyAlterTable(tables, "form_templates", "ALTER COLUMN group_id SET NOT NULL")
	if tables["form_templates"].Columns[1].Raw != "group_id UUID NOT NULL" {
		t.Fatalf("group_id SET NOT NULL failed: %q", tables["form_templates"].Columns[1].Raw)
	}

	// 000021: ALTER COLUMN role SET DEFAULT 'admin'
	applyAlterTable(tables, "users", "ALTER COLUMN role SET DEFAULT 'admin'")
	if tables["users"].Columns[3].Raw != "role TEXT NOT NULL DEFAULT 'admin'" {
		t.Fatalf("role SET DEFAULT failed: %q", tables["users"].Columns[3].Raw)
	}

	// 000023: ALTER COLUMN email DROP NOT NULL
	applyAlterTable(tables, "users", "ALTER COLUMN email DROP NOT NULL")
	if tables["users"].Columns[1].Raw != "email TEXT DEFAULT ''" {
		t.Fatalf("email DROP NOT NULL failed: %q", tables["users"].Columns[1].Raw)
	}

	// 000023: ALTER COLUMN email DROP DEFAULT
	applyAlterTable(tables, "users", "ALTER COLUMN email DROP DEFAULT")
	if tables["users"].Columns[1].Raw != "email TEXT" {
		t.Fatalf("email DROP DEFAULT failed: %q", tables["users"].Columns[1].Raw)
	}

	// 000023: ALTER COLUMN password_hash DROP NOT NULL
	applyAlterTable(tables, "users", "ALTER COLUMN password_hash DROP NOT NULL")
	if tables["users"].Columns[2].Raw != "password_hash TEXT DEFAULT ''" {
		t.Fatalf("password_hash DROP NOT NULL failed: %q", tables["users"].Columns[2].Raw)
	}

	// 000023: ALTER COLUMN password_hash DROP DEFAULT
	applyAlterTable(tables, "users", "ALTER COLUMN password_hash DROP DEFAULT")
	if tables["users"].Columns[2].Raw != "password_hash TEXT" {
		t.Fatalf("password_hash DROP DEFAULT failed: %q", tables["users"].Columns[2].Raw)
	}

	// 000028: ALTER COLUMN payment_type SET DEFAULT 'prepaid'
	applyAlterTable(tables, "contracts", "ALTER COLUMN payment_type SET DEFAULT 'prepaid'")
	if tables["contracts"].Columns[1].Raw != "payment_type TEXT NOT NULL DEFAULT 'prepaid'" {
		t.Fatalf("payment_type SET DEFAULT failed: %q", tables["contracts"].Columns[1].Raw)
	}

	// 000036: ALTER COLUMN unit_id DROP NOT NULL
	applyAlterTable(tables, "contracts", "ALTER COLUMN unit_id DROP NOT NULL")
	if tables["contracts"].Columns[2].Raw != "unit_id UUID" {
		t.Fatalf("unit_id DROP NOT NULL failed: %q", tables["contracts"].Columns[2].Raw)
	}

	// 000023: DROP CONSTRAINT IF EXISTS users_email_key
	applyAlterTable(tables, "users", "DROP CONSTRAINT IF EXISTS users_email_key")
	if len(tables["users"].Constraints) != 1 {
		t.Fatalf("expected 1 constraint after drop, got %d", len(tables["users"].Constraints))
	}

	// 000029: ADD CONSTRAINT uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)
	applyAlterTable(tables, "tenant_billings", "ADD CONSTRAINT uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)")
	if len(tables["tenant_billings"].Constraints) != 1 {
		t.Fatalf("expected 1 constraint after add, got %d", len(tables["tenant_billings"].Constraints))
	}
}
