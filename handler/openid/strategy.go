package openid

import (
	"context"

	"github.com/MatthewHartstonge/fosite"
)

type OpenIDConnectTokenStrategy interface {
	GenerateIDToken(ctx context.Context, requester fosite.Requester) (token string, err error)
}
