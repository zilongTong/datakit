// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package http

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

type httpRouteInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var httpRouteList = make(map[string]*httpRouteInfo)

func RegHTTPHandler(method, path string, handler http.HandlerFunc) {
	method = strings.ToUpper(method)
	if _, ok := httpRouteList[method+path]; ok {
		l.Warnf("failed to register %s %s by handler %s to HTTP server because of exists",
			method, path, getFunctionName(handler, '/'))
	} else {
		httpRouteList[method+path] = &httpRouteInfo{
			Method:  method,
			Path:    path,
			Handler: func(c *gin.Context) { handler(c.Writer, c.Request) },
		}
	}
}

func RegGinHandler(method, path string, handler gin.HandlerFunc) {
	method = strings.ToUpper(method)
	if _, ok := httpRouteList[method+path]; ok {
		l.Warnf("failed to register %s %s by handler %s to HTTP server because of exists",
			method, path, getFunctionName(handler, '/'))
	} else {
		httpRouteList[method+path] = &httpRouteInfo{
			Method:  method,
			Path:    path,
			Handler: handler,
		}
	}
}

func applyHTTPRoute(router *gin.Engine) {
	for _, routeInfo := range httpRouteList {
		method := routeInfo.Method
		path := routeInfo.Path
		handler := routeInfo.Handler

		l.Infof("register %s %s by handler %s to HTTP server", method, path, getFunctionName(handler, '/'))

		switch method {
		case http.MethodPost:
			router.POST(path, ginWraper(reqLimiter), handler)
		case http.MethodGet:
			router.GET(path, ginWraper(reqLimiter), handler)
		case http.MethodHead:
			router.HEAD(path, ginWraper(reqLimiter), handler)
		case http.MethodPut:
			router.PUT(path, ginWraper(reqLimiter), handler)
		case http.MethodPatch:
			router.PATCH(path, ginWraper(reqLimiter), handler)
		case http.MethodDelete:
			router.DELETE(path, ginWraper(reqLimiter), handler)
		case http.MethodOptions:
			router.OPTIONS(path, ginWraper(reqLimiter), handler)
		}
	}
}

func getFunctionName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}
