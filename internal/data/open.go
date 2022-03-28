package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-layout/internal/biz"
)

type openRepo struct {
	data *Data
	log  *log.Helper
}

// NewOpenRepo .
func NewOpenRepo(data *Data, logger log.Logger) biz.OpenRepo {
	return &openRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *openRepo) Save(ctx context.Context, g *biz.Open) (*biz.Open, error) {
	return g, nil
}

func (r *openRepo) Update(ctx context.Context, g *biz.Open) (*biz.Open, error) {
	return g, nil
}

func (r *openRepo) FindByID(context.Context, int64) (*biz.Open, error) {
	return nil, nil
}
