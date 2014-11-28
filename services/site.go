package services

import (
	"database/sql"
	"github.com/venjiang/mons/models"
)

func (this *CoreService) GetSite() (*models.Site, error) {
	site := models.Site{}
	err := this.DbMap.SelectOne(&site, `select * from site limit 1`)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &site, nil
}
func (this *CoreService) TestSite() *models.Site {
	site := models.Site{}
	err := this.DbMap.SelectOne(&site, `select * from site limit 1`)
	if err == sql.ErrNoRows {
		return nil
	}
	return &site
}
func (this *CoreService) CreateSite(site *models.Site) error {
	return this.DbMap.Insert(site)
}

func (this *CoreService) ExistsSite() bool {
	rows, _ := this.DbMap.SelectInt(`SELECT count(0) FROM site`)
	return rows > 0
}
