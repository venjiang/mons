package services

import (
	"github.com/venjiang/mons/models"
)

func (this *CoreService) CreateTag(tag *models.Tag) error {
	return this.DbMap.Insert(tag)
}
