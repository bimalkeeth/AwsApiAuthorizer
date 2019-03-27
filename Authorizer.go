package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strings"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	authorizationHeader := request.AuthorizationToken
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		println(bearerToken)

	} else {

	}

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			aud := "YOUR_API_IDENTIFIER"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://dev-8xmkmv6f.au.auth0.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			return nil, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
	println(jwtMiddleware)
	return events.APIGatewayCustomAuthorizerResponse{}, nil
}

func main() {
	lambda.Start(handler)
}
