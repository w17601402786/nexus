// Code generated by hertz generator. DO NOT EDIT.

package user_microservice

import (
	user_microservice "github.com/AdrianWangs/nexus/go-service/user/biz/handler/user_microservice"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/get_user", append(_getuserMw(), user_microservice.GetUser)...)
	root.POST("/login", append(_loginMw(), user_microservice.Login)...)
	root.POST("/register", append(_registerMw(), user_microservice.Register)...)
	root.POST("/third_party_login", append(_thirdpartyloginMw(), user_microservice.ThirdPartyLogin)...)
	root.POST("/update_user_profile", append(_updateuserprofileMw(), user_microservice.UpdateUserProfile)...)
}