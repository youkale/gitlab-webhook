package event

import (
	"testing"
	"regexp"
	"fmt"
	"html/template"
	"bytes"
	"strings"
)

func TestRegexp(t *testing.T) {

	a := `1.0.2-test`
	b := `-test`
	m, err := regexp.MatchString(b, a)

	fmt.Println(regexp.Match("H.* ", []byte("Hello World!")))

	t.Log(err)
	t.Log(m)
}

type Data struct {
	Name string
}

func TestTemplate(t *testing.T) {
	d := Data{"1.0.1-dev"}
	a := `show me the money {{ .Name}}`

	tmpl, err := template.New("cmd").Parse(a)

	t.Log(err)
	buf := bytes.NewBuffer(make([]byte, 4096))

	tmpl.Execute(buf, d)

	t.Log(buf.String())

}

func TestStringTrim(t *testing.T) {
	fmt.Println(strings.Trim(`   echo 1.0.1-test`, " "))
}

func TestExeCmd(t *testing.T) {
	execCmd("ls -l")
	execCmd("echo 1.0.1-test")
	execCmd("echo 1.0.1-test")
	execCmd(`   echo 1.0.1-test`)
}
