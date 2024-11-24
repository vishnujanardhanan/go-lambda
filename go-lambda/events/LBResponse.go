package events

type LBResponse struct {
	StatusCode        int               `json:"statusCode"`        // HTTP status code (e.g., 200, 404)
	StatusDescription string            `json:"statusDescription"` // Status description (optional)
	Headers           map[string]string `json:"headers"`           // Response headers
	Body              string            `json:"body"`              // Response body
	IsBase64Encoded   bool              `json:"isBase64Encoded"`   // Whether the body is base64-encoded
}
