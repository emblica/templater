package main

import (
	"github.com/flosch/pongo2"
  "encoding/json"
  "io/ioutil"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  dat, err := ioutil.ReadFile(os.Args[1])
  check(err)

  var tpl = pongo2.Must(pongo2.FromFile(os.Args[2]))

  var y map[string]interface{}
  json.Unmarshal([]byte(dat), &y)

  out, err := tpl.Execute(pongo2.Context(y))
  check(err)


	check(ioutil.WriteFile(os.Args[3], []byte(out), 0644))
}
