// Command generate writes the generated layer of the PropRaven Go SDK from
// openapi.json:
//
//	go run ./internal/cmd/generate            # from the repo root
//	go run ./internal/cmd/generate -spec path/to/openapi.json -out . -readme README.md
//
// Output (all gofmt'd, deterministic): gen_types.go (components/schemas),
// gen_<group>.go per x-sdk-group (service, methods, params, responses),
// gen_client.go (the Client struct), and the method table in README.md
// between the GENERATED METHODS markers. Stale gen_*.go files are removed.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	specPath := flag.String("spec", "openapi.json", "OpenAPI 3.1 spec to read")
	outDir := flag.String("out", ".", "directory of the propraven package")
	readme := flag.String("readme", "README.md", "README to update between the GENERATED METHODS markers (empty to skip)")
	pkg := flag.String("package", "propraven", "Go package name")
	flag.Parse()

	data, err := os.ReadFile(*specPath)
	if err != nil {
		fatal(err)
	}
	res, err := Generate(data, *pkg)
	if err != nil {
		fatal(err)
	}
	for _, w := range res.Warnings {
		fmt.Fprintln(os.Stderr, "generate: warning:", w)
	}
	if err := writeFiles(*outDir, res.Files); err != nil {
		fatal(err)
	}
	if *readme != "" {
		if err := updateReadme(*readme, res.MethodTable); err != nil {
			fatal(err)
		}
	}
	fmt.Fprintf(os.Stderr, "generate: %d operations, %d files\n", res.Operations, len(res.Files))
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "generate:", err)
	os.Exit(1)
}

// writeFiles writes files and deletes gen_*.go files no longer produced.
func writeFiles(dir string, files map[string][]byte) error {
	existing, err := filepath.Glob(filepath.Join(dir, "gen_*.go"))
	if err != nil {
		return err
	}
	for _, p := range existing {
		if _, keep := files[filepath.Base(p)]; !keep {
			if err := os.Remove(p); err != nil {
				return err
			}
		}
	}
	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		p := filepath.Join(dir, n)
		if old, err := os.ReadFile(p); err == nil && string(old) == string(files[n]) {
			continue
		}
		if err := os.WriteFile(p, files[n], 0o644); err != nil {
			return err
		}
	}
	return nil
}

const (
	readmeBegin = "<!-- BEGIN GENERATED METHODS (go run ./internal/cmd/generate) -->"
	readmeEnd   = "<!-- END GENERATED METHODS -->"
)

func updateReadme(path, table string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	s := string(data)
	i := strings.Index(s, readmeBegin)
	j := strings.Index(s, readmeEnd)
	if i < 0 || j < i {
		fmt.Fprintln(os.Stderr, "generate: README has no GENERATED METHODS markers; skipped")
		return nil
	}
	out := s[:i+len(readmeBegin)] + "\n" + table + s[j:]
	if out == s {
		return nil
	}
	return os.WriteFile(path, []byte(out), 0o644)
}
