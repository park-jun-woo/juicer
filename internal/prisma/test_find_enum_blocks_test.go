//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what findEnumBlocks가 enum 블록만 수집(model/generator/datasource 무시)하는지 검증
package prisma

import (
	"reflect"
	"testing"
)

func TestFindEnumBlocks(t *testing.T) {
	const src = `
generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

enum Role {
  USER
  ADMIN
}

model User {
  id   Int  @id
  role Role @default(USER)
}

enum TokenType {
  ACCESS
  REFRESH
}
`
	cases := []struct {
		name string
		key  string
		want []string
	}{
		{"Role enum collected", "Role", []string{"USER", "ADMIN"}},
		{"TokenType enum collected", "TokenType", []string{"ACCESS", "REFRESH"}},
	}

	blocks := findEnumBlocks(stripComments(src))

	t.Run("only enums collected", func(t *testing.T) {
		if len(blocks) != len(cases) {
			t.Fatalf("expected %d enum blocks, got %d: %v", len(cases), len(blocks), blocks)
		}
		if _, ok := blocks["User"]; ok {
			t.Errorf("model block %q must not be collected as enum", "User")
		}
		if _, ok := blocks["client"]; ok {
			t.Errorf("generator block %q must not be collected as enum", "client")
		}
		if _, ok := blocks["db"]; ok {
			t.Errorf("datasource block %q must not be collected as enum", "db")
		}
		for _, c := range cases {
			got := parseEnumValues(blocks[c.key])
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%s: values = %v, want %v", c.name, got, c.want)
			}
		}
	})
}
