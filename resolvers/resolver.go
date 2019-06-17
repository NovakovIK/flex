package resolvers

import (
	"context"
	"github.com/NovakovIK/flex"
	"github.com/NovakovIK/flex/storage"
	log "github.com/sirupsen/logrus"
)

type Resolver struct {
	storage *storage.Storage
}

func NewResolver(storage *storage.Storage) *Resolver {
	return &Resolver{storage: storage}
}

func (r *Resolver) Query() flex.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() flex.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *queryResolver) Media(ctx context.Context, id *int) ([]*flex.Media, error) {
	var data []storage.Media
	var err error

	if id != nil {
		data, err = r.storage.MediaDAO.FetchByID(*id)
	} else {
		data, err = r.storage.MediaDAO.FetchAll()
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var media []*flex.Media
	for i := range data {
		d := &data[i]
		media = append(media, &flex.Media{
			ID:        d.ID,
			Name:      d.Name,
			Path:      d.Path,
			Duration:  d.Duration,
			Created:   d.Created,
			Status:    d.Status.String(),
			TimePoint: d.TimePoint,
			LastSeen:  d.LastSeen,
			Thumbnail: string(d.Thumbnail),
			Width:     d.Width,
			Heigth:    d.Height,
			Size:      d.Size,
		})
	}

	return media, nil
}

func (r *mutationResolver) UpdateMedia(ctx context.Context, input flex.MediaInput) (*flex.Media, error) {
	media, err := r.storage.MediaDAO.Update(input.ID, input.Name, input.LastSeen, input.TimePoint)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &flex.Media{
		ID:        media.ID,
		Name:      media.Name,
		Path:      media.Path,
		Duration:  media.Duration,
		Created:   media.Created,
		Status:    media.Status.String(),
		TimePoint: media.TimePoint,
		LastSeen:  media.LastSeen,
		Thumbnail: string(media.Thumbnail),
		Width:     media.Width,
		Heigth:    media.Height,
		Size:      media.Size,
	}, nil
}
