package bootstrap

import (
	"web/provider/authservice"
	"web/provider/vocabulary"
)

func Init() {
	vocabulary.Init()
	authservice.Init()
}
