package gen

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestDocGenerate(t *testing.T) {
	tCases := initTestCases(t)
	for _, tCase := range tCases {
		genContext := GenContext{
			PackagePath:      tCase.PackagePath,
			Format:           Markdown,
			IgnoreDeprecated: false,
			Target:           tCase.GotMd,
			EscapeHtml:       true,
		}
		err := genContext.GenDoc()
		if err != nil {
			t.Fatalf("generate failed: %s", err)
		}
		// check the content of expected and actual
		err = CompareDir(tCase.ExpectMd, tCase.GotMd)
		if err != nil {
			t.Fatal(err)
		}
		// if test failed, keep generate files for checking
		os.RemoveAll(genContext.Target)
	}
}

func initTestCases(t *testing.T) []*TestCase {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("get work directory failed")
	}

	testdataDir := filepath.Join("testdata", "doc")
	sourcePkgs := []string{
		"k8s",
		"pkg",
		"reimport",
	}
	tcases := make([]*TestCase, len(sourcePkgs))

	for i, p := range sourcePkgs {
		resultDir := filepath.Join(cwd, testdataDir, p)
		if runtime.GOOS == "windows" {
			resultDir = filepath.Join(resultDir, "windows")
		} else {
			resultDir = filepath.Join(resultDir, "unixlike")
		}
		tcases[i] = &TestCase{
			PackagePath: filepath.Join(testdataDir, p),
			ExpectMd:    filepath.Join(resultDir, "md"),
			ExpectHtml:  filepath.Join(resultDir, "html"),
			GotMd:       filepath.Join(resultDir, "md_got"),
			GotHtml:     filepath.Join(resultDir, "html_got"),
		}
	}
	return tcases
}

type TestCase struct {
	PackagePath string
	ExpectMd    string
	ExpectHtml  string
	GotMd       string
	GotHtml     string
}

func CompareDir(a string, b string) error {
	dirA, err := os.ReadDir(a)
	if err != nil {
		return fmt.Errorf("read dir %s failed when comparing with %s", a, b)
	}
	dirB, err := os.ReadDir(b)
	if err != nil {
		return fmt.Errorf("read dir %s failed when comparing with %s", b, a)
	}
	if len(dirA) != len(dirB) {
		return fmt.Errorf("dirs contains different number of files:\n%s: %v\n%s: %v", a, len(dirA), b, len(dirB))
	}
	for _, fA := range dirA {
		// check if the same file exist in dirB
		aPath := filepath.Join(a, fA.Name())
		bPath := filepath.Join(b, fA.Name())
		_, err := os.Open(bPath)
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s exist in %s, but missing in %s", fA.Name(), a, b)
		}
		if err != nil {
			return fmt.Errorf("open file failed when compare, file path: %s", bPath)
		}
		if fA.IsDir() {
			err := CompareDir(aPath, bPath)
			if err != nil {
				return err
			}
			continue
		}
		linesA, err := readLines(aPath)
		if err != nil {
			return fmt.Errorf("failed to readlins from %s when compare files", aPath)
		}
		linesB, err := readLines(bPath)
		if err != nil {
			return fmt.Errorf("failed to readlins from %s when compare files", bPath)
		}
		for i, line := range linesA {
			if line != linesB[i] {
				lineNo := i + 1
				return fmt.Errorf(
					"file content different: \n%s:%v:%s\n%s:%v:%s",
					aPath, lineNo, line, bPath, lineNo, linesB[i],
				)
			}
		}
		if len(linesA) < len(linesB) {
			return fmt.Errorf("file content different, contains more lines in file %s:%v - %v:\n%s", aPath, len(linesA), len(linesB), strings.Join(linesB[len(linesA):], "\n"))
		}
	}
	return nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
