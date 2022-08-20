// 与falcon交互的模型

package inspector

import "github.com/bai-lu/go-playground/define"

type Config struct {
	domain string
	api_1  string
	api_2  string
	api_3  string
}

type falconInspector struct {
	Config Config
}

func init() {
	var _ define.Inspector = &falconInspector{}
}

func init() {
	var _ define.Inspector = &falconInspector{}
}

func NewFalconInspector(config Config) *falconInspector {
	return &falconInspector{}
}

func (fi *falconInspector) Run() error {
	return nil
}

func (fi *falconInspector) Stop() error {
	return nil
}

func (fi *falconInspector) Progress() (uint, error) {
	return 0, nil
}
