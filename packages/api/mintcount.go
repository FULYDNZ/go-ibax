package api

import (
	"github.com/IBAX-io/go-ibax/packages/conf"
	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/converter"
	"github.com/IBAX-io/go-ibax/packages/model"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (m Mode) getMintCountHandler(w http.ResponseWriter, r *http.Request) {
	logger := getLogger(r)
	params := mux.Vars(r)
	ret := model.Response{}
	blockID := converter.StrToInt64(params["id"])
	if blockID == 0 {
		mc := &model.MintCount{}
		f, err := mc.Get(blockID)
		if err != nil {
			logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("getting Key for wallet")
			ret.ReturnFailureString(err.Error())
			JsonCodeResponse(w, &ret)
			return
		}
		if f {
			ret.Return(mc, model.CodeSuccess)
			JsonCodeResponse(w, &ret)
			return
		} else {
			ret.ReturnFailureString("not find")
			JsonCodeResponse(w, &ret)
			return
		}
	} else {
		ret.ReturnFailureString("PoolPub.Enable false")
		JsonCodeResponse(w, &ret)
		return
	}
}