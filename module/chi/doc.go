// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

/*
Package mir provider mir.Engine implement backend [Chi](https://github.com/go-chi/chi).

Define handler in struct type like below:

type site struct {
	count    uint32
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"//localhost:8013/articles/{category}/{id:[0-9]+}#GetArticles"`
	chainFunc1 mir.Get   `mir:"/chainfunc1#-ChainFunc"`
	chainFunc2 mir.Get   `mir:"/chainfunc2#GetChainFunc2&ChainFunc"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(rw http.ResponseWriter, r *http.Request) {
	h.count++
	rw.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/{category}/{id:[0-9]+})
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles)
func (h *site) GetArticles(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetArticles"))
}

// ChainFunc1 handler with chain func info.
// Field online middleware info defined in field's tag string (eg: /chainfunc1#-ChainFunc)
func (e *entry) ChainFunc1(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("ChainFunc1"))
}

// GetChainFunc2 handler with chain func info.
// Field online middleware info defined in field's tag string (eg: /chainfunc2#GetChainFunc2&ChainFunc)
func (e *entry) GetChainFunc2(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("GetChainFunc2"))
}

// ChainFunc return field's online middleware
func (e *entry) ChainFunc() chi.Middlewares {
	return chi.Middlewares{
		simpleMiddleware,
		simpleMiddleware,
	}
}

Then register entry such use gin engine:

func main() {
	r := chi.NewRouter()

	// Instance a mir engine to register handler for mux router by mir
	mirE := muxE.Mir(r)
	mir.Register(mirE, &site{})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8013", r))
}
*/

package chi
