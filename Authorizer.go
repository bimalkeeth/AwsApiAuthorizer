package main

import (
	"fmt"
	"github.com/auth0-community/go-auth0"
	"github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"
	"os"
	"strings"
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

	tok := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik5UZEdOVE14T1RoRE56bEdSVFkzUlVNME16SkRNRFZHUWprMFJqY3pPVEZHTTBZelJUWkNNZyJ9.eyJpc3MiOiJodHRwczovL2Fzc2V0aWMtdGVzdC5hdS5hdXRoMC5jb20vIiwic3ViIjoiSmNRcUhCUUdNa01nY0NEdEtsRHZYQ0huM0pCd1hJMUxAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXNzZXRpYy10ZXN0LmF1LmF1dGgwLmNvbS9hcGkvdjIvIiwiaWF0IjoxNTUzNjg2MzgwLCJleHAiOjE1NTM3NzI3ODAsImF6cCI6IkpjUXFIQlFHTWtNZ2NDRHRLbER2WENIbjNKQndYSTFMIiwic2NvcGUiOiJyZWFkOmNsaWVudF9ncmFudHMgY3JlYXRlOmNsaWVudF9ncmFudHMgZGVsZXRlOmNsaWVudF9ncmFudHMgdXBkYXRlOmNsaWVudF9ncmFudHMgcmVhZDp1c2VycyB1cGRhdGU6dXNlcnMgZGVsZXRlOnVzZXJzIGNyZWF0ZTp1c2VycyByZWFkOnVzZXJzX2FwcF9tZXRhZGF0YSB1cGRhdGU6dXNlcnNfYXBwX21ldGFkYXRhIGRlbGV0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgY3JlYXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBjcmVhdGU6dXNlcl90aWNrZXRzIHJlYWQ6Y2xpZW50cyB1cGRhdGU6Y2xpZW50cyBkZWxldGU6Y2xpZW50cyBjcmVhdGU6Y2xpZW50cyByZWFkOmNsaWVudF9rZXlzIHVwZGF0ZTpjbGllbnRfa2V5cyBkZWxldGU6Y2xpZW50X2tleXMgY3JlYXRlOmNsaWVudF9rZXlzIHJlYWQ6Y29ubmVjdGlvbnMgdXBkYXRlOmNvbm5lY3Rpb25zIGRlbGV0ZTpjb25uZWN0aW9ucyBjcmVhdGU6Y29ubmVjdGlvbnMgcmVhZDpyZXNvdXJjZV9zZXJ2ZXJzIHVwZGF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGRlbGV0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGNyZWF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIHJlYWQ6ZGV2aWNlX2NyZWRlbnRpYWxzIHVwZGF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgZGVsZXRlOmRldmljZV9jcmVkZW50aWFscyBjcmVhdGU6ZGV2aWNlX2NyZWRlbnRpYWxzIHJlYWQ6cnVsZXMgdXBkYXRlOnJ1bGVzIGRlbGV0ZTpydWxlcyBjcmVhdGU6cnVsZXMgcmVhZDpydWxlc19jb25maWdzIHVwZGF0ZTpydWxlc19jb25maWdzIGRlbGV0ZTpydWxlc19jb25maWdzIHJlYWQ6ZW1haWxfcHJvdmlkZXIgdXBkYXRlOmVtYWlsX3Byb3ZpZGVyIGRlbGV0ZTplbWFpbF9wcm92aWRlciBjcmVhdGU6ZW1haWxfcHJvdmlkZXIgYmxhY2tsaXN0OnRva2VucyByZWFkOnN0YXRzIHJlYWQ6dGVuYW50X3NldHRpbmdzIHVwZGF0ZTp0ZW5hbnRfc2V0dGluZ3MgcmVhZDpsb2dzIHJlYWQ6c2hpZWxkcyBjcmVhdGU6c2hpZWxkcyBkZWxldGU6c2hpZWxkcyByZWFkOmFub21hbHlfYmxvY2tzIGRlbGV0ZTphbm9tYWx5X2Jsb2NrcyB1cGRhdGU6dHJpZ2dlcnMgcmVhZDp0cmlnZ2VycyByZWFkOmdyYW50cyBkZWxldGU6Z3JhbnRzIHJlYWQ6Z3VhcmRpYW5fZmFjdG9ycyB1cGRhdGU6Z3VhcmRpYW5fZmFjdG9ycyByZWFkOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGRlbGV0ZTpndWFyZGlhbl9lbnJvbGxtZW50cyBjcmVhdGU6Z3VhcmRpYW5fZW5yb2xsbWVudF90aWNrZXRzIHJlYWQ6dXNlcl9pZHBfdG9rZW5zIGNyZWF0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIGRlbGV0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIHJlYWQ6Y3VzdG9tX2RvbWFpbnMgZGVsZXRlOmN1c3RvbV9kb21haW5zIGNyZWF0ZTpjdXN0b21fZG9tYWlucyByZWFkOmVtYWlsX3RlbXBsYXRlcyBjcmVhdGU6ZW1haWxfdGVtcGxhdGVzIHVwZGF0ZTplbWFpbF90ZW1wbGF0ZXMgcmVhZDptZmFfcG9saWNpZXMgdXBkYXRlOm1mYV9wb2xpY2llcyByZWFkOnJvbGVzIGNyZWF0ZTpyb2xlcyBkZWxldGU6cm9sZXMgdXBkYXRlOnJvbGVzIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.UTvG_Pne9MIPtp6hx_o8nz_TyV0DMNb36dYXiz2DQ6uHE7VnvPYX7hNTfN_PZifQrz3P2VJG7QWW-gRE-goivymaUpBtLcrU1KLt7Di5XU2XbFikPRqsHVYnutk_KXLBLW8ZaMWJhduUM8T-d4WL9okUC-Udjj6beWhK3IjU4B7cnEyZvnomgA8N7dqFpAGhay3jpwTnf253NAPsLVYkMXdtc1Eoig-9-RtsXP7Ff3Anagm5L_xJkP3jDIfWQ-Vg43YIJHjB9qey5Fta4OOs_7YD_qouPOyHWj3BYMYVYEX5v488Ifowf4OUabsqt60H9S6eeW6VLmzSeo0zaC74cg"

	token, _ := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	println(token)

	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: "https://mydomain.eu.auth0.com/.well-known/jwks.json"}, nil)
	audience := os.Getenv("AUTH0_CLIENT_ID")
	configuration := auth0.NewConfiguration(client, []string{audience}, "https://mydomain.eu.auth0.com/", jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	token2, err := validator.ValidateRequest(token)

	lambda.Start(handler)
}
