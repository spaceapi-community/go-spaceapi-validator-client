/*
 * SpaceApi Validator
 *
 * This is the SpaceApi Validator api
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spaceapivalidatorclient

// ValidateUrlV2Response struct for ValidateUrlV2Response
type ValidateUrlV2Response struct {
	Valid         bool                   `json:"valid"`
	Message       string                 `json:"message,omitempty"`
	IsHttps       bool                   `json:"isHttps"`
	HttpsForward  bool                   `json:"httpsForward"`
	Reachable     bool                   `json:"reachable"`
	Cors          bool                   `json:"cors"`
	ContentType   bool                   `json:"contentType"`
	CertValid     bool                   `json:"certValid"`
	ValidatedJson map[string]interface{} `json:"validatedJson,omitempty"`
	SchemaErrors  []SchemaError          `json:"schemaErrors,omitempty"`
}
