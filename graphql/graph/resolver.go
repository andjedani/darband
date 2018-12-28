package graph

import (
	"context"
	"tange/bigv"
	"tange/common"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignIn(ctx context.Context, input UserAuth) (Token, error) {
	panic("not implemented")
}
func (r *mutationResolver) SignUp(ctx context.Context, input UserAuth) (Token, error) {
	db := common.GetDB()
	var model bigv.UserModel
	err := db.Where(&bigv.UserModel{Username: input.Username}).First(&model).Error
	if err != nil {

	}
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (User, error) {
	panic("not implemented")
}
