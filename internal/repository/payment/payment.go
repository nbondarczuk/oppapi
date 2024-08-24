package payment

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"oppapi/internal/logging"
	"oppapi/internal/model"
	"oppapi/internal/repository"
)

const PaymentCollectionName = "payment"

// PaymentRepository is a container for resource accerss action state.
type PaymentRepository struct {
	repository *repository.MongoRepository
	ctx        context.Context
	collection *mongo.Collection
}

// NewPaymentRepository handles resource access action in its own context.
func NewPaymentRepository() (*PaymentRepository, error) {
	repository, err := repository.WithMongo()
	if err != nil {
		return nil, err
	}
	collection := repository.Client.Database(repository.DBName).Collection(PaymentCollectionName)
	tc := PaymentRepository{
		repository: repository,
		ctx:        context.Background(),
		collection: collection,
	}
	return &tc, nil
}

// Create an object with new oid allocated.
func (tc *PaymentRepository) Create(payment *model.Payment) (*model.Payment, error) {
	logging.Logger.Debug("Creating payment")
	if payment.ID.IsZero() {
		payment.ID = primitive.NewObjectID()
	}
	payment.Created = time.Now()
	result, err := tc.collection.InsertOne(tc.ctx, payment)
	if err != nil {
		return nil, err
	}
	logging.Logger.Debug("Created payment", slog.String("ID", fmt.Sprintf("%v", result.InsertedID)))
	return payment, nil
}

// ReadOne fetches one object by primary key.
func (tc *PaymentRepository) ReadOne(id string) (model.Payment, error) {
	logging.Logger.Debug("Reading payment")
	var payment model.Payment
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Payment{}, err
	}
	err = tc.collection.FindOne(tc.ctx, bson.M{"_id": ID}).Decode(&payment)
	if err != nil {
		return model.Payment{}, err
	}
	logging.Logger.Debug("Read payment", slog.String("ID", fmt.Sprintf("%v", payment.ID)))
	return payment, nil
}

// SetStatus changes the payment record
func (tc *PaymentRepository) SetStatus(id string, status string) error {
	updated := time.Now()
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = tc.collection.UpdateOne(tc.ctx,
		bson.M{"_id": ID},
		bson.D{{"$set",
			bson.D{
				{"Status", status},
				{"Updated", updated},
			}}})
	if err != nil {
		return err
	}
	logging.Logger.Debug("Updated payment status", slog.String("ID", fmt.Sprintf("%v", ID)))
	return nil
}
