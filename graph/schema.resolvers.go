package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/yush/go-orders-graphql-api/graph/generated"
	"github.com/yush/go-orders-graphql-api/graph/model"
)

func (r *mutationResolver) CreateCreator(ctx context.Context, name string, email string) (*model.Creator, error) {
	
	var creator = model.Creator{
		ID:    strconv.Itoa(rand.Int()),
		Name:  name,
		Email: email,
	}
	r.DB.Create(&creator)
	fmt.Println(creator, r.DB)
	return &creator, nil
}

func (r *mutationResolver) CreateAudio(ctx context.Context, input model.NewAudio) (*model.Audio, error) {
	var creator *model.Creator
	r.DB.First(&creator, input.CreatorID)
	newVideo := model.Audio{
		ID:          strconv.Itoa(rand.Int()),
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		CreatorID:   input.CreatorID,
		Creator:     creator,
	}
	err := r.DB.Create(&newVideo).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &newVideo, nil
}

func (r *queryResolver) Audios(ctx context.Context) ([]*model.Audio, error) {
	var videos []*model.Audio
	err := r.DB.Find(&videos).Error
	for _, video := range videos {
		var creator *model.Creator
		r.DB.Find(&creator, video.CreatorID)
		video.Creator = creator
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return videos, nil
}

func (r *queryResolver) Audio(ctx context.Context, input string) (*model.Audio, error) {
	var video *model.Audio
	r.DB.First(&video, input)
	var creator *model.Creator
	r.DB.Find(&creator, video.CreatorID)
	video.Creator = creator
	return video, nil
}

func (r *queryResolver) Creators(ctx context.Context) ([]*model.Creator, error) {
	var creators []*model.Creator
	r.DB.Find(&creators)
	return creators, nil
}

func (r *queryResolver) Creator(ctx context.Context, input string) (*model.Creator, error) {
	var creator *model.Creator
	r.DB.Find(&creator, input)
	return creator, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
