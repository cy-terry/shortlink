package service

import (
	"errors"
	"github.com/Nathan-CY/shortlink/model"
	"github.com/Nathan-CY/shortlink/schema"
	"github.com/Nathan-CY/shortlink/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Service struct {
	Converter util.URLConverter
	UniqueID  util.UniqueID
	DB        *gorm.DB
}

func (s Service) CreateShortURL(urlApply *model.UrlApply) (string, error) {
	// If there is a direct return
	shortURL := model.ShortURL{
		ID:         s.UniqueID.Generate(),
		LangUrl:    urlApply.LangUrl,
		Duration:   urlApply.Duration,
		Token:      urlApply.Token,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	// If a long address exists, the short address corresponding to the long link is returned directly
	langUrl, err := schema.QueryShortUrlByLangUrl(shortURL, s.DB)
	if err != nil || langUrl != "" {
		return langUrl, err
	}
	// Calculate the available short address groups based on long addresses
	// Select an unused short address in the available group as the short address of this long address
	shortUrls := s.Converter.ToShortURL(urlApply.LangUrl)
	err = schema.TxAction(s.DB, func(tx *gorm.DB) error {
		for _, v := range shortUrls {
			shortURL.ShortUrl = v
			err := schema.AddShortUrl(shortURL, tx)
			if err == nil {
				return nil
			}
			if !schema.IsDataExistError(err) {
				return err
			}
		}
		return errors.New("Insufficient resources")
	})
	return shortURL.ShortUrl, err
}

func (s Service) DeleteShortURL(shortUrl string, token string) error {
	return schema.TxAction(s.DB, func(tx *gorm.DB) error {
		return schema.MoveDeleteData(model.ShortURL{
			ShortUrl: shortUrl,
			Token:    token,
		}, s.DB)
	})
}

func (s Service) UnexpiredAddr(shortUrl string) string {
	if langUrl, err := schema.QueryUnexpiredUrl(shortUrl, s.DB); err == nil {
		return langUrl
	}
	return ""
}
