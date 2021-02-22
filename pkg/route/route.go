package route

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	service_B_domain = "http://show:9000"
	v1               = "v1"
	v2               = "v2"
)

func show() *CodeMessage {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", service_B_domain, nil)
	response, _ := client.Do(request)
	if response != nil {
		if response.StatusCode == 200 {
			body, _ := ioutil.ReadAll(response.Body)
			if strings.Contains(string(body), v1) {
				return &CodeMessage{
					Code:    200,
					Version: "v1",
				}
			} else if strings.Contains(string(body), v2) {
				return &CodeMessage{
					Code:    200,
					Version: "v2",
				}
			}
		}
	} else {
		return nil
	}

	return nil
}

type CodeMessage struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
}

func (c *CodeMessage) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}
func print(c *gin.Context) {
	cm := show()
	if cm != nil {
		if cm.Version == v1 {
			c.String(200, v1)
		} else if cm.Version == v2 {
			c.String(200, v2)
		}
	} else {
		c.String(500, "not find page")
	}
}
func Route(r *gin.Engine) {
	r.GET("/show", func(c *gin.Context) {
		print(c)
	})
}
