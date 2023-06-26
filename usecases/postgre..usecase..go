package usecase

import "tes/models"

func (uc *usecase) InsertDB(req models.ItemList) error {
	err := uc.postgre.Insert(req)
	if err != nil {
		return err
	}

	return nil
}
