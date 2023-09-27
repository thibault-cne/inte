package mongo

import (
	"backend/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User's CRUD operations
func (b *Backend) CreateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := b.TimeoutContext(ctx)
	defer cancel()

	user.CreatedAt = time.Now().Unix()

	_, err := b.db.Collection(UsersCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}


func (b *Backend) GetUser(ctx context.Context, id string) (*models.User, error) {
	ctx, cancel := b.TimeoutContext(ctx)
	defer cancel()

	var user models.User
	err := b.db.Collection(UsersCollection).FindOne(ctx, bson.M{
		"google_id": id,

		"$or": []bson.M{
			{
				"deleted_at": bson.M{
					"$exists": false,
				},
			},
			{
				"deleted_at": nil,
			},
		},
	}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (b *Backend) UpdateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := b.TimeoutContext(ctx)
	defer cancel()

	res := b.db.Collection(UsersCollection).FindOneAndUpdate(ctx,
		bson.M{
			"id": user.User.GoogleId,

			"$or": []bson.M{
				{
					"deleted_at": bson.M{
						"$exists": false,
					},
				},
				{
					"deleted_at": nil,
				},
			},
		},
		bson.M{
			"$set": user,
		},
		options.FindOneAndUpdate().SetUpsert(true))
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (b *Backend) DeleteUser(ctx context.Context, id string) error {
	ctx, cancel := b.TimeoutContext(ctx)
	defer cancel()

	res := b.db.Collection(UsersCollection).FindOneAndDelete(ctx,
		bson.M{
			"google_id": id,
		},
	)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (b *Backend) RestoreUser(ctx context.Context, id string) error {
	ctx, cancel := b.TimeoutContext(ctx)
	defer cancel()

	res := b.db.Collection(UsersCollection).FindOneAndUpdate(ctx,
		bson.M{
			"google_id": id,
		},
		bson.M{
			"$unset": bson.M{
				"deleted_at": "",
				"deleted_by": "",
			},
		},
	)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}