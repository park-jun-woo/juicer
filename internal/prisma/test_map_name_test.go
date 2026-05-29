//ff:func feature=prisma type=test control=sequence
//ff:what @@map/@map 테이블·컬럼명 치환 테스트 (Parse 엔트리포인트 경유)
package prisma

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMapNameSubstitution(t *testing.T) {
	const src = `
model User {
  id     Int    @id
  email  String @map("email_address")
  posts  Post[]
  @@map("users")
}

model Post {
  id       Int  @id
  authorId Int  @map("author_id")
  author   User @relation(fields: [authorId], references: [id])
  @@map("posts")
}
`
	dir := t.TempDir()
	path := filepath.Join(dir, "schema.prisma")
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}

	tables, _, err := Parse(path)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	t.Run("@@map renames table key", func(t *testing.T) {
		if _, ok := tables["users"]; !ok {
			t.Errorf("expected table %q from @@map, got tables %v", "users", tableKeys(tables))
		}
		if _, ok := tables["User"]; ok {
			t.Errorf("model name %q must not survive as a table key after @@map", "User")
		}
	})

	post := tables["posts"]
	if post == nil {
		t.Fatalf("expected table %q from @@map, got tables %v", "posts", tableKeys(tables))
	}

	t.Run("@map renames column", func(t *testing.T) {
		col := findColumn(post.Columns, "author_id")
		if col == nil {
			t.Fatalf("expected column %q from @map; columns: %v", "author_id", columnNames(post.Columns))
		}
		if !strings.HasPrefix(col.Raw, "author_id ") {
			t.Errorf("column Raw should start with mapped name, got %q", col.Raw)
		}
		if findColumn(post.Columns, "authorId") != nil {
			t.Errorf("original field name %q must be replaced by @map", "authorId")
		}
	})

	t.Run("FK uses mapped table and column names", func(t *testing.T) {
		fk := strings.Join(post.Constraints, " | ")
		for _, want := range []string{"FOREIGN KEY (author_id)", "REFERENCES users (id)"} {
			if !strings.Contains(fk, want) {
				t.Errorf("expected FK to contain %q, got %q", want, fk)
			}
		}
	})
}
