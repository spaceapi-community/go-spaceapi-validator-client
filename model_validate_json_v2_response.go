/*
 * SpaceApi Validator
 *
 * This is the SpaceApi Validator api
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spaceapivalidatorclient
// ValidateJsonV2Response struct for ValidateJsonV2Response
type ValidateJsonV2Response struct {
	Valid bool `json:"valid"`
	Message string `json:"message"`
	CheckedVersions []string `json:"checkedVersions,omitempty"`
	ValidatedJson map[string]interface{} `json:"validatedJson,omitempty"`
	SchemaErrors []SchemaError `json:"schemaErrors,omitempty"`
}
