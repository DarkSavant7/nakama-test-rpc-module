package rpc

import (
	"context"
	"crypto/sha256"
	"dark.savant.org/nakama-test-rpc-module/database"
	"dark.savant.org/nakama-test-rpc-module/structures"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	SLASH               = "/"
	JSON_FILE_EXTENSION = ".json"
)

func TestRpcFunction(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("[TEST RPC MODULE]: got a request with the payload: %s", payload)
	data := structures.Request{}

	// Parse payload if it's not empty
	if payload != "" {
		err := json.Unmarshal([]byte(payload), &data)
		if err != nil {
			logger.Error("[TEST RPC MODULE]: Failed to parse payload: %v", err)
			return "", errInvalidPayload
		}
	}

	data.UpdateRequestPayloadWithDefaults()

	// Construct file path
	filePath := data.Type + SLASH + data.Version + JSON_FILE_EXTENSION

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errFileDoesNotExist
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error("[TEST RPC MODULE]: Failed to read file: %v", err)
		return "", errReadingFile
	}

	// Calculate file content hash
	hasher := sha256.New()
	hasher.Write(content)
	contentHash := hex.EncodeToString(hasher.Sum(nil))

	// Save data to the database
	insertErr := database.SaveToDB(ctx, db, logger, data)
	if insertErr != nil {
		logger.Error("[TEST RPC MODULE]: Failed to write request log: %v", err)
		return "", insertErr
	}

	// Prepare response
	var responseContent string
	if data.Hash != "" && data.Hash != contentHash {
		responseContent = ""
	} else {
		responseContent = base64.StdEncoding.EncodeToString(content)
	}
	response := structures.Response{Type: data.Type, Version: data.Version, Hash: contentHash, Content: responseContent}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logger.Error("[TEST RPC MODULE]: Failed to marshal response: %v", err)
		return "", errMarshalling
	}

	return string(responseJSON), nil
}
