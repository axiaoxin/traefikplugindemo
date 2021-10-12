// Package traefikplugindemo Traefik插件示例
// 给请求响应头添加 resp:xxx
package traefikplugindemo

import (
	"context"
	"net/http"

	"github.com/iancoleman/strcase"
)

// Config the plugin configuration.
type Config struct {
	// resp header值的字符串风格：snake, camel
	ValueStrCase string
	// resp header的默认值
	DefaultValue string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		ValueStrCase: "",
		DefaultValue: "",
	}
}

// HeaderResp a plugin.
type HeaderResp struct {
	next http.Handler
	name string
	conf *Config
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HeaderResp{
		next: next,
		name: name,
		conf: config,
	}, nil
}

func (h *HeaderResp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 默认返回头 resp:axiaoxin
	val := h.conf.DefaultValue
	// 如果请求头带有 x-resp，则返回头 resp:x-resp的值
	if customVal := req.Header.Get("x-resp"); customVal != "" {
		val = customVal
	}

	// 使用外部依赖包，发布插件时需要采用vendor模式
	switch h.conf.ValueStrCase {
	case "camel":
		val = strcase.ToCamel(val)
	case "snake":
		val = strcase.ToSnake(val)
	}

	// 设置返回头
	rw.Header().Add("resp", val)

	// 继续后续请求处理
	h.next.ServeHTTP(rw, req)
}
