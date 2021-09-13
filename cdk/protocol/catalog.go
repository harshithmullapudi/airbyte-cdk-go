package protocol

const (
	FULL_REFRESH_SYNC_MODE = "FULL_REFRESH"
	INCREMENTAL_SYNC_MODE  = "INCREMENTAL"
)

const (
	APPEND_DEST_SYNC_MODE    = "APPEND"
	OVERWRITE_DEST_SYNC_MODE = "OVERWRITE"
	APPEND_DEDUP_DEST_SYNC_MODE   = "APPEND_DEDUP"
)

type AirbyteStream struct {
	name                    string
	jsonSchema              map[string]interface{} `json:json_schema`
	supportedSyncModes      []string               `json:supported_sync_modes`
	sourceDefinedCursor     bool                   `json:source_defined_cursor`
	defaultCursorField      []string               `json:default_cursor_field`
	sourceDefinedPrimaryKey []string               `json:source_defined_primary_key`
	namespace               string
}

type AirbyteCatalog struct {
	streams []AirbyteStream
}

type 
