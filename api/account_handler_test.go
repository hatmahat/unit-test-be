package api

import (
	mockdb "be-4-rs-crud/src/db/mock"
	db "be-4-rs-crud/src/db/sqlc"
	"be-4-rs-crud/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountAPI(t *testing.T) {
	acc := randomAccount()

	// anonymous struct
	testCases := []struct {
		name            string
		accID           int64
		buildStubs      func(store *mockdb.MockStore)
		responseChecker func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "Status 200",
			accID: acc.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountById(gomock.Any(), gomock.Eq(acc.ID)).
					Times(1).
					Return(acc, nil)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, acc)
			},
		},
		{
			name:  "Status 404",
			accID: acc.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountById(gomock.Any(), gomock.Eq(acc.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, recorder.Code)
				requireBodyMatchResponse(t, recorder.Body, map[string]interface{}{"error": "sql: no rows in result set"})
			},
		},
		{
			name:  "Status 500",
			accID: acc.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountById(gomock.Any(), gomock.Eq(acc.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "Status 400",
			accID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountById(gomock.Any(), gomock.Eq(acc.ID)).
					Times(0)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/account/%d", tc.accID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.responseChecker(t, recorder)
		})
	}
}

func TestCreateAccountAPI(t *testing.T) {
	acc := randomAccount()

	testCases := []struct {
		name            string
		body            gin.H
		buildStubs      func(store *mockdb.MockStore)
		responseChecker func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Status 200",
			body: gin.H{
				"user_name": acc.UserName,
				"balance":   acc.Balance,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateAccountParams{
					UserName: acc.UserName,
					Balance:  acc.Balance,
				}
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Eq(arg)).Times(1).Return(acc, nil)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Status 400",
			body: gin.H{
				"user_name": "",
				"balance":   "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		// {
		// 	name: "Status 500",
		// 	body: gin.H{
		// 		"balance": acc.Balance,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {

		// 		store.EXPECT().
		// 			CreateAccount(gomock.Any(), db.CreateAccountParams{
		// 				// UserName: acc.UserName,
		// 				Balance: acc.Balance,
		// 			}).
		// 			Times(1).
		// 			Return(db.Account{}, sql.ErrConnDone)
		// 	},
		// 	responseChecker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body ke JSON
			data, err := json.Marshal(tc.body)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/account", bytes.NewReader(data))
			assert.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.responseChecker(t, recorder)
		})
	}
}

func randomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(99999, 999999),
		UserName: utils.RandomNumString(20),
		Balance:  utils.RandomNumString(10),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, acc db.Account) {
	data, err := ioutil.ReadAll(body)
	assert.NoError(t, err)

	var gotAcc db.Account
	err = json.Unmarshal(data, &gotAcc)

	assert.NoError(t, err)
	assert.Equal(t, acc, gotAcc)
}

func requireBodyMatchResponse(t *testing.T, body *bytes.Buffer, expected interface{}) {
	data, err := ioutil.ReadAll(body)
	assert.NoError(t, err)

	var gotResponse interface{}
	err = json.Unmarshal(data, &gotResponse)

	assert.NoError(t, err)
	assert.Equal(t, expected, gotResponse)
}
