package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nathan-CY/shortlink/config"
	"github.com/Nathan-CY/shortlink/model"
	"github.com/Nathan-CY/shortlink/service"
	"github.com/Nathan-CY/shortlink/util"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func RunServer(config *config.WebConf, service *service.Service, auth util.Auth) error {
	handler := initHandler(service)
	initDefFilter(auth)
	router := initRouter(handler, config.RootPath)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(config.Port), 10), router)
	if err != nil {
		return err
	}
	return nil
}

func initRouter(handler *Handler, rootPath string) *httprouter.Router {
	router := httprouter.New()

	router.POST(rootPath+"/sorturl/create", wrapper(handler.createShortURL()))

	router.DELETE(rootPath+"/delete/:sort_hash", wrapper(handler.deleteShortURL()))

	router.GET("/:sort_url", handler.toRedirectAddr())

	return router
}

func wrapper(handle handlerFunc, filters ...filterFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// filter
		if ok := defFilter.doFilter(w, r); !ok {
			return
		}
		if filters != nil {
			for _, v := range filters {
				if ok := v(w, r); !ok {
					return
				}
			}
		}

		// error handle definition
		errorOutput := func(w http.ResponseWriter, err error) {
			data, _ := json.Marshal(model.NewErrorResult(err))
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			_, err1 := fmt.Fprint(w, string(data))
			if err1 != nil {
				log.Println(err1)
			}
		}
		defer func() {
			if e := recover(); e != nil {
				switch e.(type) {
				case error:
					errorOutput(w, e.(error))
				case string:
					errorOutput(w, errors.New(e.(string)))
				default:
					errorOutput(w, errors.New("Server internal error"))
				}
			}
		}()
		// request handler
		if err := handle(w, r, params); err != nil {
			errorOutput(w, err)
		}
	}
}
