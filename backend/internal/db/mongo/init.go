package mongo

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collections = []string{
		"points_transactions",
		"stars",
		"suggestions",
		"news",
		"calendar",
		"users",
	}

	// Add indexes to index "id" which is text & unique
	indexes = map[string][]mongo.IndexModel{
		"points_transactions": {
			// index descending by created_at
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": -1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			},
		},
		"stars": {
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": 1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			}},
		"suggestions": {
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": 1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			}},
		"news": {
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": 1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			}},
		"calendar": {
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": 1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			}},
		"user": {
			mongo.IndexModel{
				Keys: bson.M{
					"created_at": 1,
				},
			},
			mongo.IndexModel{
				Keys: bson.M{
					"id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"id": bson.M{
						"$exists": true,
					},
				}),
			},
			mongo.IndexModel{
				Keys: bson.M{
					"google_id": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"google_id": bson.M{
						"$exists": true,
					},
				}),
			},
			mongo.IndexModel{
				Keys: bson.M{
					"email_address": 1,
				},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"email_address": bson.M{
						"$exists": true,
					},
				}),
			},
		},
	}

	PointsTransactionCollection   = "points_transactions"
	StarsCollection        = "stars"
	SuggestionsCollection          = "suggestions"
	NewsCollection     = "news"
	CalendarCollection  = "calendar"
	UsersCollection = "users"
)

func (b *Backend) CreateCollections() error {
	ctx, cancel := b.GetContext()
	defer cancel()

	for _, collection := range Collections {
		b.db.CreateCollection(ctx, collection)

		if err := b.CreateIndexes(collection); err != nil {
			logrus.Error(err)
			continue
		}
	}

	return nil
}

func (b *Backend) CreateIndexes(collection string) error {
	ctx, cancel := b.GetContext()
	defer cancel()

	v, ok := indexes[collection]
	if !ok {
		return nil
	}

	_, err := b.db.Collection(collection).Indexes().CreateMany(ctx, v)
	if err != nil {
		return err
	}

	return nil
}