package kunpeng

import (
	"encoding/json"
	"github.com/medasz/kunpeng/config"
	"github.com/medasz/kunpeng/plugin"
	_ "github.com/medasz/kunpeng/plugin/go"
	_ "github.com/medasz/kunpeng/plugin/json"
	"github.com/medasz/kunpeng/util"
)

var VERSION string

type greeting string

func (g greeting) Check(taskJSON string) []map[string]interface{} {
	var task plugin.Task
	json.Unmarshal([]byte(taskJSON), &task)
	return plugin.Scan(task)
}

func (g greeting) GetPlugins() []map[string]interface{} {
	return plugin.GetPlugins()
}

func (g greeting) SetConfig(configJSON string) {
	config.Set(configJSON)
}

func (g greeting) ShowLog() {
	config.SetDebug(true)
}

func (g greeting) GetVersion() string {
	return VERSION
}

func (g greeting) StartBuffer() {
	util.Logger.StartBuffer()
}

func (g greeting) GetLog(sep string) string {
	return util.Logger.BufferContent(sep)
}

func ShowLog() {
	config.SetDebug(true)
}

func StartBuffer() {
	util.Logger.StartBuffer()
}

var Greeter greeting
