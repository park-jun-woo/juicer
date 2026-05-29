//ff:func feature=ddl type=test control=sequence
//ff:what consumeBlockCommentLine의 미종료/종료-빈나머지/종료-나머지 3분기 테스트
package ddl

import "testing"

func TestConsumeBlockCommentLine(t *testing.T) {
	// Case 1: no "*/" — still inside the block, line consumed.
	rest, inside := consumeBlockCommentLine([]string{"still comment", "next"}, "still comment")
	if !inside {
		t.Fatal("expected still inside block when no */ present")
	}
	if len(rest) != 1 || rest[0] != "next" {
		t.Fatalf("expected first line consumed, got %v", rest)
	}

	// Case 2: "*/" with empty remainder — block ends, line consumed.
	rest, inside = consumeBlockCommentLine([]string{"end of comment */", "next"}, "end of comment */")
	if inside {
		t.Fatal("expected block ended on */")
	}
	if len(rest) != 1 || rest[0] != "next" {
		t.Fatalf("expected first line consumed, got %v", rest)
	}

	// Case 3: "*/" with trailing content — block ends, remainder kept on lines[0].
	lines := []string{"*/ CREATE TABLE t (id INT)", "next"}
	rest, inside = consumeBlockCommentLine(lines, "*/ CREATE TABLE t (id INT)")
	if inside {
		t.Fatal("expected block ended with trailing content")
	}
	if len(rest) != 2 || rest[0] != "CREATE TABLE t (id INT)" {
		t.Fatalf("expected remainder retained as first line, got %v", rest)
	}
}
