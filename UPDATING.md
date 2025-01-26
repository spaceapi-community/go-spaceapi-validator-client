# Updating the OpenAPI Client

```shell-session
$ docker run -v $PWD:/local openapitools/openapi-generator-cli:v4.2.2 \
  generate \
  -i https://validator.spaceapi.io/openapi.json \
  -g go \
  -o /local \
  --additional-properties=packageName=spaceapivalidatorclient \
  --git-user-id spaceapi-community \
  --git-repo-id go-spaceapi-validator-client
```
