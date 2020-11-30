/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package oauth2

import (
	"context"
	"time"

	"github.com/ory/fosite"
	"gopkg.in/square/go-jose.v2"
)

type CoreStorage interface {
	AuthorizeCodeStorage
	AccessTokenStorage
	RefreshTokenStorage
}

// AuthorizeCodeStorage handles storage requests related to authorization codes.
type AuthorizeCodeStorage interface {
	// GetAuthorizeCodeSession stores the authorization request for a given authorization code.
	CreateAuthorizeCodeSession(ctx context.Context, code string, request fosite.Requester) (err error)

	// GetAuthorizeCodeSession hydrates the session based on the given code and returns the authorization request.
	// If the authorization code has been invalidated with `InvalidateAuthorizeCodeSession`, this
	// method should return the ErrInvalidatedAuthorizeCode error.
	//
	// Make sure to also return the fosite.Requester value when returning the fosite.ErrInvalidatedAuthorizeCode error!
	GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (request fosite.Requester, err error)

	// InvalidateAuthorizeCodeSession is called when an authorize code is being used. The state of the authorization
	// code should be set to invalid and consecutive requests to GetAuthorizeCodeSession should return the
	// ErrInvalidatedAuthorizeCode error.
	InvalidateAuthorizeCodeSession(ctx context.Context, code string) (err error)
}

type AccessTokenStorage interface {
	CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error)

	GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error)

	DeleteAccessTokenSession(ctx context.Context, signature string) (err error)
}

type RefreshTokenStorage interface {
	CreateRefreshTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error)

	GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error)

	DeleteRefreshTokenSession(ctx context.Context, signature string) (err error)
}

// JWTAuthGrantStorage holds information needed to validate jwt assertion in authorization grants.
type JWTAuthGrantStorage interface {
	// GetPublicKey returns public key, issued by 'issuer', and assigned for subject. Public key is used to check
	// signature of jwt assertion in authorization grants.
	GetPublicKey(ctx context.Context, issuer string, subject string, keyId string) (*jose.JSONWebKey, error)

	// GetPublicKeys returns public key, set issued by 'issuer', and assigned for subject.
	GetPublicKeys(ctx context.Context, issuer string, subject string) (*jose.JSONWebKeySet, error)

	// GetPublicKeyScopes returns assigned scope for assertion, identified by public key, issued by 'issuer'.
	GetPublicKeyScopes(ctx context.Context, issuer string, subject string, keyId string) ([]string, error)

	// IsJWTUsed returns true, if JWT is not known yet or it can not be considered valid, because it must be already
	// expired.
	IsJWTUsed(ctx context.Context, jti string) (bool, error)

	// MarkJWTUsedForTime marks JWT as used for a time passed in exp parameter. This helps ensure that JWTs are not
	// replayed by maintaining the set of used "jti" values for the length of time for which the JWT would be
	// considered valid based on the applicable "exp" instant. (https://tools.ietf.org/html/rfc7523#section-3)
	MarkJWTUsedForTime(ctx context.Context, jti string, exp time.Time) error
}
