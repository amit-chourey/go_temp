package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Authtoken string
	Username string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type databaseInterface interface {
	GetUserLoginDetails(username string) * LoginDetails
	GetUserCoins(username string) * CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockDB{}
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}