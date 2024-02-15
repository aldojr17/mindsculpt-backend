package api

const (
	KEY_API_HEADER_X_KEY    = "X-Key"
	KEY_API_HEADER_X_SECRET = "X-Secret"
)

const (
	PATH_GET_MODELS     = "/key/api/v1/models"
	PATH_GET_GENERATION = "/key/api/v1/text2image/status/%s"
)

const (
	IMAGE_EXTENSION         = "jpg"
	FIREBASE_TOKEN_METADATA = "firebaseStorageDownloadTokens"
	FIREBASE_IMAGE_URL      = "https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s"
	STATUS_DONE             = "DONE"
	TIMEOUT                 = 3 // Seconds
)
