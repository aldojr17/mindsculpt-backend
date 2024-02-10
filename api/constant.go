package api

const (
	KEY_API_HEADER_X_KEY    = "X-Key"
	KEY_API_HEADER_X_SECRET = "X-Secret"
)

const (
	PATH_GET_MODELS     = "/key/api/v1/models"
	PATH_GENERATE       = "/key/api/v1/text2image/run"
	PATH_GET_GENERATION = "/key/api/v1/text2image/status/%s"
)
