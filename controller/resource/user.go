package resource

type UserResource struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Cipher    string `json:"password"`
	AuthToken AuthTokenResource
}

func (res UserResource) ColsIdent() map[string]interface{} {
	return map[string]interface{}{
		"id":          "ID",
		"name":        "Name",
		"password":    "Cipher",
		"auth_tokens": "AuthToken",
	}
}

type AuthTokenResource struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func (res AuthTokenResource) ColsIdent() map[string]interface{} {
	return map[string]interface{}{
		"id":      "ID",
		"user_id": "UserID",
		"token":   "Token",
	}
}
