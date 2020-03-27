package file_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/proto_file"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"strconv"
)

var Conn *grpc.ClientConn

func FileServerInit() {
	var err error
	Conn, err = grpc.Dial(viper.GetString("FileServer.Addr"), grpc.WithInsecure())
	if err != nil {
		logrus.Error("[FileServerInit] err msg: ", err.Error())
	}
}

func UploadAvatar(c context.Context, request *json_struct.UploadAvatarRequest) (*json_struct.UploadAvatarResponse, error) {
	client := proto_file.NewFileServerClient(Conn)
	id := util.GenId()
	resp, err := client.UploadFile(c, &proto_file.UploadFileRequsest{
		Id:                   id,
		FileContent:          request.FileContent,
		FileType:             "png",
	})
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, errors.New("status isnt success")
	}
	return &json_struct.UploadAvatarResponse{
		ImageId: fmt.Sprintf("%d", id),
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}
func GetAvatoar(c context.Context, request *json_struct.GetAvatarRequest) (*json_struct.GetAvatarResponse, error) {
	client := proto_file.NewFileServerClient(Conn)

	id, err := strconv.ParseInt(request.ImageId, 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := client.DownloadFile(c, &proto_file.DownloadFileRequest{
		Id:                   id,
		FileType:             "png",
	})
	if err != nil {
		return nil, err
	}
	return &json_struct.GetAvatarResponse{
		ImageFile: resp.FileContent,
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}