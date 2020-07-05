package tripletex

import (
	"encoding/json"
	"fmt"
	"time"

	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/bjerkio/tripletex-go/client/session"
	"github.com/bjerkio/trippl/internal/pkg/db"
)

type tokenDataStruct struct {
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
}

func createTokenRequest(consumerToken string, employeeToken string) (*tokenDataStruct, error) {
	fmt.Println("Getting a new token")
	year, month, day := time.Now().Date()
	expirationDate := time.Date(year, month, day+1, 0, 0, 0, 0, time.Now().Location())
	client := apiclient.Default
	sessionReq := &session.TokenSessionCreateCreateParams{
		ConsumerToken:  consumerToken,
		EmployeeToken:  employeeToken,
		ExpirationDate: expirationDate.Format("2006-01-02"),
	}

	res, err := client.Session.TokenSessionCreateCreate(sessionReq.WithTimeout(10 * time.Second))

	if err != nil {
		return nil, err
	}

	return &tokenDataStruct{
		Token:          res.Payload.Value.Token,
		ExpirationDate: expirationDate,
	}, nil
}

func getTokenData(db db.KeyValueStore) (*tokenDataStruct, error) {
	tokenRes, err := db.Get([]byte("tripletex-token"))

	if tokenRes == nil {
		return nil, nil
	}

	var tokenData tokenDataStruct
	err = json.Unmarshal(tokenRes, &tokenData)
	return &tokenData, err
}

func storeTokenData(db db.KeyValueStore, tokenData *tokenDataStruct) error {
	jsonData, err := json.Marshal(&tokenData)
	if err != nil {
		return err
	}

	return db.Set([]byte("tripletex-token"), jsonData)
}

// CreateToken retrieves a new tripletex token
func CreateToken(consumerToken string, employeeToken string, db db.KeyValueStore) (*string, error) {
	tokenData, err := getTokenData(db)
	if err != nil {
		return nil, err
	}

	if tokenData == nil {
		tokenData, err = createTokenRequest(consumerToken, employeeToken)
		if err != nil {
			return nil, err
		}
		err = storeTokenData(db, tokenData)
	}

	if tokenData.ExpirationDate.Before(time.Now()) {
		tokenData, err = createTokenRequest(consumerToken, employeeToken)
		if err != nil {
			return nil, err
		}
		err = storeTokenData(db, tokenData)
	}

	return &tokenData.Token, err
}
