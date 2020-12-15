package mongo

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName   = "productsDB"
	collName = "products"
)

type product struct {
	InternalID primitive.ObjectID `bson:"_id, omitempty"`
	pb.Product `bson:",inline"`
}

type Storage struct {
	client   *mongo.Client
	products *mongo.Collection
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
	return &Storage{products: client.Database(dbName).Collection(collName)}, nil
}

func (s *Storage) Fetch(ctx context.Context) ([]*pb.Product, error) {
	cur, err := s.products.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	products := make([]*product, 0, cur.RemainingBatchLength())
	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}
	result := make([]*pb.Product, 0, len(products))
	for _, p := range products {
		p.Id = p.InternalID.Hex()
		result = append(result, &p.Product)
	}
	return result, nil
}
