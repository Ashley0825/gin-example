package crud

import (
	"example/model"
	"time"
)

func GetActivity() model.ActivityInDB {
	today := time.Now()
	for _, a := range model.Activities {
		if today.After(a.StartAt) && today.Before(a.EndAt) {
			return a
		}
	}
	return model.ActivityInDB{}
}
