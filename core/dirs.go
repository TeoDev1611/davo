package core

import (
	"os"
	"path"

	"github.com/TeoDev1611/davo/utils"
	"github.com/i582/cfmt/cmd/cfmt"
)

func DavoPath() string {
	home, err := os.UserHomeDir()
	utils.CheckErrors(err)
	return path.Join(home, "davo", "bin")
}

func DavoSetup() {
	if _, err := os.Stat(DavoPath()); os.IsNotExist(err) {
		os.MkdirAll(DavoPath(), os.ModePerm)
		utils.Info("Created succesfully the path")
	}

	cfmt.Printf(`
Davo 🥬!
{{For the correct work you need add this}}::underline

{{- zsh}}::green|bold -> {{.zshrc}}::magenta|underline

  {{export PATH=$PATH:$HOME/davo/bin}}::cyan

{{- fish}}::green|bold -> {{$HOME/.config/fish/config.fish}}::magenta|underline

  {{fish_add_path $HOME/davo/bin}}::cyan

{{- other}}::green|bold

  Append this directory to {{%s}}::cyan
`, DavoPath())
}
