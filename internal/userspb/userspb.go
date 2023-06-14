package userspb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/msik-404/micro-appoint-users/internal/database"
	"github.com/msik-404/micro-appoint-users/internal/models"
)

type Server struct {
	UnimplementedApiServer
	Client mongo.Client
}

func (s *Server) FindOneCustomer(
	ctx context.Context,
	request *CustomerRequest,
) (*CustomerReply, error) {
	customerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	type Customer struct {
		Mail    *string `bson:"mail,omitempty"`
		Name    *string `bson:"name,omitempty"`
		Surname *string `bson:"surname,omitempty"`
	}
	var result Customer
	err = models.FindOneCustomer(ctx, db, customerID).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &CustomerReply{
		Mail:    result.Mail,
		Name:    result.Name,
		Surname: result.Surname,
	}, nil
}

func (s *Server) FindOneOwner(
	ctx context.Context,
	request *OwnerRequest,
) (*OwnerReply, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	type Owner struct {
		Mail      *string              `bson:"mail,omitempty"`
		Name      *string              `bson:"name,omitempty"`
		Surname   *string              `bson:"surname,omitempty"`
		Companies []primitive.ObjectID `bson:"companies,omitempty"`
	}
	var result Owner
	err = models.FindOneOwner(ctx, db, ownerID).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	reply := OwnerReply{
		Mail:    result.Mail,
		Name:    result.Name,
		Surname: result.Surname,
	}
	for _, id := range result.Companies {
		reply.Companies = append(reply.Companies, id.Hex())
	}
	return &reply, nil
}

type Credentials struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	HashedPwd string             `bson:"pwd,omitempty"`
}

func (s *Server) FindOneCustomerCredentials(
	ctx context.Context,
	request *CustomerCredentialsRequest,
) (*CredentialsReply, error) {
	if request.Mail == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Mail field is required",
		)
	}
	err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	var result Credentials
	err = models.FindOneCustomerCredentials(ctx, db, request.GetMail()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	customerID := result.ID.Hex()
    hashedPwd := result.HashedPwd
	reply := CredentialsReply{
		Id:        &customerID,
		HashedPwd: &hashedPwd,
	}
	return &reply, nil
}

func (s *Server) FindOneOwnerCredentials(
	ctx context.Context,
	request *OwnerCredentialsRequest,
) (*CredentialsReply, error) {
	if request.Mail == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Mail field is required",
		)
	}
	err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	var result Credentials
	err = models.FindOneOwnerCredentials(ctx, db, request.GetMail()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	ownerID := result.ID.Hex()
    hashedPwd := result.HashedPwd
	reply := CredentialsReply{
		Id:        &ownerID,
		HashedPwd: &hashedPwd,
	}
	return &reply, nil
}

func (s *Server) AddCustomer(
	ctx context.Context,
	request *AddCustomerRequest,
) (*emptypb.Empty, error) {
	if request.Mail == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Mail field is required",
		)
	}
	err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.HashedPwd == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Hashed password field is required",
		)
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	newCustomer := models.Customer{
		Mail:      request.GetMail(),
		HashedPwd: request.GetHashedPwd(),
		Name:      request.GetName(),
		Surname:   request.GetSurname(),
	}
	db := s.Client.Database(database.DBName)
	_, err = newCustomer.InsertOne(ctx, db)
	if err != nil {
        if mongo.IsDuplicateKeyError(err) {
		    return nil, status.Error(codes.AlreadyExists, err.Error())
        }
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) AddOwner(
	ctx context.Context,
	request *AddOwnerRequest,
) (*emptypb.Empty, error) {
	if request.Mail == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Mail field is required",
		)
	}
	err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.HashedPwd == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Hashed password field is required",
		)
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	newOwner := models.Owner{
		Mail:      request.GetMail(),
		HashedPwd: request.GetHashedPwd(),
		Name:      request.GetName(),
		Surname:   request.GetSurname(),
	}
	db := s.Client.Database(database.DBName)
	_, err = newOwner.InsertOne(ctx, db)
	if err != nil {
        if mongo.IsDuplicateKeyError(err) {
		    return nil, status.Error(codes.AlreadyExists, err.Error())
        }
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) AddOwnedCompany(
	ctx context.Context,
	request *AddOwnedCompanyRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	companyID, err := primitive.ObjectIDFromHex(request.GetCompanyId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
    db := s.Client.Database(database.DBName)
	result, err := models.InsertOneOwnedCompany(ctx, db, ownerID, companyID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Owner with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteOwnedCompany(
	ctx context.Context,
	request *DeleteOwnedCompanyRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	companyID, err := primitive.ObjectIDFromHex(request.GetCompanyId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
    db := s.Client.Database(database.DBName)
	result, err := models.DeleteOneOwnedCompany(ctx, db, ownerID, companyID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Owner with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateCustomer(
	ctx context.Context,
	request *UpdateCustomerRequest,
) (*emptypb.Empty, error) {
	customerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	customerUpdate := models.CustomerUpdate{
		Mail:      request.Mail,
		HashedPwd: request.HashedPwd,
		Name:      request.Name,
		Surname:   request.Surname,
	}
	db := s.Client.Database(database.DBName)
	result, err := customerUpdate.UpdateOne(ctx, db, customerID)
	if err != nil {
        if mongo.IsDuplicateKeyError(err) {
		    return nil, status.Error(codes.AlreadyExists, err.Error())
        }
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Customer with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateOwner(
	ctx context.Context,
	request *UpdateOwnerRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ownerUpdate := models.OwnerUpdate{
		Mail:      request.Mail,
		HashedPwd: request.HashedPwd,
		Name:      request.Name,
		Surname:   request.Surname,
	}
	db := s.Client.Database(database.DBName)
	result, err := ownerUpdate.UpdateOne(ctx, db, ownerID)
	if err != nil {
        if mongo.IsDuplicateKeyError(err) {
		    return nil, status.Error(codes.AlreadyExists, err.Error())
        }
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Owner with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil

}

func (s *Server) DeleteCustomer(
	ctx context.Context,
	request *DeleteCustomerRequest,
) (*emptypb.Empty, error) {
	customerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	result, err := models.DeleteOneCustomer(ctx, db, customerID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Customer with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteOwner(
	ctx context.Context,
	request *DeleteOwnerRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	result, err := models.DeleteOneOwner(ctx, db, ownerID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Owner with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}
