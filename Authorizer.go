package main

import (
	"errors"
	"fmt"
	"github.com/auth0-community/go-auth0"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/square/go-jose.v2"
	"net/http"
	"strings"
	"time"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authorizationHeader := request.AuthorizationToken
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		println(bearerToken)

	} else {

	}
	return events.APIGatewayCustomAuthorizerResponse{}, nil
}

func main() {

	tok := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik5UZEdOVE14T1RoRE56bEdSVFkzUlVNME16SkRNRFZHUWprMFJqY3pPVEZHTTBZelJUWkNNZyJ9.eyJpc3MiOiJodHRwczovL2Fzc2V0aWMtdGVzdC5hdS5hdXRoMC5jb20vIiwic3ViIjoiSmNRcUhCUUdNa01nY0NEdEtsRHZYQ0huM0pCd1hJMUxAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXNzZXRpYy10ZXN0LmF1LmF1dGgwLmNvbS9hcGkvdjIvIiwiaWF0IjoxNTUzNzM0OTE3LCJleHAiOjE1NTM4MjEzMTcsImF6cCI6IkpjUXFIQlFHTWtNZ2NDRHRLbER2WENIbjNKQndYSTFMIiwic2NvcGUiOiJyZWFkOmNsaWVudF9ncmFudHMgY3JlYXRlOmNsaWVudF9ncmFudHMgZGVsZXRlOmNsaWVudF9ncmFudHMgdXBkYXRlOmNsaWVudF9ncmFudHMgcmVhZDp1c2VycyB1cGRhdGU6dXNlcnMgZGVsZXRlOnVzZXJzIGNyZWF0ZTp1c2VycyByZWFkOnVzZXJzX2FwcF9tZXRhZGF0YSB1cGRhdGU6dXNlcnNfYXBwX21ldGFkYXRhIGRlbGV0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgY3JlYXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBjcmVhdGU6dXNlcl90aWNrZXRzIHJlYWQ6Y2xpZW50cyB1cGRhdGU6Y2xpZW50cyBkZWxldGU6Y2xpZW50cyBjcmVhdGU6Y2xpZW50cyByZWFkOmNsaWVudF9rZXlzIHVwZGF0ZTpjbGllbnRfa2V5cyBkZWxldGU6Y2xpZW50X2tleXMgY3JlYXRlOmNsaWVudF9rZXlzIHJlYWQ6Y29ubmVjdGlvbnMgdXBkYXRlOmNvbm5lY3Rpb25zIGRlbGV0ZTpjb25uZWN0aW9ucyBjcmVhdGU6Y29ubmVjdGlvbnMgcmVhZDpyZXNvdXJjZV9zZXJ2ZXJzIHVwZGF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGRlbGV0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGNyZWF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIHJlYWQ6ZGV2aWNlX2NyZWRlbnRpYWxzIHVwZGF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgZGVsZXRlOmRldmljZV9jcmVkZW50aWFscyBjcmVhdGU6ZGV2aWNlX2NyZWRlbnRpYWxzIHJlYWQ6cnVsZXMgdXBkYXRlOnJ1bGVzIGRlbGV0ZTpydWxlcyBjcmVhdGU6cnVsZXMgcmVhZDpydWxlc19jb25maWdzIHVwZGF0ZTpydWxlc19jb25maWdzIGRlbGV0ZTpydWxlc19jb25maWdzIHJlYWQ6ZW1haWxfcHJvdmlkZXIgdXBkYXRlOmVtYWlsX3Byb3ZpZGVyIGRlbGV0ZTplbWFpbF9wcm92aWRlciBjcmVhdGU6ZW1haWxfcHJvdmlkZXIgYmxhY2tsaXN0OnRva2VucyByZWFkOnN0YXRzIHJlYWQ6dGVuYW50X3NldHRpbmdzIHVwZGF0ZTp0ZW5hbnRfc2V0dGluZ3MgcmVhZDpsb2dzIHJlYWQ6c2hpZWxkcyBjcmVhdGU6c2hpZWxkcyBkZWxldGU6c2hpZWxkcyByZWFkOmFub21hbHlfYmxvY2tzIGRlbGV0ZTphbm9tYWx5X2Jsb2NrcyB1cGRhdGU6dHJpZ2dlcnMgcmVhZDp0cmlnZ2VycyByZWFkOmdyYW50cyBkZWxldGU6Z3JhbnRzIHJlYWQ6Z3VhcmRpYW5fZmFjdG9ycyB1cGRhdGU6Z3VhcmRpYW5fZmFjdG9ycyByZWFkOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGRlbGV0ZTpndWFyZGlhbl9lbnJvbGxtZW50cyBjcmVhdGU6Z3VhcmRpYW5fZW5yb2xsbWVudF90aWNrZXRzIHJlYWQ6dXNlcl9pZHBfdG9rZW5zIGNyZWF0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIGRlbGV0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIHJlYWQ6Y3VzdG9tX2RvbWFpbnMgZGVsZXRlOmN1c3RvbV9kb21haW5zIGNyZWF0ZTpjdXN0b21fZG9tYWlucyByZWFkOmVtYWlsX3RlbXBsYXRlcyBjcmVhdGU6ZW1haWxfdGVtcGxhdGVzIHVwZGF0ZTplbWFpbF90ZW1wbGF0ZXMgcmVhZDptZmFfcG9saWNpZXMgdXBkYXRlOm1mYV9wb2xpY2llcyByZWFkOnJvbGVzIGNyZWF0ZTpyb2xlcyBkZWxldGU6cm9sZXMgdXBkYXRlOnJvbGVzIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.w_dPw7lXfoBvdCWD5vfkjFL7jf5a2F-uT7BSkZoV2GIGC3pHs6s6JSlrSqFdGdvISwodVCTgNdmmXnjPtPRyaUPZwEOQMPdDwPa57QuSGrXFDPJGw7vhzPft90Gp_uIQa20OORa0yamF3z4d3vIBrALh3336CkeOKoHI8Y-Cw1whl4UvnL6NgvRvSqjpXQIAFFSeRgu6jdtWVdsZV78L846kQFkMy3vx7DxMFaQSzMzf639Sp6kDz02upAOSCfctUo8J-qpeF8jIWK3KpANQzCMQkNuJ7RH-BeYyc0JLXW2jnKes_xYPxoBWeza3nXVLXR84E8pBIMR2rWbgOkIMkw"

	token, _ := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	println(token)

	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: "https://assetic-test.au.auth0.com/.well-known/jwks.json"}, nil)
	audience := "https://assetic-test.au.auth0.com/api/v2/"
	configuration := auth0.NewConfiguration(client, []string{audience}, "https://assetic-test.au.auth0.com/", jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("Authorization", tok)

	token2, err := validator.ValidateRequest(req)
	if err != nil {
		fmt.Println("Token is not valid")
	}

	claims := map[string]interface{}{}
	err = validator.Claims(req, token2, &claims)
	if err != nil {
		fmt.Println("Token claims are not valid")
	}

	tokena, erra := validator.ValidateRequestWithLeeway(req, time.Duration(time.Second*60))
	if erra != nil {
		fmt.Println("Token claims are not valid")
	}

	println(tokena)

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
	println(jwtMiddleware.CheckJWT(nil, req))
	//jwtMiddleware.CheckJWT(http.ResponseWriter,req)

	lambda.Start(handler)
}
