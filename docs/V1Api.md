# \V1Api

All URIs are relative to *https://validator.spaceapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Get**](V1Api.md#V1Get) | **Get** /v1 | 
[**V1ValidatePost**](V1Api.md#V1ValidatePost) | **Post** /v1/validate/ | validate an input against the SpaceApi schema



## V1Get

> ServerInformation V1Get(ctx, )



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


## V1ValidatePost

> ValidateV1Response V1ValidatePost(ctx, validateV1)

validate an input against the SpaceApi schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validateV1** | [**ValidateV1**](ValidateV1.md)|  | 

### Return type

[**ValidateV1Response**](ValidateV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

