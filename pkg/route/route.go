package route

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	ShowServiceDomin = ""
)

func setshowServiceDomin() string {
	ShowServiceDomin = os.Getenv("ShowServiceDomin")
	if ShowServiceDomin == "" {
		panic("ShowServiceDomin没有设置")
	}
	return "http://" + ShowServiceDomin
}

type CodeMessage struct {
	Code    int `json:"code"`
	Version int `json:"version"`
}

func (c *CodeMessage) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}
func show(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, setshowServiceDomin())
}
func Route(r *gin.Engine) {
	r.GET("/show", func(c *gin.Context) {
		show(c)
	})
}
