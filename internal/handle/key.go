package handle

import (
	"github.com/zerodays/sistem-users/internal/config"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/util"
	"net/http"
)

// SigningKeyHandle serves key that is used to check
// the signature of access token.
func SigningKeyHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	err := util.EncodePubKey(config.Login.SigningPublicKey, w)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}
}
