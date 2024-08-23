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
	"oppapi/internal/repository"
)

const PaymentCollectionName = "payment"

// swagger:model isocurrencycode
type ISOCurrencyCode [3]byte

// swagger:model creditcard
type CreditCard struct {
	NameAndSurname string `json:"nameandsurename" bson:"nameandsurename"`
	CardNo         string `json:"cardno" bson:"cardno"`
	CCV            string `json:"ccv" bson:"ccv"`
	ExpiryDate     string `json:"expirtdate" bson:"expirtdate"`
}

// swagger:model paymentmethod
type PaymentMethod struct {
	CreditCard `json:"creditcard" bson:"creditcard"`
}

// Payment is the entity mnaged by the repository.
// swagger:model payment
type Payment struct {
	// the id of the payment
	// required: true
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// the amount of the payment
	// required: true
	Amount string `json:"amount" bson:"amount"`

	// the currency of the amount of the payment
	// required: true
	Currency ISOCurrencyCode `json:"currency" bson:"currency"`

	// the payment methid, it may be Payment card
	Method PaymentMethod `json:"method" bson:"method"`

	// date of creation of the payment
	// required: true
	Created time.Time `json:"created" bson:"created"`
}

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
func (tc *PaymentRepository) Create(payment *Payment) (*Payment, error) {
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
func (tc *PaymentRepository) ReadOne(id string) (Payment, error) {
	var payment Payment
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Payment{}, err
	}
	err = tc.collection.FindOne(tc.ctx, bson.M{"_id": ID}).Decode(&payment)
	if err != nil {
		return Payment{}, err
	}
	logging.Logger.Debug("Read payment", slog.String("ID", fmt.Sprintf("%v", payment.ID)))
	return payment, nil
}
