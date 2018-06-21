package httprouter

import (
	"net/http"
	"testing"
)

func TestGroup(t *testing.T) {
	var (
		index, hello bool
	)
	gs := NewGroup("/v1",
		NSGroup("/api",
			NSRouter("/index", "get",
				func(w http.ResponseWriter, r *http.Request, _ Params) {
					index = true
				}),
			NSRouter("/hello", "POST",
				func(w http.ResponseWriter, r *http.Request, _ Params) {
					hello = true
				}),
		),
	)
	r := AddGroups(gs)
	w := new(mockResponseWriter)
	ri, _ := http.NewRequest("GET", "/v1/api/index", nil)
	rh, _ := http.NewRequest("POST", "/v1/api/hello", nil)
	r.ServeHTTP(w, ri)
	r.ServeHTTP(w, rh)
	if !index {
		t.Error("/v1/api/index testing error!!")
	}
	if !hello {
		t.Error("/v1/api/hello testing errror!")
	}

}
