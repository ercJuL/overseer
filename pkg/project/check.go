package project

import (
	"github.com/ercJuL/overseer/pkg/utils"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"io/ioutil"
)

const tempVersion = `v1.99999999.999999999`

type checkResult struct {
	AutoUpdateVersionMods []module.Version
	ConstVersionMods      []module.Version
}

func GoModCheck(filePath string, privateDomains ...string) (*checkResult, error) {
	modContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	f, err := modfile.Parse("./test.mod", modContent, func(path, version string) (string, error) {
		if version == "latest" {
			return tempVersion, nil
		}
		return version, nil
	})
	if err != nil {
		return nil, err
	}
	cr := new(checkResult)
	for _, item := range f.Require {
		if !utils.HasAnyPrefix(item.Mod.Path, privateDomains...) {
			continue
		}
		if item.Mod.Version == tempVersion {
			cr.AutoUpdateVersionMods = append(cr.AutoUpdateVersionMods, item.Mod)
		} else {
			cr.ConstVersionMods = append(cr.ConstVersionMods, item.Mod)
		}
	}
	return cr, nil
}
