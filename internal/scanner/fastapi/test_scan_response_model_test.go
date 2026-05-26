//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestScan_ResponseModel 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_ResponseModel(t *testing.T) {
	dir := t.TempDir()
	src := `from fastapi import FastAPI
from pydantic import BaseModel

class Item(BaseModel):
    name: str
    price: float

app = FastAPI()

@app.get("/items", response_model=Item)
async def get_item():
    pass
`
	os.WriteFile(filepath.Join(dir, "main.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].TypeName != "Item" {
		t.Fatalf("expected Item, got %s", ep.Responses[0].TypeName)
	}
}
