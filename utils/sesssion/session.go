package sesssion

import (
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/models"
	"github.com/wangluozhe/requests/url"
	"strings"
)

var (
	Pxy = ""
)

type Client struct {
	Session *requests.Session
}

func NewClient(w config.WebConfig) *Client {
	if !config.IsLoad {
		config.InitConfig()
		config.IsLoad = true
	}
	session := requests.NewSession()
	session.Proxies = Pxy
	session.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	if w.Authorization != "" {
		if strings.Contains(w.Authorization, "Bearer") {
			session.Headers.Set("Authorization", w.Authorization)
		} else {
			session.Headers.Set("Authorization", "Bearer "+w.Authorization)
		}
	}
	if w.Proxy != "" {
		session.Proxies = w.Proxy
	} else if config.Conf.Proxy != "" {
		session.Proxies = config.Conf.Proxy
	}
	for k, v := range w.Headers {
		session.Headers.Set(k, v)

	}
	return &Client{
		Session: session,
	}
}

func (c *Client) Do(method, url string, req *url.Request) (*models.Response, error) {
	return c.Session.Request(method, url, req)
}
