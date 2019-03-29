package main

import (
	"AwsApiAuthorizer/contracts"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	authorizationHeader := request.AuthorizationToken
	authorizationHeader = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik5UZEdOVE14T1RoRE56bEdSVFkzUlVNME16SkRNRFZHUWprMFJqY3pPVEZHTTBZelJUWkNNZyJ9.eyJpc3MiOiJodHRwczovL2Fzc2V0aWMtdGVzdC5hdS5hdXRoMC5jb20vIiwic3ViIjoiUDZzd3huc1B2NlRybWE2MDlzTENrOVdYYlI2bVlGNUlAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXBpLnByZWRpY3Rvci5jb20vYXNkIiwiaWF0IjoxNTUzODI4MDE1LCJleHAiOjE1NTM5MTQ0MTUsImF6cCI6IlA2c3d4bnNQdjZUcm1hNjA5c0xDazlXWGJSNm1ZRjVJIiwic2NvcGUiOiJyZWFkOmFwcG9pbnRtZW50cyByZWFkOmZpbGVzIHdyaXRlOmZpbGVzIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.BAmphNLkkf-NeCJYw_-TaSHIOVuMo9H4VLEz0g1KaxXKqzia8jwpbU2VmrBEQWVVFq-5ToXUVZbyIw_31RMX0Ob8q5npjj4ZABFuxABoG5hEdxTYy9rV6LuOEfyeZ0XyreEvoRRjiY0S4jFabLVmgqA9NEjXFzFFCT_kGayxbqOWOPf0tIS5m2EqXZFmdDIKqx6bajGZmeonWwuMaa7QNed8EfCTI6NHAcLjpQCFA0z9MeU6tfjnJrnHw3oDdQaogJU5iDdrF4H6wgnEL_G3UcHrJwnl8WAJopNQ3-eDLJZ2a3udX4cGO1Ke9Hw2t5FQ_JfEm20QGwCD9JPqEAKy6g"
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
		if errResponse != nil && errResponse.Error() != "" {
			return responseRequest, errResponse
		}
		responseRequest.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{}
	} else {
		return responseRequest, errors.New("Token is empty")
	}
	return responseRequest, nil
}

func main() {

	authorizationHeader := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik5UZEdOVE14T1RoRE56bEdSVFkzUlVNME16SkRNRFZHUWprMFJqY3pPVEZHTTBZelJUWkNNZyJ9.eyJpc3MiOiJodHRwczovL2Fzc2V0aWMtdGVzdC5hdS5hdXRoMC5jb20vIiwic3ViIjoiUDZzd3huc1B2NlRybWE2MDlzTENrOVdYYlI2bVlGNUlAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXBpLnByZWRpY3Rvci5jb20vYXNkIiwiaWF0IjoxNTUzODI4Njg0LCJleHAiOjE1NTM5MTUwODQsImF6cCI6IlA2c3d4bnNQdjZUcm1hNjA5c0xDazlXWGJSNm1ZRjVJIiwic2NvcGUiOiJyZWFkOmFwcG9pbnRtZW50cyByZWFkOmZpbGVzIHdyaXRlOmZpbGVzIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.c-tgTWqSygRJuO5vvRNtzHy7lF2MDju7dXi7ALF-rgqm7-rQCoCMWM9CQBPIemPzcOfGy2hWvHACLh2x29VcJn9u1Ln5s-95aaqQVr6MTLvOCiY3s_rxnOwsfFtiZ6djvpzI7t2SgciUstMCtbXAq1n14TYlrmph4jE9QHFLpzclRXsBVHoqFOR4vqqgiDcP2Cg2IYPOUWN5lBG4BaDhGGgD86sd0CSIU_XgQgw69MsAuimPr1Af6yElyip-PafFiYnbWgszmPvWW8ycTVIoHb1Hqmg-ijRQgTau_-mE4bdUT38g28rW_jaoxiFtcIeujAigfwK-vop7U0I_aSDyeQ"
	responseRequest := events.APIGatewayCustomAuthorizerResponse{}
	if authorizationHeader != "" {

		_, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if err != nil {

		}
		req, _ := http.NewRequest("GET", "", nil)
		req.Header.Set("Authorization", authorizationHeader)
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{

			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

				valid := token.Claims.(jwt.MapClaims).Valid()
				if valid != nil {
					return token, errors.New("Invalid token")
				}
				aud := "https://api.predictor.com/asd"
				checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAud {
					return token, errors.New("Invalid audience.")
				}
				iss := "https://assetic-test.au.auth0.com/"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return token, errors.New("Invalid issuer.")
				}
				now := time.Time{}.Unix()
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
		if errResponse != nil && errResponse.Error() != "" {
			fmt.Println(errResponse.Error())
		}
		responseRequest.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{}
	} else {

	}

	//lambda.Start(handler)
}
