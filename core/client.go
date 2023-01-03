package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TeoDev1611/davo/utils"
	"github.com/go-resty/resty/v2"
	"github.com/manifoldco/promptui"
)

type DavoPkgInfo struct {
	ReleaseName string
	PreRelease  string
	TagName     string
	URL         string
}

// Helper Functions
func indexOf(element string, data []interface{}) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func GetGitHubInformation(app string) DavoPkgInfo {
	var URL string
	var response map[string]interface{}
	var Davo DavoPkgInfo

	// Check the URL
	if !strings.HasPrefix(app, "https://github.com") && len(strings.Split(app, "/")) == 2 {
		URL = fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", strings.Join(strings.Split(app, "/"), "/"))
	} else if strings.HasPrefix(app, "https://github.com") {
		splitted := strings.Split(strings.ReplaceAll(app, "https://github.com", ""), "/")
		URL = fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", splitted[1], splitted[2])
	} else {
		utils.Error("Not valid URL only valid -> User/RepoName or https://github.com/User/RepoName")
	}

	// Make the Request!
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(URL)
	utils.CheckErrors(err)

	// Unmarshal and make all into map[string]interface{}
	err = json.Unmarshal(resp.Body(), &response)
	utils.CheckErrors(err)

	Davo.URL = fmt.Sprintf("%v", response["assets_url"])
	Davo.PreRelease = fmt.Sprintf("%v", response["prerelease"])
	Davo.ReleaseName = fmt.Sprintf("%v", response["name"])
	Davo.TagName = fmt.Sprintf("%v", response["tag_name"])

	return Davo
}

func DownloadNow(app string) {
	var data []map[string]interface{}
	var name []interface{}
	var date []interface{}
	var count []interface{}
	var download []interface{}

	DavoInfo := GetGitHubInformation(app)
	println(DavoInfo.URL)

	client := resty.New()
	resp, err := client.R().EnableTrace().Get(DavoInfo.URL)
	utils.CheckErrors(err)
	err = json.Unmarshal(resp.Body(), &data)
	utils.CheckErrors(err)

	for _, v := range data {
		name = append(name, v["name"])
		date = append(date, v["created_at"])
		count = append(count, v["download_count"])
		download = append(download, v["browser_download_url"])
	}

	// Prompt for select!
	prompt := promptui.Select{
		Label: "Select Download Option!",
		Items: name,
	}
	_, result, err := prompt.Run()
	utils.CheckErrors(err)

	index := indexOf(result, name)

	fmt.Println(download[index])
}
