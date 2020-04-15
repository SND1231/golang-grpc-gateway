package app_service

import (
	"app/user/db"
	"app/user/model"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    pb "app/user/pb"
    errors "app/errors"
)

func GetUserList() ([]*pb.User, error) {
    var users []model.User
    var userList []*pb.User

    db := db.Connection()
    defer db.Close()
    db.Find(&users).Scan(&userList)

    return userList, nil
}

func GetUser(id int32) (pb.User, error) {
    var users model.User
    var userList pb.User

    db := db.Connection()
    defer db.Close()
    db.Find(&users, id).Scan(&userList)

    return userList, nil
}

func CreateUser(request pb.CreateUserRequest) (int32, error) {
    user_param := model.User{Name: request.Name, Email: request.Email,
                             PhotoUrl: request.PhotoUrl}

    db := db.Connection()
    defer db.Close()

    db.Create(&user_param)
    if db.NewRecord(user_param) == false {
        return user_param.ID, nil
    }
    return 0, errors.NewError("fail to create user")
}

func UpdateUser(request pb.UpdateUserRequest) (int32, error) {

    user_param := model.User{Name: request.Name, Email: request.Email,
                             PhotoUrl: request.PhotoUrl}
    var id = request.Id

    db := db.Connection()
    defer db.Close()
    user := model.User{}
    db.Find(&user, id)
    
    db.Model(&user).UpdateColumns(user_param)
    return id, nil
    
}
