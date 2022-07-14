package filter

import (
	"context"
	"net/http"
	"strconv"
)

const (
	OptionsContextKey = "filter_options"
)

func Middleware(h http.HandlerFunc, defaultLimit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitFromQuery := r.URL.Query().Get("limit")

		limit := defaultLimit
		var limitParseError error
		if limitFromQuery != "" {
			if limit, limitParseError = strconv.Atoi(limitFromQuery); limitParseError != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad limit"))
				return
			}
		}
		optionsI := NewOptions(limit)
		ctx := context.WithValue(r.Context(), OptionsContextKey, optionsI)
		r = r.WithContext(ctx)

		h(w, r)
	}
}
