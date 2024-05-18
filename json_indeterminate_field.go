package main

import (
	"github.com/artschu23/go-exp/indeterminate"
	"github.com/artschu23/go-exp/internal/scenes/iot"
	"github.com/google/uuid"
)

type NaiveGroupOverview struct {
	iot.Group

	AutoUpdate interface{} `json:"auto_update"`
}

type GroupOverview struct {
	iot.Group

	AutoUpdate indeterminate.Field[bool] `json:"auto_update"`
}

func GetGroupOverview(_ iot.Group) GroupOverview {
	devices := []iot.Device{
		{
			ID:         uuid.New(),
			Name:       "neko",
			AutoUpdate: true,
		},
		{
			ID:         uuid.New(),
			Name:       "usagi",
			AutoUpdate: false,
		},
	}

	group := iot.Group{
		ID:      uuid.New(),
		Name:    "Chiikawa Group",
		Devices: devices,
	}

	groupOverview := GroupOverview{
		// select required fields
		Group: iot.Group{
			ID:   group.ID,
			Name: group.Name,
		},

		// initialize indeterminate fields
		AutoUpdate: indeterminate.Field[bool]{
			Value: group.Devices[0].AutoUpdate,
		},
	}

	// logic to check if all devices are set to update automatically
	for _, device := range group.Devices {
		if device.AutoUpdate != group.Devices[0].AutoUpdate {
			groupOverview.AutoUpdate.IsIndeterminate = true
			break
		}
	}

	return groupOverview
}
