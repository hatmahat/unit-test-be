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

	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// store := mockdb.NewMockStore(ctrl) // mock yg udah dibuat

	// // build stubs
	// store.EXPECT().
	// 	GetAccountById(gomock.Any(), gomock.Eq(acc.ID)).
	// 	Times(1).
	// 	Return(acc, nil)

	// // start test server and send request
	// server := NewServer(store)

	// t.Run("Status 200", func(t *testing.T) {
	// 	recorder200 := httptest.NewRecorder()

	// 	url := fmt.Sprintf("/account/%d", acc.ID)
	// 	request, err := http.NewRequest(http.MethodGet, url, nil)
	// 	assert.NoError(t, err)

	// 	server.router.ServeHTTP(recorder200, request)
	// 	assert.Equal(t, http.StatusOK, recorder200.Code)
	// 	requireBodyMatchAccount(t, recorder200.Body, acc)
	// })

	// t.Run("Status 400", func(t *testing.T) {
	// 	recorder400 := httptest.NewRecorder()

	// 	url := fmt.Sprintf("/account/-1")
	// 	request, err := http.NewRequest(http.MethodGet, url, nil)
	// 	assert.NoError(t, err)

	// 	server.router.ServeHTTP(recorder400, request)
	// 	assert.Equal(t, http.StatusBadRequest, recorder400.Code)

	// 	var emptyAcc db.Account
	// 	requireBodyMatchAccount(t, recorder400.Body, emptyAcc)
	// })

	// t.Run("Status 404", func(t *testing.T) {
	// 	recorder404 := httptest.NewRecorder()

	// 	url := fmt.Sprintf("/account/2")
	// 	request, err := http.NewRequest(http.MethodGet, url, nil)
	// 	assert.NoError(t, err)

	// 	fmt.Println("requst", request)
	// 	fmt.Println("idd", acc.ID)

	// 	server.router.ServeHTTP(recorder404, request)
	// 	assert.Equal(t, http.StatusNotFound, recorder404.Code)
	// })

	//assert.Equal(t, "http.StatusOK", recorder)
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
