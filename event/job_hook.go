package event

import (
	"github.com/mokeoo/gitlab-webhook/config"
	"os/exec"
	"strings"
	"html/template"
	"bytes"
	"strconv"
	"io"
	"fmt"
	"log"
)

type JobHook struct {
	BeforeSha         string     `json:"before_sha"`
	BuildAllowFailure bool       `json:"build_allow_failure"`
	BuildDuration     float64    `json:"build_duration"`
	BuildFinishedAt   string     `json:"build_finished_at"`
	BuildID           int        `json:"build_id"`
	BuildName         string     `json:"build_name"`
	BuildStage        string     `json:"build_stage"`
	BuildStartedAt    string     `json:"build_started_at"`
	BuildStatus       string     `json:"build_status"`
	Commit            Commit     `json:"commit"`
	ObjectKind        string     `json:"object_kind"`
	ProjectID         int        `json:"project_id"`
	ProjectName       string     `json:"project_name"`
	Ref               string     `json:"ref"`
	Repository        Repository `json:"repository"`
	Sha               string     `json:"sha"`
	Tag               bool       `json:"tag"`
	User              User       `json:"user"`
}

func (j *JobHook) Handle(setting config.Setting) {
	//工程名
	isNameSame := j.Repository.Name == setting.ProjectName
	//匹配tag
	m := strings.LastIndex(j.Ref, setting.Ref) != -1
	//stage
	isEqualsBuild := j.BuildStage == setting.BuildStage
	//build result
	isEqualsStatus := j.BuildStatus == setting.BuildStatus
	//build name
	isEqualsBuildName := j.BuildName == setting.BuildName
	if isNameSame && m && isEqualsBuild && isEqualsStatus && isEqualsBuildName {
		for i, cmd := range setting.Command {
			c := cmd
			if tmpl, err := template.New(strconv.Itoa(i)).Parse(cmd); nil == err {
				buf := bytes.NewBufferString("")
				tmpl.Execute(buf, j)
				c = buf.String()
			}
			execCmd(c)
		}
	}
}

func execCmd(c string) {
	sc := strings.Trim(c, " ")
	fmt.Printf("~ $ %s\n", c)
	cmd := exec.Command("/bin/bash", "-c", sc)
	stderr, _ := cmd.StderrPipe()
	std, _ := cmd.StdoutPipe()
	go func() {
		printCmdOut(stderr)
	}()
	go func() {
		printCmdOut(std)
	}()
	ch := make(chan error, 1)
	if err := cmd.Start(); err != nil {
		log.Printf("PATH %v", cmd.Env)
		ch <- err
	}
	go func() {
		ch <- cmd.Wait()
	}()
	e := <-ch
	if nil != e {
		fmt.Printf("%v\n", e.Error())
	}
}

func printCmdOut(read io.ReadCloser) {
	buf := make([]byte, 512)
	for {
		n, e := read.Read(buf)
		if e == io.EOF {
			break
		} else {
			fmt.Printf("%v", string(buf[:n]))
		}
	}
	fmt.Println()
}
