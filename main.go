package main

import (
	"net/http"
	"github.com/mokeoo/gitlab-webhook/event"
	"github.com/mokeoo/gitlab-webhook/config"
	"io/ioutil"
	"encoding/json"
	"log"
)

const XGitlabEventHeaderName = "X-Gitlab-Event"

const XGitlabTokenHeaderName = "X-Gitlab-Token"

var cfg = config.GetConfig()

type eventHandler struct {
}

func main() {
	http.Handle(cfg.Path, &eventHandler{})
	log.Printf("gitlab web hook server is listen on %s ", cfg.Address)
	log.Fatal(http.ListenAndServe(cfg.Address, nil))
}

func (evt *eventHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	token := request.Header.Get(XGitlabTokenHeaderName)
	if cfg.SecretToken != token {
		log.Println("gitlab secret token is wrong, please check!")
		return
	}
	reqEvent := request.Header.Get(XGitlabEventHeaderName)

	if reqBody, err := ioutil.ReadAll(request.Body); nil == err {
		//log.Printf("method=%s,url=%v,gitlab-event=%s,body=%v\n", request.Method, request.URL, reqEvent, string(reqBody))
		if handler, err := event.GetEventHandler(reqEvent); nil == err {
			json.Unmarshal(reqBody, handler)
			for _, val := range cfg.Settings {
				if reqEvent == val.Event {
					handler.Handle(val)
				}
			}
		}
	}
}
