package api

const (
	KEY_API_HEADER_X_KEY    = "X-Key"
	KEY_API_HEADER_X_SECRET = "X-Secret"
)

const (
	PATH_GET_MODELS        = "/key/api/v1/models"
	PATH_GET_GENERATION    = "/key/api/v1/text2image/status/%s"
	PATH_CREATE_GENERATION = "/key/api/v1/text2image/run"
)

const (
	MODEL_STATUS_ACTIVE            = "ACTIVE"
	MODEL_STATUS_DISABLED_BY_QUEUE = "DISABLED_BY_QUEUE"
)

const (
	IMAGE_STATUS_INITIAL    = "INITIAL"
	IMAGE_STATUS_PROCESSING = "PROCESSING"
	IMAGE_STATUS_DONE       = "DONE"
	IMAGE_STATUS_FAIL       = "FAIL"
)

// 1:1 / 2:3 / 3:2 / 9:16 / 16:9
// You and your users can make requests in Russian, English or any other
// language. It is also allowed to use emoji in the text description.
// The maximum size of a text description is 1000 characters.

const (
	IMAGE_EXTENSION         = "jpg"
	FIREBASE_TOKEN_METADATA = "firebaseStorageDownloadTokens"
	FIREBASE_IMAGE_URL      = "https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s"
	TIMEOUT                 = 3 // Seconds
)

const GENERATE_PARAMS = `{
	"type": "GENERATE",
	"width": %d,
	"height": %d,
	"num_images": 1,
	"negativePromptUnclip": "%s",
	"generateParams": {
		"query": "%s"
	}
}`
