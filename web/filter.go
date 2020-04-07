package web

import (
	"github.com/Nathan-CY/shortlink/util"
	"net/http"
)

type filterFunc func(w http.ResponseWriter, r *http.Request) bool

var defFilter = &routerFilter{
	filterChain: make([]filterFunc, 0),
}

type routerFilter struct {
	filterChain []filterFunc
}

func (rf *routerFilter) doFilter(w http.ResponseWriter, r *http.Request) bool {
	if rf.filterChain != nil && len(rf.filterChain) > 0 {
		for _, fun := range rf.filterChain {
			if ok := fun(w, r); !ok {
				return ok
			}
		}
	}
	return true
}

func (rf *routerFilter) link(filterFunc func(w http.ResponseWriter, r *http.Request) bool) *routerFilter {
	rf.filterChain = append(rf.filterChain, filterFunc)
	return rf
}

func initDefFilter(auth util.Auth) {
	defFilter.link(func(w http.ResponseWriter, r *http.Request) bool {
		return auth.Action()
	})
}
