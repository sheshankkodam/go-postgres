package db_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sheshankkodam/go-postgres/db"
	"github.com/sheshankkodam/go-postgres/db/mocker"
	"github.com/stretchr/testify/assert"
	"testing"
)

var insertTests = []struct {
	name, userName, deptName, date string
	expectedError                  error
}{
	{"happy path", "sheshank", "sse", "05-08-2018", nil},
	{"invalid date", "sheshank", "sse", "05-2018", errors.New("invalid date")},
}

func TestPostgresConnection_InsertHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostgresOps := mocker.NewMockPostgresOperations(ctrl)
	p := db.PostgresService{Ops: mockPostgresOps}

	for _, tc := range insertTests {
		mockPostgresOps.EXPECT().Insert(tc.userName, tc.deptName, tc.date).Return(tc.expectedError).Times(1)
		err := p.Ops.Insert(tc.userName, tc.deptName, tc.date)
		assert.Equal(t, tc.expectedError, err)
	}
}
