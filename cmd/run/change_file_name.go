package run

import (
	"os"
	"path/filepath"
	"strings"
)

func ChangeFileName() {
	srcDir := "/Users/chisato/code/ios_rule_script/rule/Clash"
	destDir := "/Users/chisato/.config/clash-verge/Rules/rule0"

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".list" {
			destPath := filepath.Join(destDir, filepath.Base(path))
			err = os.Rename(path, destPath)
			if err != nil {
				return err
			}
			newDestPath := strings.TrimSuffix(destPath, filepath.Ext(destPath))
			return os.Rename(destPath, newDestPath)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
