package graph

import (
	"context"
	"tange/bigv"
	"tange/common"

	log "github.com/sirupsen/logrus"
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
	db := common.GetDB()
	var model []bigv.UserModel
	db.Where(&bigv.UserModel{Username: input.Username}).Find(&model)
	if len(model) == 1 {
		user := model[0]
		if user.Authenticate(input.Password) == nil {
			token := Token{
				Msg:     "Doroste",
				Success: true,
			}
			return token, nil
		} else {
			token := Token{
				Msg:     "Ghalate",
				Success: false,
			}
			return token, nil
		}
	}
	token := Token{
		Msg:     "Nist",
		Success: false,
	}
	return token, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, input UserAuth) (Token, error) {
	log.Info(input.Password)
	db := common.GetDB()
	var model []bigv.UserModel
	db.Where(&bigv.UserModel{Username: input.Username}).Find(&model)
	log.Info(input.Username)
	if len(model) == 0 {
		user, err := bigv.CreateUser(input.Username, input.Password)
		if err != nil {
			panic("panic")
		}
		token := Token{
			Msg:     "shod" + user.PasswordHash,
			Success: true,
		}
		return token, nil
	} else {
		token := Token{
			Msg:     "nashod" + model[0].Username,
			Success: false,
		}
		return token, nil
	}

}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (User, error) {
	panic("not implemented")
}
