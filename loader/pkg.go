package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

// trimExtension removes the file extension if exists.
func trimExtension(path string) string {
	if ext := filepath.Ext(path); ext != "" {
		base := filepath.Base(path)
		path = path[:len(path)-len(base)]
	}
	return path
}

// packagePath returns the github package path from any given path,
// e.g. path/to/github.com/your-repo/your-pkg returns github.com/your-repo/your-pkg
// If your package is not hosted on github, you may need to override $PKG to
// set the prefix of your package.
func packagePath(prefix, path string) string {
	path = trimExtension(path)
	path = strings.TrimRight(path, "/")
	idx := strings.Index(path, prefix)
	return path[idx:]
}

// packageName returns the base package name.
func packageName(prefix, path string) string {
	return filepath.Base(packagePath(prefix, path))
}

func loadPackage(path string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports,
	}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		panic(fmt.Errorf("loader: failed to load package: %v", err))
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}
	return pkgs[0]
}
