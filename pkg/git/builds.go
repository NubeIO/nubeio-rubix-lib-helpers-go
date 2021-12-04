package git

import (
	"bufio"
	"fmt"
	"strings"
)

var Targets = struct {
	Amd64 string `json:"amd64"`
	Armv7 string `json:"armv7"`
}{
	Amd64: "amd64",
	Armv7: "armv7",
}

type Git struct {
	Owner        string `json:"owner"`
	Repo         string `json:"repo"`
	Zip          string `json:"zip"`
	Target       string `json:"target"`
	Token        string `json:"token"`         //git token
	BuildVersion string `json:"build_version"` // v0.1.8
	FolderName   string `json:"zip_dir_name"`  // system-0.1.8-c4f9e7a8.amd64.zip
	DownloadPath string `json:"download_path"` //home/user/.tmp-download
	UnzipPath    string `json:"unzip_path"`    //data/rubix-wires/
	IsLocalhost  bool   `json:"is_localhost"`
	URL          string `json:"url"`
}

//ResultSplit used for converting the sting to an object
func (g *Git) ResultSplit(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

const (
	urlBase = iota
	urlReleasesTags
	urlReleases
	urlReleasesLatest

	CurlReleasesTags   //releases tags [v0.1, v0.2]
	CurlReleases       //releases for a version
	CurlReleasesLatest // releases for latest version
	CurlReleaseDownload
)

func (g *Git) buildURL(url int) string {
	base := fmt.Sprintf("https://api.github.com/repos/%s/%s", g.Owner, g.Repo)
	switch url {
	case urlBase:
		return base
	case urlReleasesTags:
		return fmt.Sprintf("%s/tags", base)
	case urlReleases:
		return fmt.Sprintf("%s/releases/%s", base, g.BuildVersion)
	case urlReleasesLatest:
		return fmt.Sprintf("%s/releases/latest", base)
	default:
		return base
	}
}

func (g *Git) BuildCURL(selection int) string {
	switch selection {
	case CurlReleasesTags:
		return fmt.Sprintf("curl -s -H 'Authorization: token %s'  %s | grep -E 'name' | cut -d '\"' -f 4 -", g.Token, g.buildURL(urlReleasesTags))
	case CurlReleases:
		return fmt.Sprintf("curl -s -H 'Authorization: token %s' %s | grep -E 'browser_download_url' | grep %s | cut -d '\"' -f 4 -", g.Token, g.buildURL(urlReleases), g.Target)
	case CurlReleasesLatest:
		return fmt.Sprintf("curl -s -H 'Authorization: token %s' %s | grep -E 'browser_download_url' | grep %s | cut -d '\"' -f 4 -", g.Token, g.buildURL(urlReleasesLatest), g.Target)
	case CurlReleaseDownload:
		return fmt.Sprintf("curl -fsSL -H 'Authorization: token %s' %s -o %s/%s", g.Token, g.URL, g.DownloadPath, g.FolderName)
	default:
		return ""
	}
}

func (g *Git) GetReleasesTags(command string) (string, error) {

	return "", nil
}

//## download all
//# curl -s https://api.github.com/repos/NubeIO/flow-framework/releases/latest | grep -E 'browser_download_url' | grep amd64 | cut -d '"' -f 4 | wget -qi -
//
//## list
//# curl -s https://api.github.com/repos/NubeIO/flow-framework/releases/latest | grep -E 'browser_download_url' | grep amd64  | cut -d '"' -f 4 -
//
//# download a zip
//# curl -fsSL https://github.com/NubeIO/flow-framework/releases/download/v0.1.8/system-0.1.8-c4f9e7a8.amd64.zip  -O
//
//# download with token
//# curl -fsSL -H 'Authorization: token hp_jCSKmylkV937Vy6aEPyOTZNlHhLGHN0Xzld' https://github.com/NubeIO/flow-framework/releases/download/v0.1.8/system-0.1.8-c4f9e7a8.amd64.zip -O

//## list all releases
