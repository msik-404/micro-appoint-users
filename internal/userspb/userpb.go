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

func (s *Server) FindOneCustomerCredentials(
	ctx context.Context,
	request *CustomerCredentialsRequest,
) (*CredentialsReply, error) {
	mail, err := verifyString(&request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	customerModel := models.Customer{}
	err = models.FindOneCustomerCredentials(ctx, db, *mail).Decode(&customerModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	reply := CredentialsReply{
		Id:        customerModel.ID.Hex(),
		HashedPwd: customerModel.HashedPwd,
	}
	return &reply, nil
}

func (s *Server) FindOneOwnerCredentials(
	ctx context.Context,
	request *OwnerCredentialsRequest,
) (*CredentialsReply, error) {
	mail, err := verifyString(&request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	db := s.Client.Database(database.DBName)
	ownerModel := models.Owner{}
	err = models.FindOneOwnerCredentials(ctx, db, *mail).Decode(&ownerModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	reply := CredentialsReply{
		Id:        ownerModel.ID.Hex(),
		HashedPwd: ownerModel.HashedPwd,
	}
	return &reply, nil
}

func (s *Server) FindManyOwnerCompanies(
	ctx context.Context,
	request *OwnerCompaniesRequest,
) (*OwnerCompaniesReply, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	db := s.Client.Database(database.DBName)
	ownerModel := models.Owner{}
	err = models.FindOneOwnerCompanies(ctx, db, ownerID).Decode(&ownerModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(ownerModel.Companies) == 0 {
		return nil, status.Error(
			codes.NotFound,
			"This owner does not own any companies",
		)
	}
	reply := OwnerCompaniesReply{Id: ownerModel.ID.Hex()}
	for _, companieID := range ownerModel.Companies {
		reply.Companies = append(reply.Companies, companieID.Hex())
	}
	return &reply, nil
}

func (s *Server) AddCustomer(
	ctx context.Context,
	request *AddCustomerRequest,
) (*emptypb.Empty, error) {
	mail, err := verifyString(&request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.HashedPwd == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Hashed password field is required",
		)
	}
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	newCustomer := models.Customer{
		Mail:      mail,
		HashedPwd: request.HashedPwd,
		Name:      name,
		Surname:   surname,
	}
	db := s.Client.Database(database.DBName)
	_, err = newCustomer.InsertOne(ctx, db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) AddOwner(
	ctx context.Context,
	request *AddOwnerRequest,
) (*emptypb.Empty, error) {
	mail, err := verifyString(&request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.HashedPwd == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Hashed password field is required",
		)
	}
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	newOwner := models.Owner{
		Mail:      mail,
		HashedPwd: request.HashedPwd,
		Name:      name,
		Surname:   surname,
	}
	db := s.Client.Database(database.DBName)
	_, err = newOwner.InsertOne(ctx, db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) AddOwnerCompany(
	ctx context.Context,
	request *AddOwnerCompanyRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	if request.CompanyId == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"CompanyID field is required",
		)
	}
	db := s.Client.Database(database.DBName)
	companyID, err := primitive.ObjectIDFromHex(*request.CompanyId)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	result, err := models.InsertOneOwnerCompany(ctx, db, ownerID, companyID)
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

func (s *Server) DeleteOwnerCompany(
    ctx context.Context,
    request *DeleteOwnerCompanyRequest,
) (*emptypb.Empty, error) {
	ownerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	if request.CompanyId == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"CompanyID field is required",
		)
	}
	db := s.Client.Database(database.DBName)
	companyID, err := primitive.ObjectIDFromHex(*request.CompanyId)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	result, err := models.DeleteOneOwnerCompany(ctx, db, ownerID, companyID)
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
	customerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	mail, err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
    customerUpdate := models.Customer{
        Mail: mail,
        HashedPwd: request.HashedPwd,
        Name: name,
        Surname: surname,
    }
	db := s.Client.Database(database.DBName)
    result, err := customerUpdate.UpdateOne(ctx, db, customerID)
	if err != nil {
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
	ownerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	mail, err := verifyString(request.Mail, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
    ownerUpdate := models.Owner{
        Mail: mail,
        HashedPwd: request.HashedPwd,
        Name: name,
        Surname: surname,
    }
	db := s.Client.Database(database.DBName)
    result, err := ownerUpdate.UpdateOne(ctx, db, ownerID)
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

func (s *Server) DeleteCustomer(
    ctx context.Context, 
    request *DeleteCustomerRequest,
) (*emptypb.Empty, error) {
	customerID, err := primitive.ObjectIDFromHex(request.Id)
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
	ownerID, err := primitive.ObjectIDFromHex(request.Id)
    if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	db := s.Client.Database(database.DBName)
    result, err := models.DeleteOneCustomer(ctx, db, ownerID)
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
