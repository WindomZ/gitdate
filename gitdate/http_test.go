package gitdate

import "testing"

func Test_httpGet(t *testing.T) {
	_, err := httpGet(script_url)
	if err != nil {
		t.Fatal(err)
	}
}
