//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestScan_FullProject 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_FullProject(t *testing.T) {
	dir := t.TempDir()
	src := `from fastapi import FastAPI, APIRouter, Query, Depends, UploadFile
from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
    email: str

class UserResponse(BaseModel):
    id: int
    name: str
    email: str

app = FastAPI()
router = APIRouter(prefix="/users")

@router.get("/")
async def list_users(skip: int = 0, limit: int = Query(default=100)):
    pass

@router.get("/{user_id}")
async def get_user(user_id: int) -> UserResponse:
    pass

@router.post("/", status_code=201)
async def create_user(user: UserCreate):
    pass

app.include_router(router)
`
	os.WriteFile(filepath.Join(dir, "main.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Method != "GET" || ep.Path != "/users" || ep.Handler != "list_users" {
		t.Fatalf("endpoint 0: got %s %s %s", ep.Method, ep.Path, ep.Handler)
	}
	if ep.Request == nil || len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %v", ep.Request)
	}

	ep = result.Endpoints[1]
	if ep.Method != "GET" || ep.Path != "/users/{user_id}" {
		t.Fatalf("endpoint 1: got %s %s", ep.Method, ep.Path)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param, got %v", ep.Request)
	}

	ep = result.Endpoints[2]
	if ep.Method != "POST" || ep.Path != "/users" {
		t.Fatalf("endpoint 2: got %s %s", ep.Method, ep.Path)
	}
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected request body")
	}
	if ep.Request.Body.TypeName != "UserCreate" {
		t.Fatalf("expected UserCreate, got %s", ep.Request.Body.TypeName)
	}
	if len(ep.Request.Body.Fields) != 2 {
		t.Fatalf("expected 2 body fields, got %d", len(ep.Request.Body.Fields))
	}
}

func TestScan_NoPyFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("hi"), 0o644)
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(result.Endpoints))
	}
}

