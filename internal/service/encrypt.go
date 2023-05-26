package service

import (
	"context"
	"crypto/aes"
	"encoding/hex"
	"encrypt-decrypt/internal/pb"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Encrypt struct {
	pb.UnimplementedEncryptServiceServer
	Key string
}

func (e *Encrypt) EncryptContent(ctx context.Context, request *pb.EncryptContentRequest) (*pb.EncryptContentResponse, error) {
	content, err := e.EncryptAes([]byte(e.Key), request.Content)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}
	return &pb.EncryptContentResponse{Content: content}, nil
}

func (e *Encrypt) EncryptFile(ctx context.Context, in *pb.EncryptFileRequest) (*pb.EncryptFileResponse, error) {
	return &pb.EncryptFileResponse{}, nil
}

func (e *Encrypt) EncryptAes(key []byte, text string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(text))
	c.Encrypt(out, []byte(text))
	fmt.Println(out)
	return hex.EncodeToString(out), nil
}

func NewEncrypt(key string) *Encrypt {
	return &Encrypt{
		Key: key,
	}
}
