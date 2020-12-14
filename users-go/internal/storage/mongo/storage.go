package mongo

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName   = "usersDB"
	collName = "users"
)

type user struct {
	InternalID primitive.ObjectID `bson:"_id, omitempty"`
	pb.User    `bson:",inline"`
}

type Storage struct {
	client *mongo.Client
	users  *mongo.Collection
}

func New(ctx context.Context, opts ...*options.ClientOptions) (*Storage, error) {
	client, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return &Storage{users: client.Database(dbName).Collection(collName)}, nil
}

func (s *Storage) Fetch(ctx context.Context) ([]*pb.User, error) {
	cur, err := s.users.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	users := make([]*user, 0, cur.RemainingBatchLength())
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	result := make([]*pb.User, 0, len(users))
	for _, u := range users {
		u.Id = u.InternalID.Hex()
		result = append(result, &u.User)
	}
	return result, nil
}
