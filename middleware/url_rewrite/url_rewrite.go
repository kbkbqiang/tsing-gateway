package set_header

import (
	"net/http"
	"strings"

	"github.com/dxvgef/tsing-gateway/global"
)

// url rewrite
type URLRewrite struct {
	Rule map[string]string `json:"rule"`
}

// 新建中间件实例
func New(config string) (*URLRewrite, error) {
	var instance URLRewrite
	err := instance.UnmarshalJSON(global.StrToBytes(config))
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func (self *URLRewrite) GetName() string {
	return "url_rewrite"
}

// 中间件行为
func (self *URLRewrite) Action(resp http.ResponseWriter, req *http.Request) (bool, error) {
	for k := range self.Rule {
		if req.URL.Path == k {
			req.URL.Path = strings.Replace(req.URL.Path, k, self.Rule[k], 1)
			req.RequestURI = req.URL.RequestURI()
		}
	}
	return true, nil
}
