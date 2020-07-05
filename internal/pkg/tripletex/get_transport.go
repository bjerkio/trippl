package tripletex

import (
	"github.com/go-openapi/runtime"

	apiclient "github.com/bjerkio/tripletex-go/client"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
)

func GetTransport(config config.TripplConfig, db db.KeyValueStore) (*httptransport.Runtime, error) {
	token, err := CreateToken(config.ConsumerToken, config.EmployeeToken, db)
	if err != nil {
		return nil, err
	}

	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth("0", *token)
	r.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	return r, nil
}
