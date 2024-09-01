package repository

import (
	"context"
	"felagi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FreelancerRepositoryImpl struct {
	collection *mongo.Collection
}

func NewfreelancerRepositoryImpl(coll *mongo.Collection) domain.FreelancerRepository {
	return &FreelancerRepositoryImpl{collection: coll}
}

func (r *FreelancerRepositoryImpl) Register(user domain.FreelancerUser) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *FreelancerRepositoryImpl) GetUserByUsernameOrEmail(username, email string) (domain.FreelancerUser, error) {
	var user domain.FreelancerUser
	err := r.collection.FindOne(context.Background(), bson.M{"baseuser.email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}



func (r *FreelancerRepositoryImpl) ActivateAccount(email string) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"baseuser.email": email}, bson.M{"$set": bson.M{"baseuser.is_active": true}})
	if err != nil {
		return err
	}

	return nil
}



func (r *FreelancerRepositoryImpl) GetUserByEmail(email string) (domain.FreelancerUser, error) {
	var user domain.FreelancerUser
	err := r.collection.FindOne(context.Background(), bson.M{"baseuser.email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
