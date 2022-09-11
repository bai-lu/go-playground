package api

import (
	"encoding/json"

	"github.com/bai-lu/go-playground/utils"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

const URLPrefix = "/grafana/api"

var JsonResponse = utils.JsonResponse{}

type PostData struct {
	Target string `json:"target"`
}

func Run() {
	m := macaron.Classic()
	RegisterBase(m)
	RegisterApi(m)
	m.Run()
}
func RegisterBase(m *macaron.Macaron) {

	// 所有GET方法，自动注册HEAD方法
	m.SetAutoHead(true)
	// 404错误
	m.NotFound(func(ctx *macaron.Context) string {
		return JsonResponse.Failure(utils.NotFound, "您访问的页面不存在")
	})
	// 50x错误
	m.InternalServerError(func(ctx *macaron.Context) string {
		return JsonResponse.Failure(utils.ServerError, "服务器内部错误, 请稍后再试")
	})
}

func RegisterApi(m *macaron.Macaron) {
	m.SetURLPrefix(URLPrefix)

	// GET / with 200 status code response. Used for "Test connection" on the datasource config page.
	m.Get("/", func() string {
		return "Hello world!"
	})

	// POST /search to return available metrics.
	m.Post("/search", binding.Bind(PostData{}), Search)

	// POST /query to return panel data or annotations.
	m.Get("/query", binding.Bind(PostData{}), Query)
}

func Search() string {
	return ""

}

func Query() string {
	// return JsonResponse.Success(utils.SuccessCode, utils.SuccessContent, LoadFeedJson())
	result, err := json.Marshal(NewMachineData())
	if err != nil {
		panic(err)
	}
	return string(result)
}
