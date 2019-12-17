# Go API client for spaceapivalidatorclient

This is the SpaceApi Validator api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./spaceapivalidatorclient"
```

## Documentation for API Endpoints

All URIs are relative to *https://validator.spaceapi.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1Api* | [**V1Get**](docs/V1Api.md#v1get) | **Get** /v1 | 
*V1Api* | [**V1ValidatePost**](docs/V1Api.md#v1validatepost) | **Post** /v1/validate/ | validate an input against the SpaceApi schema
*V2Api* | [**V2Get**](docs/V2Api.md#v2get) | **Get** /v2 | 
*V2Api* | [**V2ValidateJSONPost**](docs/V2Api.md#v2validatejsonpost) | **Post** /v2/validateJSON | validate an input against the SpaceApi schema
*V2Api* | [**V2ValidateURLPost**](docs/V2Api.md#v2validateurlpost) | **Post** /v2/validateURL | validate an input against the SpaceApi schema


## Documentation For Models

 - [ServerInformation](docs/ServerInformation.md)
 - [ValidateJsonV2Response](docs/ValidateJsonV2Response.md)
 - [ValidateUrlV2](docs/ValidateUrlV2.md)
 - [ValidateUrlV2Response](docs/ValidateUrlV2Response.md)
 - [ValidateV1](docs/ValidateV1.md)
 - [ValidateV1Response](docs/ValidateV1Response.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author



