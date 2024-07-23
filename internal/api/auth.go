package api

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pietervdwerk/tasksapi/pkg/openapi3"
)

// Token endpoint
// (POST /token)
func (api *API) PostToken(ctx context.Context, req PostTokenRequestObject) (PostTokenResponseObject, error) {
	clientID, clientSecret, err := validateAuthorizationHeader(req.Params.Authorization)
	if err != nil {
		return PostToken400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	// Validate parameters
	if req.Body.GrantType != "authorization_code" {
		return PostToken400JSONResponse{
			Message: "invalid grant type",
		}, nil
	}

	clientIDGuid, err := uuid.Parse(clientID)
	if err != nil {
		return PostToken400JSONResponse{
			Message: "invalid client id",
		}, nil
	}

	client, err := api.clientsRepo.GetByID(context.Background(), clientIDGuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return PostToken401JSONResponse{
				Message: "clientid not found",
			}, nil
		}
		return nil, err
	}

	if client.Secret != clientSecret {
		return PostToken401JSONResponse{
			Message: "invalid client secret",
		}, nil
	}

	// Generate JWT token
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": clientID,
		"iat": time.Now().Unix(),
		"exp": exp,
	})

	// Sign the token
	tokenString, err := token.SignedString(api.conf.privateKey)
	if err != nil {
		return nil, err
	}

	return PostToken200JSONResponse{
		AccessToken: tokenString,
		ExpiresIn:   int(exp),
		TokenType:   openapi3.Bearer,
		Scope:       req.Body.Scope,
	}, nil
}

// Validate the authorization header
func validateAuthorizationHeader(header string) (clientId string, clientSecret string, err error) {
	if header == "" {
		return "", "", errors.New("missing Authorization header")
	}

	// Extract the client ID and secret from the authorization header
	if !strings.HasPrefix(header, "Basic ") {
		return "", "", errors.New("invalid Authorization header format")
	}

	// Decode the client ID and secret
	encodedCreds := strings.TrimPrefix(header, "Basic ")
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return "", "", errors.New("authorization header invalid base64")
	}

	// Extract the client ID and secret
	creds := strings.Split(string(decodedCreds), ":")
	if len(creds) != 2 {
		return "", "", errors.New("invalid Authorization header format")
	}
	return creds[0], creds[1], nil
}
