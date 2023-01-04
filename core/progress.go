package core

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/TeoDev1611/davo/utils"
	"github.com/schollz/progressbar/v3"
)

func DownloadFileWithProgressBar(url string, filename string) {
	req, err := http.NewRequest("GET", url, nil)
	utils.CheckErrors(err)
	resp, err := http.DefaultClient.Do(req)
	utils.CheckErrors(err)
	defer resp.Body.Close()

	f, _ := os.OpenFile(path.Join(DavoPath(), filename), os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Davo 🥬! Downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}
