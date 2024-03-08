package middleware

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/auth"
	pErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func NewCheckAuth(uc auth.Usecase, log *zap.Logger) func(h http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//sessionCookie, err := r.Cookie(constants.SessionName)
			//if err != nil {
			//	log.Debug("Failed to get session cookie", zap.Error(err))
			//	pHTTP.HandleError(w, r, pErrors.ErrSessionNotFound)
			//	return
			//}
			//
			//id, authToken, err := parseSessionCookie(sessionCookie)
			//if err != nil {
			//	pHTTP.HandleError(w, r, pErrors.ErrBadSessionCookie)
			//	return
			//}
			//
			//userID, err := uc.CheckAuth(r.Context(), id, authToken)
			//if err != nil {
			//	pHTTP.HandleError(w, r, err)
			//	return
			//}

			var userID int64 = 1
			ctx := context.WithValue(r.Context(), ContextUserID, userID)
			//ctx = context.WithValue(ctx, ContextAuthToken, authToken)

			h(w, r.WithContext(ctx))
		}
	}
}

func parseSessionCookie(c *http.Cookie) (int64, string, error) {
	tmp := strings.Split(c.Value, "$")
	if len(tmp) != 2 {
		return 0, "", pErrors.ErrBadSessionCookie
	}

	id, err := strconv.ParseInt(tmp[0], 10, 64)
	if err != nil {
		return 0, "", err
	}

	return id, c.Value, nil
}
