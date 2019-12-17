# \V2Api

All URIs are relative to *https://validator.spaceapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2Get**](V2Api.md#V2Get) | **Get** /v2 | 
[**V2ValidateJSONPost**](V2Api.md#V2ValidateJSONPost) | **Post** /v2/validateJSON | validate an input against the SpaceApi schema
[**V2ValidateURLPost**](V2Api.md#V2ValidateURLPost) | **Post** /v2/validateURL | validate an input against the SpaceApi schema



## V2Get

> ServerInformation V2Get(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServerInformation**](ServerInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ValidateJSONPost

> ValidateJsonV2Response V2ValidateJSONPost(ctx, body)

validate an input against the SpaceApi schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**ValidateJsonV2Response**](ValidateJsonV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ValidateURLPost

> ValidateUrlV2Response V2ValidateURLPost(ctx, validateUrlV2)

validate an input against the SpaceApi schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validateUrlV2** | [**ValidateUrlV2**](ValidateUrlV2.md)|  | 

### Return type

[**ValidateUrlV2Response**](ValidateUrlV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

