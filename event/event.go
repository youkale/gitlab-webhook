package event

import (
	"github.com/mokeoo/gitlab-webhook/config"
	"errors"
)

var handlerRegister = make(map[string]HandleEvent)

func init() {
	handlerRegister["Job Hook"] = &JobHook{}
}

type HandleEvent interface {
	Handle(setting config.Setting)
}

func GetEventHandler(eventHook string) (HandleEvent, error) {
	handler := handlerRegister[eventHook]
	if nil == handler {
		return nil, errors.New("")
	}
	return handler, nil
}

type Repository struct {
	Description     string `json:"description"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	Homepage        string `json:"homepage"`
	Name            string `json:"name"`
	VisibilityLevel int    `json:"visibility_level"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Commit struct {
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	Duration    int    `json:"duration"`
	FinishedAt  string `json:"finished_at"`
	ID          int    `json:"id"`
	Message     string `json:"message"`
	Sha         string `json:"sha"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
}

type User struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}
