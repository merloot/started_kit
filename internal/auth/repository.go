package auth

import (
	"context"
	"log"
	"started_kit/internal/entities"
	"started_kit/pkg/dbcontext"
)

type Repository interface {
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func (r repository) Get(ctx context.Context, id string) (entities.User, error) {
	var user entities.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) Create(ctx context.Context, user entities.User) error {
	return r.db.With(ctx).Model(&user).Insert()
}

func (r repository) Update(ctx context.Context, user entities.User) error {
	return r.db.With(ctx).Model(&user).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	user, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("users").Row(&count)
	return count, err
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entities.User, error) {
	var users []entities.User
	err := r.db.With(ctx).Select().OrderBy("createdAt").Offset(int64(offset)).Limit(int64(limit)).All(&users)
	return users, err
}
