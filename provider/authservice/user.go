package authservice

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
	"web/data"
	"web/server/kind"
)

func Auth(filter data.Filter) (token string) {
	dbConnection := data.DBConnection()
	query := data.Query{
		Model:        UserModel,
		DBConnection: dbConnection,
	}

	rowsCount := query.Select(UserModel).Filter(filter).Get()
	if rowsCount == 0 {
		return
	}

	tokenQuery := data.Query{
		Model:        AuthTokenModel,
		DBConnection: dbConnection,
	}
	tokenFilter := AuthTokenFilter()
	tokenFilter["user_id"] = query.Data[0]["id"]
	tokens := tokenQuery.Select(AuthTokenModel).Filter(tokenFilter).Get()
	if tokens == 0 {
		tokenQuery = data.Query{
			Model:        AuthTokenModel,
			DBConnection: dbConnection,
		}
		tokenQuery.Insert(map[string]interface{}{
			"user_id": tokenFilter["user_id"],
			"token":   getToken(tokenFilter["user_id"].(int)),
		})

		result := tokenQuery.Exec(dbConnection)
		tokenID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
			return
		}

		tokenFilter := AuthTokenFilter()
		tokenFilter["id"] = tokenID
		tokens = tokenQuery.Select(AuthTokenModel).Filter(tokenFilter).Get()
		if tokens == 0 {
			log.Fatal(kind.ERROR_UNKNOWN)
			return
		}
	}

	token = tokenQuery.Data[0]["token"].(string)

	return
}

func getToken(userID int) (token string) {
	currentTime := time.Now()
	tokenCipher := sha256.Sum256([]byte(fmt.Sprint(userID, currentTime)))
	token = hex.EncodeToString(tokenCipher[:])

	return
}
