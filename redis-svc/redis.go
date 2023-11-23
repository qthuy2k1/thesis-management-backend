package main

import (
	"context"
	"fmt"

	redispb "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redispb.UnimplementedRedisServiceServer
	Redis *redis.Client
}

func NewRedis(redis *redis.Client) *Redis {
	return &Redis{
		Redis: redis,
	}
}

type UserRedis struct {
	Id             string `redis:"id" json:"id"`
	Class          string `redis:"class,omitempty" json:"class,omitempty"`
	Major          string `redis:"major,omitempty" json:"major,omitempty"`
	Phone          string `redis:"phone,omitempty" json:"phone,omitempty"`
	PhotoSrc       string `redis:"photoSrc" json:"photoSrc`
	Role           string `redis:"role" json:"role"`
	Name           string `redis:"name" json:"name"`
	Email          string `redis:"email" json:"email"`
	HashedPassword string `redis:"hashed_password" json:"hashed_password"`
}

func (r *Redis) GetUser(ctx context.Context, req *redispb.GetUserRequest) (*redispb.GetUserResponse, error) {
	res := r.Redis.HGetAll(ctx, fmt.Sprintf("user:%s", req.Id))

	if len(res.Val()) == 0 {
		return &redispb.GetUserResponse{
			Response: &redispb.CommonRedisResponse{
				StatusCode: 404,
				Message:    "user not found",
			},
		}, nil
	}

	var userScan UserRedis
	if err := res.Scan(&userScan); err != nil {
		return nil, err
	}

	return &redispb.GetUserResponse{
		Response: &redispb.CommonRedisResponse{
			StatusCode: 200,
			Message:    "found user",
		},
		User: &redispb.User{
			Id:             userScan.Id,
			Class:          &userScan.Class,
			Major:          &userScan.Major,
			Phone:          &userScan.Phone,
			PhotoSrc:       userScan.PhotoSrc,
			Role:           userScan.Role,
			Name:           userScan.Name,
			Email:          userScan.Email,
			HashedPassword: &userScan.HashedPassword,
		},
	}, nil
}

func (r *Redis) SetUser(ctx context.Context, req *redispb.SetUserRequest) (*redispb.SetUserResponse, error) {
	class := ""
	if req.User.Class != nil {
		class = *req.User.Class
	}

	phone := ""
	if req.User.Phone != nil {
		phone = *req.User.Phone
	}

	major := ""
	if req.User.Major != nil {
		major = *req.User.Major
	}

	hashedPassword := ""
	if req.User.HashedPassword != nil {
		hashedPassword = *req.User.HashedPassword
	}

	userCache := UserRedis{
		Id:             req.User.Id,
		Class:          class,
		Major:          major,
		Phone:          phone,
		PhotoSrc:       req.User.PhotoSrc,
		Role:           req.User.Role,
		Name:           req.User.Name,
		Email:          req.User.Email,
		HashedPassword: hashedPassword,
	}

	if err := r.Redis.HSet(ctx, fmt.Sprintf("user:%s", req.User.Id), userCache); err.Err() != nil {
		return nil, err.Err()
	}

	return &redispb.SetUserResponse{
		Response: &redispb.CommonRedisResponse{
			StatusCode: 200,
			Message:    "success",
		},
	}, nil
}
