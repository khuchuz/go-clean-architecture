package repository

import (
	"context"
	"testing"

	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_CreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		id := (primitive.NewObjectID()).Hex()
		err := repo.CreateUser(context.Background(), &models.User{
			ID:       id,
			Username: "usermock",
			Email:    "usermock@gmail.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		})

		assert.Nil(t, err)
	})

	mt.Run("custom error duplicate", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		id := (primitive.NewObjectID()).Hex()
		err := repo.CreateUser(context.Background(), &models.User{
			ID:       id,
			Username: "usermock",
			Email:    "usermock@gmail.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		})

		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
	mt.Run("simple error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(bson.D{{"ok", 0}})

		id := (primitive.NewObjectID()).Hex()
		err := repo.CreateUser(context.Background(), &models.User{
			ID:       id,
			Username: "usermock",
			Email:    "usermock@gmail.com",
		})

		assert.NotNil(t, err)
	})
}

func Test_GetUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.ID},
			{Key: "username", Value: expectedUser.Username},
			{Key: "email", Value: expectedUser.Email},
			{Key: "password", Value: expectedUser.Password},
		}))

		user, err := repo.GetUser(context.Background(), expectedUser.Username, expectedUser.Password)
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	mt.Run("usernotfound", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "user not found",
		}))

		user, err := repo.GetUser(context.Background(), expectedUser.Username, expectedUser.Password)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	mt.Run("simple error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(bson.D{{"ok", 0}})

		user, err := repo.GetUser(context.Background(), expectedUser.Username, expectedUser.Password)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}

func Test_UpdatePassword(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.UpdatePassword(context.Background(), "usermock", "11f5639f22525155cb0b43573ee4212838c78d87")
		assert.Nil(t, err)
	})

	mt.Run("cannot find expected user", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "user not found",
		}))

		err := repo.UpdatePassword(context.Background(), expectedUser.Username, expectedUser.Password)
		assert.NotNil(t, err)
	})
}

func Test_IsUserExistByUsername(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success empty", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.IsUserExistByUsername(context.Background(), "usermock")
		assert.False(t, err)
	})

	mt.Run("success exist", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.ID},
			{Key: "username", Value: expectedUser.Username},
			{Key: "email", Value: expectedUser.Email},
			{Key: "password", Value: expectedUser.Password},
		}))

		err := repo.IsUserExistByUsername(context.Background(), expectedUser.Username)
		assert.True(t, err)
	})

	mt.Run("cannot find expected user", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "user not found",
		}))

		err := repo.IsUserExistByUsername(context.Background(), expectedUser.Username)
		assert.False(t, err)
	})
}

func Test_IsUserExistByEmail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success empty", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.IsUserExistByEmail(context.Background(), "john.doe@test.com")
		assert.False(t, err)
	})

	mt.Run("success exist", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.ID},
			{Key: "username", Value: expectedUser.Username},
			{Key: "email", Value: expectedUser.Email},
			{Key: "password", Value: expectedUser.Password},
		}))

		err := repo.IsUserExistByEmail(context.Background(), expectedUser.Username)
		assert.True(t, err)
	})

	mt.Run("cannot find expected user", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB, "users")
		expectedUser := &User{
			ID:       primitive.NewObjectID(),
			Username: "john",
			Email:    "john.doe@test.com",
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "user not found",
		}))

		err := repo.IsUserExistByEmail(context.Background(), expectedUser.Username)
		assert.False(t, err)
	})
}
