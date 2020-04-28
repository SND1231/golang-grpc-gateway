package post_service

import (
	"app/post/db"
	"app/post/model"
	pb "app/post/proto"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//error_list [], err_msg string
func CreateError(code codes.Code, error_list []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: error_list,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequest_FieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckGetPostsRequest(request pb.GetPostsRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Limit == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Limit", "値が設定されていません"))
	}
	if request.Limit == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("id", "値が設定されていません"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckCreatePostRequest(request pb.CreatePostRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Tittle == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("内容", "必須です"))
	}
	if request.UserId == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("ユーザID", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckUpdatePostRequest(request pb.UpdatePostRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Tittle == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("内容", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckDeletePostRequest(request pb.DeletePostRequest) error {
	var post model.Post
	db := db.Connection()
	defer db.Close()

	id := request.Id
	user_id := request.UserId

	db.Where("id = ? AND user_id = ?", id, user_id).First(&post)
	fmt.Println(post.ID)
	if post.ID == 0 {
		return status.New(codes.AlreadyExists, "ユーザが違うか、投稿がすでに存在しません").Err()
	}
	return nil
}
