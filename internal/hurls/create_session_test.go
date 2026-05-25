package hurls

import "testing"

func TestCreateSession_MissingHost(t *testing.T) {
	err := createSession("", "tests", "repo")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateSession_MissingTests(t *testing.T) {
	err := createSession("http://localhost", "", "repo")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateSession_MissingRepo(t *testing.T) {
	err := createSession("http://localhost", "tests", "")
	if err == nil {
		t.Fatal("expected error")
	}
}
