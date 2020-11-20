package output_plugin_file_d

import (
	"fmt"

	"github.com/ozonru/file.d/fd"
	"github.com/ozonru/file.d/pipeline"
)

type Plugin struct {
	controller pipeline.OutputPluginController
	outFn      func(event *pipeline.Event)
}

type Config struct {
}

func init() {
	fd.DefaultPluginRegistry.RegisterOutput(&pipeline.PluginStaticInfo{
		Type:    "new_stdout",
		Factory: Factory,
	})
	fmt.Println("NEW STDOUT - INIT")
}

func Factory() (pipeline.AnyPlugin, pipeline.AnyConfig) {
	return &Plugin{}, &Config{}
}

func (p *Plugin) Start(_ pipeline.AnyConfig, params *pipeline.OutputPluginParams) {
	p.controller = params.Controller
	fmt.Println("NEW STDOUT - START")
}

func (p *Plugin) Stop() {
}

func (p *Plugin) Out(event *pipeline.Event) {
	//fmt.Println(event.Root.EncodeToString())
	fmt.Println("NEW STDOUT - OUT")
	p.controller.Commit(event)
}
