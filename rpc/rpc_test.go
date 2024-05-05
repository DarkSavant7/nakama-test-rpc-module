package rpc

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"os"
	"testing"
)

type MockLogger struct {
	runtime.Logger
}

type MockNakamaModule struct {
	runtime.NakamaModule
}

func (m *MockLogger) Debug(format string, v ...interface{}) {
	//do nothing
}

func (m *MockLogger) Error(format string, v ...interface{}) {
	//do nothing
}

func TestTestRpcFunction(t *testing.T) {
	testCases := []struct {
		name           string
		payload        string
		db             *sql.DB
		mockFileExists bool
		mockFileData   string
		expectedError  string
	}{
		{
			name:           "ValidPayload",
			payload:        `{"hash": "","type": "test","version": "v1"}`,
			mockFileExists: true,
			mockFileData:   "test data",
		},
		{
			name:           "InvalidPayload",
			payload:        `\{`,
			mockFileExists: true,
			mockFileData:   "test data",
			expectedError:  "Payload is invalid",
		},
		{
			name:           "FileNotExist",
			payload:        `{"hash": "","type": "test","version": "v1"}`,
			mockFileExists: false,
			mockFileData:   "",
			expectedError:  "File not found",
		},
	}

	mkdirErr := os.Mkdir("test", 0666)
	if mkdirErr != nil {
		t.Fatalf("Failed to create mock file folder: %v", mkdirErr)
	}
	defer os.Remove("test")
	logger := new(MockLogger)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, dbErr := sqlmock.New()
			if dbErr != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", dbErr)
			}
			defer db.Close()
			mock.ExpectExec("INSERT INTO test_module_request_log").WillReturnResult(sqlmock.NewResult(1, 1))
			if tt.mockFileExists {
				err := os.WriteFile("test/v1.json", []byte(tt.mockFileData), 0644)
				if err != nil {
					t.Fatalf("Failed to create mock file: %v", err)
				}
				defer os.Remove("test/v1.json")
			}

			_, err := TestRpcFunction(
				context.Background(),
				logger,
				db,
				&MockNakamaModule{},
				tt.payload,
			)

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("Expected error %v, but got error: %v", tt.expectedError, err)
			} else if err == nil && tt.expectedError != "" {
				t.Errorf("Expected error %v, but got nil", tt.expectedError)
			}
		})
	}
}
