package service

import (
	"forum/server/data"
	"forum/server/shareddata"
)

type InfoService struct {
	InfoData data.InfoData
}

func (s *InfoService) GetInfoData(userUID string) (shareddata.InfoData, error) {
	var info shareddata.InfoData

	// Get username and user id
	username, id := GetUser(s.InfoData.Db, userUID)
	// if id = 0 that means the user doesn't exist
	info.Authorize = id != 0
	info.Username = username

	// Get Categories
	categories, err := s.InfoData.GetCategories()
	if err != nil {
		return shareddata.InfoData{}, err
	}
	info.Categories = categories

	return info, nil
}