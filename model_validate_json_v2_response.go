/*
 * SpaceApi Validator
 *
 * This is the SpaceApi Validator api
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spaceapivalidatorclient

// ValidateJsonV2Response struct for ValidateJsonV2Response
type ValidateJsonV2Response struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}