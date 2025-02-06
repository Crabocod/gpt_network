package grpc

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	pb "github.com/Crabocod/gpt_network/api-service/internal/proto"
)

type GrpcController struct {
	pb.UnimplementedApiServiceServer
	service service.Service
}

func New(service service.Service) *GrpcController {
	return &GrpcController{
		service: service,
	}
}

//func (s *GrpcController) SaveComment(ctx context.Context, req *pb.SaveCommentRequest) (*pb.SaveCommentResponse, error) {
//	var comment models.Comment
//
//	User, err := models.GetUserByName(req.GetAuthorName())
//	if err != nil {
//		return &pb.SaveCommentResponse{Success: false}, err
//	}
//
//	comment.Text = req.GetText()
//	comment.PostID, _ = strconv.Atoi(req.GetPostId())
//	comment.AuthorID = User.ID
//	err = comment.Save()
//	if err != nil {
//		return &pb.SaveCommentResponse{Success: false}, err
//	}
//
//	return &pb.SaveCommentResponse{Success: true}, nil
//}
//
//func (s *GrpcController) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
//	Post, err := models.GetLatestFilteredPost(req.GetAuthorName())
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, status.Errorf(codes.NotFound, "Post not found for author: %s", req.GetAuthorName())
//		}
//		return nil, status.Errorf(codes.Internal, "Error retrieving post: %v", err)
//	}
//
//	return &pb.GetPostResponse{
//		PostId:   strconv.Itoa(Post.ID),
//		PostText: Post.Text,
//	}, nil
//}
//
//func (s *GrpcController) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
//	var post models.Post
//
//	User, err := models.GetUserByName(req.GetAuthorName())
//	if err != nil {
//		return &pb.SavePostResponse{Success: false}, err
//	}
//
//	post.Text = req.GetText()
//	post.AuthorID = User.ID
//	err = post.Save()
//	if err != nil {
//		return &pb.SavePostResponse{Success: false}, err
//	}
//
//	return &pb.SavePostResponse{Success: true}, nil
//}
