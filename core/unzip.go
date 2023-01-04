package core

import (
	"fmt"
	"os"

	"github.com/TeoDev1611/davo/utils"
	"github.com/evilsocket/islazy/zip"
)

func UnzipFiles(filepath string) {
	files, err := zip.Unzip(filepath, DavoPath())
	utils.CheckErrors(err)
	utils.Info(fmt.Sprintf("Files extracted: %v", files))
	err = os.Remove(filepath)
	utils.CheckErrors(err)
  utils.Info(fmt.Sprintf("Removed the Zip File: %s", filepath))
}
