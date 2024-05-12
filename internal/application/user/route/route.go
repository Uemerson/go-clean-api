package route

import (
	"net/http"

	"github.com/Uemerson/clean-go-api/internal/application/user/handler"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/encrypter"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/helper"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/repository"
	"github.com/Uemerson/clean-go-api/internal/domain/user/service"
	"gorm.io/gorm"
)

func UserRoute(mux *http.ServeMux, db *gorm.DB) {
	e := encrypter.NewEncrypter()
	r := repository.NewRepository(db)
	h := helper.NewHelper()
	s := service.NewService(r, e, h)
	uh := handler.NewUserHandler(s)

	r.Migration()

	mux.HandleFunc("POST /", uh.Add)
	mux.HandleFunc("GET /", uh.Load)
}
