package lion

import "net/http"

type Handlefunc func(c *Context)

// handlefunc

type HandleFunc func(w http.ResponseWriter, r *http.Request)
