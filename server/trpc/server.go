package trpc

import (
	"context"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/website/Bilibili"
	"github.com/Tontonnow/ddddhmlist/website/Friday"
	"github.com/Tontonnow/ddddhmlist/website/Hami"
	"github.com/Tontonnow/ddddhmlist/website/IQ"
	"github.com/Tontonnow/ddddhmlist/website/KKTV"
	"github.com/Tontonnow/ddddhmlist/website/LETV"
	"github.com/Tontonnow/ddddhmlist/website/Litv"
	"github.com/Tontonnow/ddddhmlist/website/MGTV"
	"github.com/Tontonnow/ddddhmlist/website/MIGU"
	"github.com/Tontonnow/ddddhmlist/website/MyVideo"
	"github.com/Tontonnow/ddddhmlist/website/MytvSuper"
	"github.com/Tontonnow/ddddhmlist/website/QQTV"
	"github.com/Tontonnow/ddddhmlist/website/VIKI"
	"github.com/Tontonnow/ddddhmlist/website/VIU"
	"github.com/Tontonnow/ddddhmlist/website/XiGUA"
	"github.com/Tontonnow/ddddhmlist/website/YOUKUTV"
	"github.com/Tontonnow/ddddhmlist/website/yangshipin"
	"github.com/google/uuid"
	"runtime/debug"
	"strings"
	"trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/filter"
	trpcHttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

var ErrCode = map[int]string{
	0: "Success",
	1: "内部错误,请排查日志,任务ID%s",
	2: "链接格式可能有误,任务ID%s",
	3: "地区限制,服务器无对应地区代理,任务ID%s",
}

type DdddListServiceImpl struct {
	UnimplementedDDDDhm
}

func (s *DdddListServiceImpl) DdddList(ctx context.Context, req *server.Request) (*server.Response, error) {
	var rsp = &server.Response{
		Code:    0,
		Message: ErrCode[0],
	}
	log.Info("DdddList", "requestId", ctx.Value("requestId"))
	data, code := GetMateInfo(ctx, req)
	if code != 0 {
		rsp.Code = int32(code)
		rsp.Message = fmt.Sprintf(ErrCode[code], ctx.Value("requestId"))
		return rsp, nil
	}
	rsp.Video = data
	return rsp, nil

}
func (s *DdddListServiceImpl) Hello(ctx context.Context, req *server.HelloRequest) (*server.HelloResponse, error) {
	rsp := &server.HelloResponse{
		Msg: req.GetMsg(),
	}
	return rsp, nil
}
func GetMateInfo(ctx context.Context, req *server.Request) (r *server.Data, code int) {
	url := req.Url
	var f func(ctx context.Context, sharerUrl string) (r *server.Data, code int)
	r = &server.Data{}
	if url == "" {
		code = 2
		return
	}
	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			log.Error("GetMateInfo", "err", err, "stack", stack)
			code = 1
			return
		}
	}()
	if strings.Contains(url, "friday") {
		f = Friday.GetMateInfo
	} else if strings.Contains(url, "hamivideo") {
		f = Hami.GetMateInfo
	} else if strings.Contains(url, ".iq.") {
		f = IQ.GetIqMateInfo
	} else if strings.Contains(url, ".iqiyi.") {
		f = IQ.GetMateInfo
	} else if strings.Contains(url, "kktv.me") {
		f = KKTV.GetMateInfo
	} else if strings.Contains(url, ".le.") {
		f = LETV.GetMateInfo
	} else if strings.Contains(url, ".litv.") {
		f = Litv.GetMateInfo
	} else if strings.Contains(url, ".mgtv.") {
		f = MGTV.GetMateInfo
	} else if strings.Contains(url, "miguvideo") {
		f = MIGU.GetMateInfo
	} else if strings.Contains(url, ".mytvsuper.") {
		f = MytvSuper.GetMateInfo
	} else if strings.Contains(url, "myvideo.net.tw") {
		f = MyVideo.GetMateInfo
	} else if strings.Contains(url, "v.qq.com") {
		f = QQTV.GetMateInfo
	} else if strings.Contains(url, "wetv.") {
		f = QQTV.GetWeTvMateInfo
	} else if strings.Contains(url, ".viki.") {
		f = VIKI.GetMateInfo
	} else if strings.Contains(url, ".viu.") {
		f = Viu.GetMateInfo
	} else if strings.Contains(url, "wetv.") {
		f = QQTV.GetWeTvMateInfo
	} else if strings.Contains(url, ".ixigua.") || strings.Contains(url, "lvdetail") {
		f = XiGUA.GetMateInfo
	} else if strings.Contains(url, "yangshipin") {
		f = yangshipin.GetMateInfo
	} else if strings.Contains(url, "youku.com") || strings.Contains(url, "youku.tv") {
		f = YOUKUTV.GetMateInfo
	} else if strings.Contains(url, ".bilibili.") {
		f = Bilibili.GetMateInfo
	} else {
		return r, 2
	}
	return f(ctx, url)
}
func Filter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (interface{}, error) {
	Head := trpcHttp.Head(ctx)
	requestId := ""
	if Head != nil {
		requestId = Head.Request.Header.Get("X-Request-Id")
	}
	if requestId == "" || strings.Count(requestId, "-") != 4 {
		requestId = uuid.New().String()
	}
	ctx = context.WithValue(ctx, "requestId", requestId)
	trpc.SetMetaData(ctx, "requestId", []byte(requestId))
	rsp, err := next(ctx, req)
	return rsp, err
}

func Init() {
	filter.Register("ddddFilter", Filter, nil)
	LogFunc := func(ctx context.Context, req, rsp interface{}) string {
		return fmt.Sprintf(", requestId:%s", ctx.Value("requestId"))
	}
	var debuglogOption []debuglog.Option
	debuglogOption = append(debuglogOption, debuglog.WithLogFunc(LogFunc))
	debuglogOption = append(debuglogOption, debuglog.WithEnableColor(true))
	filter.Register("SimpleLogFunc",
		debuglog.ServerFilter(debuglogOption...),
		debuglog.ClientFilter(debuglogOption...))
	s := trpc.NewServer()
	RegisterDDDDhmService(s.Service("DDDDhm.DdddListService"),
		&DdddListServiceImpl{})
	RegisterDDDDhmService(s.Service("DDDDhm.DdddListServiceHTTP"),
		&DdddListServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
