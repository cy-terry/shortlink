package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nathan-CY/shortlink/config"
	"github.com/Nathan-CY/shortlink/model"
	"github.com/Nathan-CY/shortlink/service"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

type Handler struct {
	s *service.Service
}

func initHandler(service *service.Service) *Handler {
	return &Handler{
		s: service,
	}
}

var urlRegexp, _ = regexp.Compile(urlRegexpString)
var urlRegexpString = `^((https|http)?://)?(([0-9a-z_!~*'().&=+$%-]+: )?[0-9a-z_!~*'().&=+$%-]+@)?(([0-9]{1,3}\.){3}[0-9]{1,3}|([0-9a-z_!~*'()-]+\.)*([0-9a-z][0-9a-z-]{0,61})?[0-9a-z]\.[a-z]{2,6})(:[0-9]{1,4})?((/?)|(/[0-9a-z_!~*'().;?:@&=+$,%#-]+)+/?)`

func (h *Handler) createShortURL() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
		data, _ := ioutil.ReadAll(r.Body)
		var urlApply model.UrlApply
		if err := json.Unmarshal(data, &urlApply); err != nil {
			return err
		}
		isLegal := urlRegexp.MatchString(urlApply.LangUrl)
		if !isLegal {
			return errors.New("Lang url illegal")
		}
		urlApply.Token = r.Header.Get("Access-Token")
		shortUrl, err := h.s.CreateShortURL(&urlApply)
		if err != nil {
			return err
		}

		result, _ := json.Marshal(&model.Result{
			Code: 200,
			Msg:  "Application is successful",
			Data: struct {
				ShortUrl string `json:"short_url"`
				LangUrl  string `json:"lang_url"`
			}{ShortUrl: config.Conf.Serve.ShortPrefix + shortUrl, LangUrl: urlApply.LangUrl},
			Timestamp: time.Now().Unix(),
		})
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_, _ = fmt.Fprint(w, string(result))
		return nil
	}
}

func (h *Handler) deleteShortURL() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
		err := h.s.DeleteShortURL(params.ByName("sort_hash"), r.Header.Get("Access-Token"))
		if err == nil {
			result, _ := json.Marshal(&model.Result{
				Code:      200,
				Msg:       "Delete is successful",
				Timestamp: time.Now().Unix(),
			})
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			_, _ = fmt.Fprint(w, string(result))
		}
		return err
	}
}

func (h *Handler) toRedirectAddr() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		langUrl := h.s.UnexpiredAddr(params.ByName("sort_url"))
		if langUrl != "" {
			http.Redirect(w, r, langUrl, http.StatusFound)
		} else {
			// TODO 404 Response
		}
	}
}
