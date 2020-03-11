package rpc

import (
	sdk "github.com/irisnet/irishub-sdk-go/types"
)

type Random interface {
	sdk.Module
	Generate(request RandomRequest) (reqID string, err sdk.Error)
	QueryRandom(reqID string) (RandomInfo, sdk.Error)
	QueryRequests(height int64) ([]RequestRandom, sdk.Error)
}

type RandomRequest struct {
	sdk.BaseTx
	BlockInterval uint64
	Callback      EventGenerateRandomCallback
}

type EventGenerateRandomCallback func(reqID, randomNum string, err sdk.Error)

// Rand represents a random number with related data
type RandomInfo struct {
	RequestTxHash string `json:"request_tx_hash"` // the original request tx hash
	Height        int64  `json:"height"`          // the height of the block used to generate the random number
	RandomNum     string `json:"random_num"`      // the actual random number
}

// RequestRandom represents a request for a random number
type RequestRandom struct {
	Height   int64  `json:"height"`   // the height of the block in which the request tx is included
	Consumer string `json:"consumer"` // the request address
	TxHash   string `json:"tx_hash"`  // the request tx hash
}