package authservice

import "web/data"

func UserFilter() data.Filter {
	return data.MakeFilter(map[string]any{
		"id":       nil,
		"name":     nil,
		"password": nil,
	})
}

func AuthTokenFilter() data.Filter {
	return data.MakeFilter(map[string]any{
		"id":      nil,
		"user_id": nil,
		"token":   nil,
	})
}

var UserModel data.Model = data.Model{
	Name:    "users",
	SQLName: "users",
}

var AuthTokenModel data.Model = data.Model{
	Name:    "auth_tokens",
	SQLName: "auth_tokens",
}

func Init() {
	UserModel.Prepare()
	AuthTokenModel.Prepare()
}
