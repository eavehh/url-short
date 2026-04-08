package link

import (
	"net/http"
	"stepik_1/pkg/middleware"
	"stepik_1/pkg/req"
	"stepik_1/pkg/res"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepo *LinkRepository
}

type LinkHandler struct {
	LinkRepo *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepo: deps.LinkRepo,
	}
	router.Handle("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update()))
	router.Handle("DELETE /link/{id}", handler.Delete())
	router.Handle("GET /link/{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepo.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepo.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.SendJson(w, createdLink, 201)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		id, err := strconv.ParseInt(id_str, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.LinkRepo.GetById(uint(id))
		if err != nil {
			res.SendJson(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.SendJson(w, nil, 200)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		println("go to link with hash:", hash)
		link, err := handler.LinkRepo.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id_str := r.PathValue("id")
		id, err := strconv.ParseUint(id_str, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepo.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.SendJson(w, link, 201)
	}
}
