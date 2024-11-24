package events

type LBRequest struct {
	RequestContext struct {
		Elb struct {
			TargetGroupArn string `json:"targetGroupArn"`
		} `json:"elb"`
	} `json:"requestContext"`
	HttpMethod            string            `json:"httpMethod"`            // HTTP method (e.g., GET, POST)
	Path                  string            `json:"path"`                  // Path of the request (e.g., "/greet/{name}")
	Headers               map[string]string `json:"headers"`               // Request headers
	QueryStringParameters map[string]string `json:"queryStringParameters"` // Query parameters
	PathParameters        map[string]string `json:"pathParameters"`        // Path parameters (e.g., "/greet/{name}")
	Body                  string            `json:"body"`                  // Request body (if any)
	IsBase64Encoded       bool              `json:"isBase64Encoded"`       // Whether the body is base64-encoded
}
