package service

import (
    "net/http"
    "github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "index", struct {
            ID      string `json:"id"`
            Content string `json:"content"`
        }{ID: "15331096", Content: "hello world!"})
    }
}
