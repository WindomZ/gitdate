package gitdate

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const script_url string = "https://gist.githubusercontent.com/WindomZ/18695b33bed478805efe8d384f08b0c4/raw/619230a81d382079e70b873209ae622359beab4c/gitdate"

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func execCommandGitDate(unix int64) error {
	script, err := httpGet(script_url)
	if err != nil {
		return err
	}
	if _, _, err = pipeCommand("git", "stash"); err != nil {
		return err
	}
	defer func() {
		pipeCommand("git", "stash", "pop")
	}()
	_, err = execCommandStatement(
		strings.Replace(script, "$1",
			strconv.FormatInt(unix, 10), -1),
	)
	return err
}
