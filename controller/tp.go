/*
	Transacction processor
*/
package controller

import "net/http"

type client struct {
	from          string
	to            string
	timestamp     string
	class         string
	ageSuggestive string
}

func (c client) tester() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("client api working"))
	}
}
func (c client) RegisterRoutes() {
	http.HandleFunc("/api/v1/client/test", c.tester())
}
