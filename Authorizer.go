package main

import (
	"AwsApiAuthorizer/contracts"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	authorizationHeader := request.AuthorizationToken
	responseRequest := events.APIGatewayCustomAuthorizerResponse{}
	if authorizationHeader != "" {

		_, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			return responseRequest, err
		}
		req, _ := http.NewRequest("GET", "", nil)
		req.Header.Set("Authorization", authorizationHeader)
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{

			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

				valid := token.Claims.(jwt.MapClaims).Valid()
				if valid != nil {
					return token, errors.New("Invalid token")
				}
				aud := ""
				checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAud {
					return token, errors.New("Invalid audience.")
				}
				iss := ""
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return token, errors.New("Invalid issuer.")
				}
				now := time.Now().Unix()
				expired := token.Claims.(jwt.MapClaims).VerifyExpiresAt(now, false)
				if expired {
					return token, errors.New("Token already expired")
				}
				return nil, nil
			},
			SigningMethod: jwt.SigningMethodRS256,
		})
		ww := &contracts.Response{}
		errResponse := error(jwtMiddleware.CheckJWT(ww, req))
		if errResponse != nil {
			return responseRequest, errResponse
		}
		responseRequest.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{}
	} else {
		return responseRequest, errors.New("Token is empty")
	}
	return responseRequest, nil
}

func main() {
	lambda.Start(handler)
}
