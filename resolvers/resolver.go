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

func (r *queryResolver) Media(ctx context.Context, id *int) ([]*flex.Media, error) {
	var data []storage.Media
	var err error

	if id != nil {
		data, err = r.storage.MediaDAO.FetchByID(int64(*id))
	} else {
		data, err = r.storage.MediaDAO.FetchAll()
	}

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
func (r *queryResolver) Profiles(ctx context.Context, id *int) ([]*flex.Profile, error) {
	var data []storage.Profile
	var err error

	if id != nil {
		data, err = r.storage.ProfileDAO.FetchByID(int64(*id))
	} else {
		data, err = r.storage.ProfileDAO.FetchAll()
	}

	if err != nil {
		return nil, err
	}

	var profiles []*flex.Profile
	for i := range data {
		d := &data[i]
		profiles = append(profiles, &flex.Profile{
			ID:   int(d.ID),
			Name: d.Name,
		})
	}

	return profiles, nil
}
func (r *queryResolver) ViewingInfo(ctx context.Context, mediaID *int, profileID *int) ([]*flex.ProfileViewingInfo, error) {
	var data []storage.ProfileViewingInfo
	var err error

	if mediaID != nil && profileID != nil {
		data, err = r.storage.ProfileViewingInfoDAO.FetchByMediaIDAndProfileID(int64(*mediaID), int64(*profileID))
	} else if mediaID != nil {
		data, err = r.storage.ProfileViewingInfoDAO.FetchByMediaID(int64(*mediaID))
	} else if profileID != nil {
		data, err = r.storage.ProfileViewingInfoDAO.FetchByProfileID(int64(*profileID))
	} else {
		data, err = r.storage.ProfileViewingInfoDAO.FetchAll()
	}

	if err != nil {
		return nil, err
	}
	var viewingInfo []*flex.ProfileViewingInfo
	for i := range data {
		d := &data[i]
		viewingInfo = append(viewingInfo, &flex.ProfileViewingInfo{
			MediaID:   int(d.MediaID),
			ProfileID: int(d.ProfileID),
			TimePoint: int(d.ProfileID),
			Timestamp: int(d.Timestamp),
		})
	}

	return viewingInfo, nil
}
