package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/edwinwalela/go-order/orders/config"
	"google.golang.org/api/option"
)

func Up(ctx context.Context, config config.Config) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("./firestore-sa.json")
	return firestore.NewClient(ctx, config.Db.ProjectId, opt)
}
