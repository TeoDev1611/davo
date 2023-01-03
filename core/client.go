package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TeoDev1611/davo/utils"
	"github.com/go-resty/resty/v2"
	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
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

func downloadFiles(url string, filename string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Davo 🥬! Downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
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
	// Helper variables!
	var data []map[string]interface{}
	var name []interface{}
	var date []interface{}
	var count []interface{}
	var download []interface{}

	// Information from the main function!
	DavoInfo := GetGitHubInformation(app)
	println(DavoInfo.URL)

	// Make the client and the request
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(DavoInfo.URL)
	utils.CheckErrors(err)

	// All information to the map[string]interface{}
	err = json.Unmarshal(resp.Body(), &data)
	utils.CheckErrors(err)

	// Iterate and make the slices!
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

	// Get the information from the selected option!
	index := indexOf(result, name)
	downloadFiles(fmt.Sprintf("%v", download[index]), result)
}
