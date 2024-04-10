package retriever

import (
	"context"
	"errors"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"

	"gorm.io/gorm"
)

type SpaceRetriever interface {
	Create(c context.Context, table *model.Space) error
	Query(c context.Context, alias string) (*model.Space, error)
}

type spaceRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewSpaceRetriever(db *gorm.DB, cache *cache.Cache) SpaceRetriever {
	return &spaceRetriever{
		db:    db,
		cache: cache,
	}
}

func (s spaceRetriever) Create(c context.Context, table *model.Space) error {

	err := s.db.WithContext(c).Create(table).Error
	return err

}

/*

 */
func (s spaceRetriever) Query(c context.Context, alias string) (*model.Space, error) {
	var space model.Space
	var token model.Token

	// 4.9test
	// var users []model.User
	// err := s.db.Find(&users)
	// if err == nil {

	// }
	// for _, user := range users {
	// 	fmt.Println(user, "123")
	// }
	//

	if err := s.db.WithContext(c).First(&space, "alias = ?", "glaxe").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	// if err := s.db.WithContext(c).First(&token, "id = ?", space.TokenID).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		// 未找到记录
	// 		return nil, err
	// 	}
	// 	// 发生了其他错误
	// 	return nil, err
	// }
	space.Token = token

	return &space, nil
}