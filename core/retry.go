package core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jxcia/go-garden/core/log"
	"github.com/opentracing/opentracing-go"
)

func retryAnalyze(retry string) ([]int, error) {
	retrySlice := make([]int, 0)
	arr := strings.Split(retry, "/")
	if len(arr) == 0 {
		return []int{}, errors.New("config retry format error")
	}
	for _, sec := range arr {
		s, err := strconv.Atoi(sec)
		if err != nil {
			return []int{}, errors.New("config retry format error")
		}
		retrySlice = append(retrySlice, s)
	}

	retrySlice = append(retrySlice, 0)

	return retrySlice, nil
}

func (g *Garden) retryGo(service, action string, retry []int, nodeIndex int, span opentracing.Span, route routeCfg, request *req, body MapData, rpcArgs, rpcReply interface{}) (int, string, http.Header, error) {
	code := httpOk
	result := infoSuccess
	addr := ""
	var err error
	var header http.Header

	for i, r := range retry {
		atomic.AddInt64(&g.services[service].Nodes[nodeIndex].Waiting, 1)

		if route.Type == "http" {
			addr, err = g.getServiceHttpAddr(service, nodeIndex)
			if err != nil {
				code = httpFail
				break
			}
			addr = "http://" + addr + route.Path
			code, result, header, err = g.requestServiceHttp(span, addr, request, body, route.Timeout)
		} else if route.Type == "rpc" {
			addr, err = g.getServiceRpcAddr(service, nodeIndex)
			if err != nil {
				code = httpFail
				break
			}
			action = capitalize(action)
			err = rpcCall(span, addr, service, action, rpcArgs, rpcReply, route.Timeout)
			if err != nil {
				code = httpFail
			}
		}

		atomic.AddInt64(&g.services[service].Nodes[nodeIndex].Waiting, -1)

		if err != nil {
			log.Error("callService", err)
			g.addFusingQuantity(g.services[service].Nodes[nodeIndex].Addr + "/" + service + "/" + action)

			// call timeout don't retry
			if strings.Contains(err.Error(), "Timeout") || strings.Contains(err.Error(), "deadline") {
				err = errors.New(fmt.Sprintf("Call %s %s %s timeout", route.Type, service, action))
				return code, infoTimeout, nil, err
			}

			// call 404 don't retry
			if code == httpNotFound {
				return code, infoNotFound, nil, err
			}

			if i == len(retry)-1 {
				return code, infoServerError, nil, err
			}
			time.Sleep(time.Millisecond * time.Duration(r))
			continue
		}

		break
	}

	atomic.AddInt64(&g.services[service].Nodes[nodeIndex].Finish, 1)

	return code, result, header, err
}
