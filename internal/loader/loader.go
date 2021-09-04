// Package loader used for compile the target source code and generate SSA for it.
package loader

import (
	"path/filepath"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"

	"github.com/bytevet/godzilla/internal/logger"
)

// LoadProject compile the target source code and generate SSA for it.
func LoadProject(path string) (ret *ssa.Program, err error) {
	if path, err = filepath.Abs(path); err != nil {
		return nil, err
	}

	workDir := filepath.Dir(path)

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedExportsFile |
			packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
		Logf: func(format string, args ...interface{}) {
			logger.Debug(format, args...)
		},

		// Switch the work dir into the project path.
		// It's here for hisotrical reason. In older version of Go,
		// packages.Load may modify go.mod in work dir.
		Dir: workDir,
		// Disable loading testcases.
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil, err
	}

	ret, _ = ssautil.AllPackages(pkgs, ssa.NaiveForm|ssa.BareInits)
	return
}
