package oauth2

import (
	"context"

	"github.com/MatthewHartstonge/fosite"
)

type ImplicitGrantStorage interface {
	CreateImplicitAccessTokenSession(ctx context.Context, token string, request fosite.Requester) (err error)
}
