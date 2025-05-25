package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/shyandsy/shygoctl/api/spec"
	apiutil "github.com/shyandsy/shygoctl/api/util"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/pkg/golang"
	"github.com/shyandsy/shygoctl/util/pathx"
	"github.com/zeromicro/go-zero/core/logx"
)

const tmpFile = "%s-%d"

var (
	tmpDir = path.Join(os.TempDir(), "goctl")
	// VarStringDir describes the directory.
	VarStringDir string
	// VarStringAPI describes the API.
	VarStringAPI string
	// VarStringHome describes the go home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the branch.
	VarStringBranch string
	// VarStringStyle describes the style of output files.
	VarStringStyle  string
	VarBoolWithTest bool
	// VarBoolTypeGroup describes whether to group types.
	VarBoolTypeGroup bool
)

// GoCommand gen go project files from command line

// DoGenProject gen go project files with api file
func DoGenProject(spec *spec.ApiSpec, dir, style string) error {
	//api, err := parser.Parse(apiFile)
	//if err != nil {
	//	return err
	//}

	if err := spec.Validate(); err != nil {
		return err
	}

	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	logx.Must(pathx.MkdirIfNotExist(dir))
	rootPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}

	logx.Must(genEtc(dir, cfg, spec))
	logx.Must(genConfig(dir, cfg, spec))
	logx.Must(genMain(dir, rootPkg, cfg, spec))
	logx.Must(genServiceContext(dir, rootPkg, cfg, spec))
	logx.Must(genTypes(dir, cfg, spec))
	logx.Must(genRoutes(dir, rootPkg, cfg, spec))
	logx.Must(genHandlers(dir, rootPkg, cfg, spec))
	logx.Must(genLogic(dir, rootPkg, cfg, spec))
	logx.Must(genMiddleware(dir, cfg, spec))

	fmt.Println(color.Green.Render("Done."))
	return nil
}

func backupAndSweep(apiFile string) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	_ = os.MkdirAll(tmpDir, os.ModePerm)

	go func() {
		_, fileName := filepath.Split(apiFile)
		_, e := apiutil.Copy(apiFile, fmt.Sprintf(path.Join(tmpDir, tmpFile), fileName, time.Now().Unix()))
		if e != nil {
			err = e
		}
		wg.Done()
	}()
	go func() {
		if e := sweep(); e != nil {
			err = e
		}
		wg.Done()
	}()
	wg.Wait()

	return err
}

func sweep() error {
	keepTime := time.Now().AddDate(0, 0, -7)
	return filepath.Walk(tmpDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		pos := strings.LastIndexByte(info.Name(), '-')
		if pos > 0 {
			timestamp := info.Name()[pos+1:]
			seconds, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil {
				// print error and ignore
				fmt.Println(color.Red.Sprintf("sweep ignored file: %s", fpath))
				return nil
			}

			tm := time.Unix(seconds, 0)
			if tm.Before(keepTime) {
				if err := os.RemoveAll(fpath); err != nil {
					fmt.Println(color.Red.Sprintf("failed to remove file: %s", fpath))
					return err
				}
			}
		}

		return nil
	})
}
