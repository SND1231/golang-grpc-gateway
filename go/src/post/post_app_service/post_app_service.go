package post_app_service

import (
	"app/post/db"
	"app/post/model"
	"app/post/post_service"
	pb "app/post/proto"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetPosts(request pb.GetPostsRequest) ([]*pb.Post, error) {
	var posts []model.Post
	var postList []*pb.Post

	err := post_service.CheckGetPostsRequest(request)
	if err != nil {
		return postList, err
	}

	limit := request.Limit
	offset := limit * (request.Offset - 1)

	db := db.Connection()
	defer db.Close()
	db.Limit(limit).Offset(offset).
		Find(&posts).Scan(&postList)

	return postList, nil
}

func GetPost(id int32) (pb.Post, error) {
	var post model.Post
	var post_param pb.Post

	db := db.Connection()
	defer db.Close()
	db.Find(&post, id).Scan(&post_param)

	return post_param, nil
}

func CreatePost(request pb.CreatePostRequest) (int32, error) {
	err := post_service.CheckCreatePostRequest(request)
	if err != nil {
		return -1, err
	}

	post_param := model.Post{Tittle: request.Tittle, Content: request.Content,
		PhotoUrl: request.PhotoUrl, UserId: request.UserId}

	db := db.Connection()
	defer db.Close()
	db.Create(&post_param)
	if db.NewRecord(post_param) == false {
		return post_param.ID, err
	}
	return -1, status.New(codes.Unknown, "作成失敗").Err()
}

func UpdatePost(request pb.UpdatePostRequest) (int32, error) {
	err := post_service.CheckUpdatePostRequest(request)
	if err != nil {
		return -1, err
	}

	id := request.Id

	post_param := model.Post{Tittle: request.Tittle, Content: request.Content,
		PhotoUrl: request.PhotoUrl}

	db := db.Connection()
	defer db.Close()
	post := model.Post{}
	db.Find(&post, id)

	db.Model(&post).UpdateColumns(post_param)
	return id, nil

}

func DeletePost(request pb.DeletePostRequest) (int32, error) {
	err := post_service.CheckDeletePostRequest(request)
	if err != nil {
		return -1, err
	}

	id := request.Id
	user_id := request.UserId

	db := db.Connection()
	defer db.Close()
	db.Where("id = ? AND user_id <> ?", id, user_id).Delete(model.Post{})
	return id, nil
}


func CreateLike(request pb.CreateLikeRequest) (int32, error) {
	err := post_service.LikeExists(request)
	if err != nil {
		return -1, err
	}
	var post model.Post
	db := db.Connection()
	defer db.Close()
	db.Where(&post, request.PostId)

	like := model.Like{UserId: request.UserId}
	db.Create(&like)
	db.Model(&post).Association("Likes").Append(&like)
	if db.NewRecord(like) == false {
		return like.ID, nil
	}
	return -1, status.New(codes.Unknown, "作成失敗").Err()
}

func DeleteLike(request pb.DeleteLikeRequest) (int32, error) {
	db := db.Connection()
	defer db.Close()
	db.Where("id = ?", request.Id).Delete(model.Like{})
	return request.Id, nil
}
