package rpc

import "github.com/heroiclabs/nakama-common/runtime"

const (
	FILE_DOES_NOT_EXIST = 0
	ERROR_READING_FILE  = 1
	INVALID_PAYLOAD     = 2
	DATABASE_ERROR      = 3
	MARSHALLING_ERROR   = 4
)

var (
	errFileDoesNotExist = runtime.NewError("File not found", FILE_DOES_NOT_EXIST)
	errReadingFile      = runtime.NewError("File cannot be read", ERROR_READING_FILE)
	errInvalidPayload   = runtime.NewError("Payload is invalid", INVALID_PAYLOAD)
	errDatabase         = runtime.NewError("Database connection issue", DATABASE_ERROR)
	errMarshalling      = runtime.NewError("Error marshalling the response", MARSHALLING_ERROR)
)
