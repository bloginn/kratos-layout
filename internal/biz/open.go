package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Open struct {
	Hello string
}

type OpenRepo interface {
	Save(context.Context, *Open) (*Open, error)
	Update(context.Context, *Open) (*Open, error)
	FindByID(context.Context, int64) (*Open, error)
}

// OpenUseCase is a Open use case.
type OpenUseCase struct {
	repo OpenRepo
	log  *log.Helper
}

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.Errorf(10001, "user not found", "user not found")
)

// NewOpenUseCase new a Open use case.
func NewOpenUseCase(repo OpenRepo, logger log.Logger) *OpenUseCase {
	return &OpenUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OpenUseCase) Hello(ctx context.Context, name string) (*Open, error) {
	uc.log.WithContext(ctx).Infof("name:%s", name)
	if name != "word" {
		return nil, ErrUserNotFound
	}

	return &Open{Hello: "hello " + name}, nil
}
