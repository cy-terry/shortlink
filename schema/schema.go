package schema

import (
	"github.com/Nathan-CY/shortlink/model"
	"github.com/jinzhu/gorm"
)

func AddShortUrl(shortUrl model.ShortURL, db *gorm.DB) error {
	var number []int
	if err := db.
		Table(shortUrl.TableName()).
		Select("count(short_url) as number").
		Where("short_url = ?", shortUrl.ShortUrl).
		Pluck("number", &number).Error; err != nil {
		return err
	}
	if number[0] < 1 {
		if err := db.Create(&shortUrl).Error; err != nil {
			return err
		}
	} else {
		return ErrDataExist
	}
	return nil
}

func QueryShortUrlByLangUrl(shortUrl model.ShortURL, db *gorm.DB) (string, error) {
	var po model.ShortURL
	if err := db.Where("lang_url = ?", shortUrl.LangUrl).Find(&po).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", nil
		} else {
			return "", err
		}
	}
	return po.LangUrl, nil
}

func QueryUnexpiredUrl(shortUrl string, db *gorm.DB) (string, error) {
	var po model.ShortURL
	if err := db.Where("short_url = ?", shortUrl).Find(&po).Error; err != nil {
		return "", err
	}
	return po.LangUrl, nil
}

func MoveDeleteData(shortUrl model.ShortURL, tx *gorm.DB) error {
	var po model.ShortURL
	if err := tx.Model(shortUrl).
		Where("short_url = ?", shortUrl.ShortUrl).
		Where("token = ?", shortUrl.Token).
		Find(&po).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		} else {
			return err
		}
	}
	if err := tx.Delete(po).Error; err != nil {
		return err
	}
	if err := tx.Table("hi_" + po.TableName()).Create(po).Error; err != nil {
		return err
	}
	return nil
}
