package main

import (
    "fmt"
    "io/ioutil"
    "path"

    "github.com/robertkrimen/otto"
)

type ScriptEngine struct {
    VM *otto.Otto
    Dir string
}

type ScriptContext struct {
    Requestor string
    Command string
    Args string
}


func NewScriptEngine(dir string)(se ScriptEngine) {
    se.VM = otto.New()
    se.Dir = dir

    return
}


func (s *ScriptEngine) Run (f string, c ScriptContext) (output otto.Value, err error) {
    var scriptContent []byte

    script := path.Join(s.Dir, fmt.Sprintf("%s.js", f))
    if scriptContent, err = ioutil.ReadFile(script); err != nil {
        return
    }

    s.VM.Set("Context", c)
    return s.VM.Run(scriptContent)

}
