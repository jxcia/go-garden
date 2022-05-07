package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/jxcia/go-garden/core/log"
	"github.com/opentracing/opentracing-go"
)

// Request datatype
type req struct {
	Method   string  `json:"method"`
	Url      string  `json:"url"`
	UrlParam string  `json:"urlParam"`
	ClientIp string  `json:"clientIp"`
	Headers  MapData `json:"headers"`
	Body     MapData `json:"body"`
}

func (g *Garden) callService(span opentracing.Span, service, action string, request *req, body MapData, args, reply interface{}) (int, string, http.Header, error) {
	s := g.cfg.Routes[service]
	if len(s) == 0 {
		return httpNotFound, infoNotFound, nil, errors.New("service not found")
	}
	route := s[action]
	if (route.Type != "http" && route.Type != "rpc") ||
		(route.Type == "http" && len(route.Path) == 0) ||
		(route.Type == "rpc" && (args == nil || reply == nil)) {
		return httpNotFound, infoNotFound, nil, errors.New("service route not found")
	}

	serviceAddr, nodeIndex, err := g.selectService(service)
	if err != nil {
		return httpNotFound, infoNotFound, nil, err
	}

	// service limiter
	if route.Limiter != "" {
		second, quantity, err := limiterAnalyze(route.Limiter)
		if err != nil {
			log.Error("limiter", err)
		} else if !g.limiterInspect(serviceAddr+"/"+service+"/"+action, second, quantity) {
			span.SetTag("break", "service limiter")
			return httpNotFound, infoServerLimiter, nil, errors.New("server limiter")
		}
	}

	// service fusing
	if route.Fusing != "" {
		second, quantity, err := g.fusingAnalyze(route.Fusing)
		if err != nil {
			log.Error("fusingAnalyze", err)
		} else if !g.fusingInspect(serviceAddr+"/"+service+"/"+action, second, quantity) {
			span.SetTag("break", "service fusing")
			return httpNotFound, infoServerFusing, nil, errors.New("server fusing")
		}
	}

	// service call retry
	retry, err := retryAnalyze(g.cfg.Service.CallRetry)
	if err != nil {
		log.Error("retryAnalyze", err)
		retry = []int{0}
	}

	code, result, header, err := g.retryGo(service, action, retry, nodeIndex, span, route, request, body, args, reply)

	return code, result, header, err
}

func (g *Garden) requestServiceHttp(span opentracing.Span, url string, request *req, body MapData, timeout int) (int, string, http.Header, error) {
	client := &http.Client{
		Timeout: time.Millisecond * time.Duration(timeout),
	}

	// encapsulation request body
	//var s string
	// debug
	//for k, v := range body {
	//	s += fmt.Sprintf("%v=%v&", k, v)
	//}
	//s = strings.Trim(s, "&")
	s, _ := json.Marshal(&body)
	r, err := http.NewRequest(request.Method, url, bytes.NewReader(s))
	if err != nil {
		return httpFail, "", nil, err
	}

	// New request request
	for k, v := range request.Headers {
		r.Header.Add(k, v.(string))
	}
	// Add the body format header
	ct := fmt.Sprint(request.Headers["Content-Type"])
	r.Header.Set("Content-Type", ct)
	//	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Increase calls to the downstream service security validation key
	r.Header.Set("Call-Key", g.cfg.Service.CallKey)
	//debug
	log.Info("debug", request.Headers)
	log.Info("debug", body)
	log.Info("debug", request)
	// add request opentracing span header
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	res, err := client.Do(r)
	if err != nil {
		return httpFail, "", nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != httpOk {
		return res.StatusCode, "", nil, errors.New("http status " + strconv.Itoa(res.StatusCode))
	}
	body2, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return httpFail, "", nil, err
	}
	return httpOk, string(body2), res.Header, nil
}

// CallRpc call other service rpc method
func (g *Garden) CallRpc(span opentracing.Span, service, action string, args, reply interface{}) error {
	_, _, _, err := g.callService(span, service, action, nil, nil, &args, &reply)
	if err != nil {
		return err
	}
	return nil
}
