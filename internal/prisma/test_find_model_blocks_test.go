//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what findModelBlocks model 블록만 name→body로 수집(다른 블록 스킵) 테스트
package prisma

import "testing"

func TestFindModelBlocks(t *testing.T) {
	src := `generator client { provider = "x" }
model User {
  id Int @id
}
enum Role {
  ADMIN
}`
	blocks := findModelBlocks(src)
	if len(blocks) != 1 {
		t.Fatalf("got %d blocks: %v", len(blocks), blocks)
	}
	body, ok := blocks["User"]
	if !ok || len(body) != 1 || body[0] != "id Int @id" {
		t.Errorf("User body: %v", body)
	}
}
