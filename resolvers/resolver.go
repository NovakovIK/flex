package resolvers

import (
	"context"
	"github.com/NovakovIK/flex"
	"github.com/NovakovIK/flex/storage"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	storage *storage.Storage
}

func NewResolver(storage *storage.Storage) *Resolver {
	return &Resolver{storage: storage}
}

func (r *Resolver) Query() flex.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Media(ctx context.Context) ([]*flex.Media, error) {
	data, err := r.storage.MediaDAO.FetchAll()
	if err != nil {
		return nil, err
	}

	var media []*flex.Media
	for i := range data {
		d := &data[i]
		media = append(media, &flex.Media{
			ID:           int(d.ID),
			Name:         d.Name,
			Duration:     int(d.Duration),
			LastModified: int(d.LastModified),
			Status:       d.Status.String(),
		})
	}

	return media, nil
}
func (r *queryResolver) Profiles(ctx context.Context) ([]*flex.Profile, error) {
	panic("not implemented")
}
func (r *queryResolver) ViewingInfo(ctx context.Context, profileID *int) ([]*flex.ProfileViewingInfo, error) {
	panic("not implemented")
}
