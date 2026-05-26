package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func getProjectRoot(t *testing.T) string {
	// Try to find project root by looking for go.mod
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get working directory: %v", err)
	}

	// Walk up the directory tree to find go.mod
	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// Check if this is the yi package (has data directory)
			if _, err := os.Stat(filepath.Join(dir, "data")); err == nil {
				return dir
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	t.Fatalf("cannot find project root from %s", cwd)
	return ""
}

func TestLoadAllJSON(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	data, err := loadAllJSON()
	if err != nil {
		t.Fatalf("loadAllJSON failed: %v", err)
	}

	// Verify bagua count
	if len(data.Bagua) != 8 {
		t.Errorf("expected 8 bagua, got %d", len(data.Bagua))
	}

	// Verify gua count
	if len(data.Gua) != 64 {
		t.Errorf("expected 64 gua, got %d", len(data.Gua))
	}

	// Verify dayan count
	if len(data.Dayan) != 81 {
		t.Errorf("expected 81 dayan, got %d", len(data.Dayan))
	}

	t.Logf("Loaded: %d bagua, %d gua, %d dayan", len(data.Bagua), len(data.Gua), len(data.Dayan))
}

func TestBaguaJSONStructure(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	data, err := loadAllJSON()
	if err != nil {
		t.Fatalf("loadAllJSON failed: %v", err)
	}

	for i, b := range data.Bagua {
		if b.Name == "" {
			t.Errorf("bagua[%d] has empty name", i)
		}
		if b.Num < 0 || b.Num > 7 {
			t.Errorf("bagua[%d] has invalid num: %d", i, b.Num)
		}
		if b.Symbol == "" {
			t.Errorf("bagua[%d] has empty symbol", i)
		}
	}

	t.Logf("All %d bagua have valid structure", len(data.Bagua))
}

func TestGuaJSONStructure(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	data, err := loadAllJSON()
	if err != nil {
		t.Fatalf("loadAllJSON failed: %v", err)
	}

	for i, g := range data.Gua {
		if g.Xu < 1 || g.Xu > 64 {
			t.Errorf("gua[%d] has invalid xu: %d", i, g.Xu)
		}
		if g.Index == "" {
			t.Errorf("gua[%d] has empty index", i)
		}
		if g.Ming == "" {
			t.Errorf("gua[%d] has empty ming", i)
		}
		if g.ShangNum < 0 || g.ShangNum > 7 {
			t.Errorf("gua[%d] has invalid shang_num: %d", i, g.ShangNum)
		}
		if g.XiaNum < 0 || g.XiaNum > 7 {
			t.Errorf("gua[%d] has invalid xia_num: %d", i, g.XiaNum)
		}
	}

	t.Logf("All %d gua have valid structure", len(data.Gua))
}

func TestDayanJSONStructure(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	data, err := loadAllJSON()
	if err != nil {
		t.Fatalf("loadAllJSON failed: %v", err)
	}

	for i, d := range data.Dayan {
		if d.Number < 1 || d.Number > 81 {
			t.Errorf("dayan[%d] has invalid number: %d", i, d.Number)
		}
		if d.JiXiong == "" {
			t.Errorf("dayan[%d] has empty jixiong", i)
		}
	}

	// Count IsMax true values
	maxCount := 0
	for _, d := range data.Dayan {
		if d.IsMax {
			maxCount++
		}
	}
	t.Logf("All %d dayan have valid structure, %d are marked as IsMax", len(data.Dayan), maxCount)
}

func TestGuaYaoData(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	data, err := loadAllJSON()
	if err != nil {
		t.Fatalf("loadAllJSON failed: %v", err)
	}

	// Count gua with yao data
	guaWithYao := 0
	totalYao := 0

	for _, g := range data.Gua {
		hasYao := false
		for _, y := range g.Yaos {
			if y.Ci != "" {
				hasYao = true
				totalYao++
			}
		}
		if hasYao {
			guaWithYao++
		}
	}

	t.Logf("%d/%d gua have yao data, total %d yao entries", guaWithYao, len(data.Gua), totalYao)
}

func TestJSONFilesAreValid(t *testing.T) {
	projectRoot := getProjectRoot(t)
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("cannot change to project root: %v", err)
	}

	files := []string{
		"data/bagua.json",
		"data/gua.json",
		"data/dayan.json",
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("cannot read %s: %v", file, err)
			continue
		}

		var data interface{}
		if err := json.Unmarshal(content, &data); err != nil {
			t.Errorf("%s is not valid JSON: %v", file, err)
		} else {
			t.Logf("%s is valid JSON", file)
		}
	}
}
