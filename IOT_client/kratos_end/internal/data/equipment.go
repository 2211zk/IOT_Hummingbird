package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos/internal/biz"
)

type wlProductsRepo struct {
	data *Data
	log  *log.Helper
}

func NewWlProductsRepo(data *Data, logger log.Logger) biz.WlProductsRepo {
	return &wlProductsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
