package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/8mamo10/bankrupt/db/mock"
	db "github.com/8mamo10/bankrupt/db/sqlc"
	"github.com/8mamo10/bankrupt/token"
	"github.com/8mamo10/bankrupt/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateTransferAPI(t *testing.T) {
	amount := int64(10)

	user1, _ := randomUser(t)
	user2, _ := randomUser(t)
	account1 := randomAccount(user1.Username)
	account2 := randomAccount(user2.Username)

	account1.Currency = util.USD
	account2.Currency = util.USD

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{}
	t.Run("OK", func(t *testing.T) {
		// ctrl
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		// stub
		store := mockdb.NewMockStore(ctrl)
		store.EXPECT().
			GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
			Times(1).
			Return(account1, nil)
		store.EXPECT().
			GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
			Times(1).
			Return(account2, nil)
		arg := db.TransferTxParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		}
		store.EXPECT().
			TransferTx(gomock.Any(), gomock.Eq(arg)).
			Times(1)
			// server
		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()
		// request body
		body := gin.H{
			"from_account_id": account1.ID,
			"to_account_id":   account2.ID,
			"amount":          amount,
			"currency":        util.USD,
		}
		data, err := json.Marshal(body)
		require.NoError(t, err)
		// http request
		url := "/transfers"
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		require.NoError(t, err)
		// auth
		addAutohrization(t, request, server.tokenMaker, authorizationTypeBearer, account1.Owner, time.Minute)
		// response
		server.router.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}
