//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_FastAPIFramework FastAPI 프레임워크 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_FastAPIFramework(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "main.py"), []byte(`from fastapi import FastAPI
app = FastAPI()
@app.get("/")
def root():
    return {"msg": "ok"}
`), 0o644)
	execScan([]string{"--framework", "fastapi", dir})
}
