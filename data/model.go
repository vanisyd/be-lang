package data

import (
	"fmt"
)

func (model *Model) TableName() string {
	return model.SQLName
}

func (model *Model) GetName() string {
	return fmt.Sprintf("%s", model.Name)
}
