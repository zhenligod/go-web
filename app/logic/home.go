package logic

import (
	"errors"

	"github.com/zhenligod/go-web/app/model"

	"github.com/jinzhu/gorm"
	"github.com/zhenligod/thingo/mysql"
)

// HomeLogic home logic.
type HomeLogic struct {
	BaseLogic
}

// GetData 模拟数据库查询
func (h *HomeLogic) GetData(name string) (map[string]interface{}, error) {
	db, err := mysql.GetDbObj("default")
	if err != nil {
		//log.Println("db connection error: ", err)
		return nil, errors.New("db connection error")
	}

	user := &model.User{}
	err = db.Where("name = ?", name).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	//log.Println("user: ", user)

	return map[string]interface{}{
		"user": user,
	}, nil
}
