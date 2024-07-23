# Tasks API

## Running the project

This project uses a in memory sqlite database, so you don't need to install anything else to run it.
Just do

```bash
go run cmd/server/main.go
```

## Assumptions

1. The goal an overview of code quality and design skills, not about all edge cases and nice DX for the user.
2. A JWT must be created so using the Authentication Code Flow would make little sense unless you want to create your own auth server, in which the time required to do both would exceed reasonable expectations. So any flow resulting in signing a JWT is acceptable.
3. The packages used are completely up to the implementer.

## Notes

1. I started creating a simple auth server with the authorization code flow for testing, my experiment can be found in the cmd/auth folder.

2. I've never used oapi-codegen before, and although it has very nice features there are certain thinks like how JSON request validation errors are thrown when the request has no json body that I don't like. Could probably be handled better, but out of the scope of this project.
