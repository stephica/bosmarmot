package abci

import (
	"sync"

	"code.monax.io/platform/bosmarmot/release"
	"github.com/tendermint/abci/types"
)

const responseInfoName = "Bosmarmot"

var version = release.Version()

type abciApp struct {
	// State commit mutex,
	mtx              sync.Mutex
	LastBlockHeight  uint64
	LastBlockAppHash []byte
}

func NewApp(genesisAppHash []byte) types.Application {
	return &abciApp{
		LastBlockAppHash: genesisAppHash,
	}
}

func (app *abciApp) Info() types.ResponseInfo {
	return types.ResponseInfo{
		Data:             responseInfoName,
		Version:          version,
		LastBlockHeight:  app.LastBlockHeight,
		LastBlockAppHash: app.LastBlockAppHash,
	}
}

func (app *abciApp) SetOption(key string, value string) string {
	return "No options available"
}

func (app *abciApp) Query(reqQuery types.RequestQuery) (respQuery types.ResponseQuery) {
	respQuery.Log = "Query not support"
	respQuery.Code = types.CodeType_UnknownRequest
	return respQuery
}

func (app *abciApp) CheckTx(tx []byte) types.Result {
	return types.NewResultOK(nil, "")
}

func (app *abciApp) InitChain(validators []*types.Validator) {
	// Could verify agreement on initial validator set here
}

func (app *abciApp) BeginBlock(hash []byte, header *types.Header) {
}

func (app *abciApp) DeliverTx(tx []byte) types.Result {
	return types.NewResultOK(nil, "")
}

func (app *abciApp) EndBlock(height uint64) (respEndBlock types.ResponseEndBlock) {
	return respEndBlock
}

func (app *abciApp) Commit() types.Result {
	app.mtx.Lock()
	defer app.mtx.Unlock()
	return types.NewResultOK(nil, "")
}
