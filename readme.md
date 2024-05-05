# Nakama RPC Module

This project is a Nakama server module written in Go. It includes an RPC function that accepts a payload with `type`, `version`, and `hash`. All parameters are optional with default values: `type=core`, `version=1.0.0`, `hash=null`.

## Functionality

The RPC function performs the following operations:

- Reads a file from the disk with the path format `\<type\>/\<version\>.json`.
- Saves data to the database.
- Calculates the hash of the file content.
- Returns a response with the fields: `type`, `version`, `hash`, `content`.
    - If the hashes are not equal, then the `content` field is an empty string.
    - If the file doesn't exist, then it returns an error.
- Uses default values if they are not present in the payload.

## Testing

The custom logic of the RPC function is covered with unit tests. These tests ensure that the function behaves as expected when reading files, interacting with the database, and handling different payload inputs.
Execute  ```go test``` to run tests


## Usage

To use this module, you need to have a Nakama server set up. Once the server is running, you can call the RPC function with a payload containing the `type`, `version`, and `hash` parameters. The function will read the corresponding file, save data to the database, calculate the file content hash, and return a response.

To run the Nakama server with the provided module you might use docker compose. In the project directory execute 
```
docker compose up
```
It will build module and run the Nakama integrating our module.


## Implementation

As the database was used Postgres as one of the available options without any particular reason. For thorough choosing the database more context is needed.
That might not look as typical Go code as it isn't my primary language, so that I probably update this project at some point.
For the storing data the table was required. Off the top of my head I made a DB init function which creates the table if required. I believe that there is some DB migration functionality in the Nakama server, so that I need to research a little bit
