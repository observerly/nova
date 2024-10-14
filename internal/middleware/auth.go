/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package middleware

/*****************************************************************************************************************/

import (
	"context"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"firebase.google.com/go/v4/auth"
)

/*****************************************************************************************************************/

func IsAuthenticated(ctx context.Context, req connect.AnyRequest, auth *auth.Client) (bool, error) {
	// Get the token from the request header:
	authorization := req.Header().Get("Authorization")

	// Remove the "Bearer " prefix from the "Authorization" header:
	idToken := strings.Replace(authorization, "Bearer ", "", 1)

	// Verify the token:
	token, err := auth.VerifyIDTokenAndCheckRevoked(ctx, idToken)

	if err != nil {
		return false, fmt.Errorf("error verifying token: %w", err)
	}

	// Check if the token is valid:
	if token == nil {
		return false, fmt.Errorf("token is invalid")
	}

	// Check if the token expiry time is not in the past:
	if token.Expires <= time.Now().Unix() {
		return false, fmt.Errorf("token has expired")
	}

	return true, nil
}

/*****************************************************************************************************************/
