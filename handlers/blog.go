package handlers

import (
	"net/http"
	
	"portfolio_reimagined/db"
)

type BlogHandler struct {
	Queries *db.Queries
}

func (h *BlogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogs, err := h.Queries.ListBlogs(ctx)
	if err != nil {
		http.Error(w, "Failed to load blogs", http.StatusInternalServerError)
		return
	}

	for _, blog := range blogs {
		_, _ = w.Write([]byte(blog.Title + "\n"))
	}
}
