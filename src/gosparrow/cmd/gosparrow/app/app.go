package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gosparrow/pkg/consts"
	"gosparrow/pkg/errors"
	"gosparrow/pkg/log"
	"gosparrow/pkg/utils"
)

// Session represents the meta info of a Individual call
type Session struct {
	CorrelationID string
	RequestIP     string
}

// HandlerAdapter used to pass handler function to ServeHTTP
type HandlerAdapter struct {
	handler func(*Session, http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func ensureCorrelationID(rw http.ResponseWriter, req *http.Request) (
	correlationID string, err error) {

	rawCorrelationID := req.Header[consts.ReqHeaderCorrelationID]
	if len(rawCorrelationID) == 0 {
		correlationID = utils.UUID()
	} else {
		correlationID = rawCorrelationID[0]
	}
	rw.Header().Set(consts.ReqHeaderCorrelationID, correlationID)
	return
}

// ServeHTTP used to achieve http.HandleFunc interface compliance
func (adapter *HandlerAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	startAt := time.Now()
	status := http.StatusInternalServerError
	var err error
	var correlationID string

	var requestIP string
	rawRequestIP := req.RemoteAddr
	if border := strings.LastIndex(rawRequestIP, ":"); border != -1 {
		requestIP = rawRequestIP[:border]
	}

	renderInternalServerError := func(err error) {
		status = http.StatusInternalServerError
		data := errors.NewAPIError(consts.ErrInternalServerErrorMsg)
		log.Individual(consts.LogLevelError, correlationID, requestIP,
			req.Method, req.RequestURI,
			fmt.Sprintf("ERROR: %s",
				err.Error(),
			),
		)
		renderJSON(rw, status, data)
	}

	correlationID, err = ensureCorrelationID(rw, req)
	if err != nil {
		renderInternalServerError(err)
		return
	}

	log.Individual(consts.LogLevelInfo, correlationID, requestIP, req.Method,
		req.RequestURI, fmt.Sprintf("[req,begin]"))
	defer func() {
		log.Individual(consts.LogLevelInfo, correlationID, requestIP,
			req.Method, req.RequestURI,
			fmt.Sprintf("%d [req,end] %.2fms",
				status,
				time.Now().Sub(startAt).Seconds()*1000,
			),
		)
	}()

	if req.Method == "POST" || req.Method == "PATCH" {
		contentType := req.Header.Get(consts.ReqHeaderContentType)
		if !utils.ContainsStr(contentType, consts.ReqHeaderValueApplicationJSON) {
			log.Individual(consts.LogLevelInfo, correlationID,
				requestIP, req.Method, req.RequestURI,
				fmt.Sprintf("Request content-type %s not supported",
					contentType,
				),
			)
			status = http.StatusUnsupportedMediaType
			err = errors.ErrHandlerUnsupportedMediaType
			renderJSON(rw, status, errors.NewAPIError(err.Error()))
			return
		}
	}

	session := &Session{
		CorrelationID: correlationID,
		RequestIP:     requestIP,
	}

	var data interface{}
	status, data, err = adapter.handler(session, rw, req)
	if err != nil {
		switch status {
		case http.StatusBadRequest,
			http.StatusUnauthorized,
			http.StatusNotAcceptable,
			http.StatusForbidden,
			http.StatusNotFound,
			http.StatusConflict:
			renderJSON(rw, status, errors.NewAPIError(err.Error()))
		default:
			renderInternalServerError(err)
		}
	} else {
		renderJSON(rw, status, data)
	}
}

// Handler used to set handler function
func Handler(f func(*Session, http.ResponseWriter, *http.Request) (
	int, interface{}, error)) *HandlerAdapter {

	return &HandlerAdapter{f}
}

func renderJSON(rw http.ResponseWriter, status int, data interface{}) {
	if data == nil {
		rw.Header().Set(consts.RespHeaderContentLength, strconv.Itoa(0))
		rw.WriteHeader(status)
		rw.Write(nil)
		return
	}

	var result []byte
	var err error
	if result, err = json.MarshalIndent(data, "", "  "); err != nil {
		status = http.StatusInternalServerError
		result = []byte(consts.ErrInternalServerErrorMsg)
		log.App(consts.LogLevelError,
			fmt.Sprintf("Marshal data failed: %s",
				err.Error(),
			),
		)
	}
	result = append(result, '\n')

	// Unescape HTML
	result = bytes.Replace(result, []byte("\\u003c"), []byte("<"), -1)
	result = bytes.Replace(result, []byte("\\u003e"), []byte(">"), -1)
	result = bytes.Replace(result, []byte("\\u0026"), []byte("&"), -1)

	rw.Header().Set(consts.RespHeaderContentType,
		consts.RespHeaderValueApplicationJSON)
	rw.Header().Set(consts.RespHeaderContentLength, strconv.Itoa(len(result)))
	rw.WriteHeader(status)
	rw.Write(result)
}
