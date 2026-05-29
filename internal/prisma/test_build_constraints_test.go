//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what buildConstraints 테이블 테스트 (필드 @id/복합 @@id/복합 @@unique/복합 FK/onDelete)
package prisma

import (
	"testing"
)

func TestBuildConstraints(t *testing.T) {
	cases := []struct {
		name   string
		models []model
		target string // model name whose constraints are checked
		want   []string
	}{
		{
			name: "field @id primary key",
			models: []model{
				{name: "User", tableName: "User", fields: []field{
					{name: "id", baseType: "Int", attrs: []string{"@id"}},
					{name: "email", baseType: "String"},
				}},
			},
			target: "User",
			want:   []string{"PRIMARY KEY (id)"},
		},
		{
			name: "composite @@id primary key",
			models: []model{
				{name: "Membership", tableName: "Membership",
					fields: []field{
						{name: "orgId", baseType: "Int"},
						{name: "userId", baseType: "Int"},
					},
					blockAttrs: []string{"@@id([orgId, userId])"}},
			},
			target: "Membership",
			want:   []string{"PRIMARY KEY (orgId, userId)"},
		},
		{
			name: "composite @@unique (multitenancy orgId,email)",
			models: []model{
				{name: "Account", tableName: "Account",
					fields: []field{
						{name: "id", baseType: "Int", attrs: []string{"@id"}},
						{name: "orgId", baseType: "Int"},
						{name: "email", baseType: "String"},
					},
					blockAttrs: []string{"@@unique([orgId, email])"}},
			},
			target: "Account",
			want:   []string{"PRIMARY KEY (id)", "UNIQUE (orgId, email)"},
		},
		{
			name: "composite FK with onDelete Cascade",
			models: []model{
				{name: "Org", tableName: "Org", fields: []field{
					{name: "id", baseType: "Int", attrs: []string{"@id"}},
				}},
				{name: "Entity", tableName: "Entity", fields: []field{
					{name: "id", baseType: "Int", attrs: []string{"@id"}},
					{name: "orgId", baseType: "Int"},
				}},
				{name: "Child", tableName: "Child", fields: []field{
					{name: "id", baseType: "Int", attrs: []string{"@id"}},
					{name: "entityId", baseType: "Int"},
					{name: "orgId", baseType: "Int"},
					{name: "entity", baseType: "Entity", attrs: []string{
						"@relation(fields: [entityId, orgId], references: [id, orgId], onDelete: Cascade)",
					}},
				}},
			},
			target: "Child",
			want: []string{
				"PRIMARY KEY (id)",
				"FOREIGN KEY (entityId, orgId) REFERENCES Entity (id, orgId) ON DELETE CASCADE",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertBuildConstraints(t, c.models, c.target, c.want)
		})
	}
}
