package database

const CREATE_TABLE = `
		CREATE TABLE IF NOT EXISTS test_module_request_log
		(
			id         BIGSERIAL PRIMARY KEY,
			type       VARCHAR(50) NOT NULL,
			version    VARCHAR(50) NOT NULL,
			hash       VARCHAR(64) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
		  `

const INSERT_LOG = `
		INSERT INTO test_module_request_log (type, version, hash)
				VALUES ($1, $2, $3)
		  `
