package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bufio"
	"bytes"
)

//var c = config.GetConfig()

func TestJobHook(t *testing.T) {

	serve := httptest.NewServer(&eventHandler{})
	defer serve.Close()

}

func BenchmarkEventHandler_ServeHTTP(b *testing.B) {
	serve := httptest.NewServer(&eventHandler{})
	defer serve.Close()
	for i := 0; i < b.N; i++ {
		if !newRequest(serve.URL, serve.Client()){
			b.Fatal("has error")
		}
	}
}

func newRequest(url string, c *http.Client) bool {
	req, _ := http.NewRequest("POST", url, bufio.NewReader(bytes.NewBuffer([]byte(getJsonString()))))
	req.Header.Set(XGitlabTokenHeaderName, "2575079e53e0605b24b1bd8df2e2f757")
	req.Header.Set(XGitlabEventHeaderName, "Job Hook")
	req.Header.Set("Content-Type", "Application/json")
	resp, _ := c.Do(req)
	return resp.StatusCode == 200
}

func getJsonString() string {
	return `
{
"object_kind": "build",
"ref": "1.0.1-test",
"tag": true,
"before_sha": "2985cc94c12c4cad74c67d0fbeb00db8e5030352",
"sha": "2985cc94c12c4cad74c67d0fbeb00db8e5030352",
"build_id": 2116,
"build_name": "qm-build",
"build_stage": "build",
"build_status": "success",
"build_started_at": "2018-05-16 08:40:27 UTC",
"build_finished_at": "2018-05-16 08:44:12 UTC",
"build_duration": 225.099269,
"build_allow_failure": false,
"project_id": 67,
"project_name": "ellice/hyd-admin",
"user": {
  "id": 2,
  "name": "Sean",
  "email": "tdycxs@gmail.com"
},
"commit": {
  "id": 866,
  "sha": "2985cc94c12c4cad74c67d0fbeb00db8e5030352",
  "message": "test ci\n",
  "author_name": "0x1024",
  "author_email": "",
  "author_url": "",
  "status": "running",
  "duration": null,
  "started_at": "2018-05-16 08:40:26 UTC",
  "finished_at": null
},
"repository": {
  "name": "hyd-admin",
  "url": "",
  "description": "",
  "homepage": "",
  "git_http_url": "",
  "git_ssh_url": "",
  "visibility_level": 0
}
}
`
}
