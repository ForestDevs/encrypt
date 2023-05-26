package service

import (
	"context"
	"crypto/aes"
	"encoding/hex"
	"encrypt-decrypt/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Decrypt struct {
	pb.UnimplementedDecryptServiceServer
	Key string
}

func (d *Decrypt) DecryptContent(ctx context.Context, request *pb.DecryptContentRequest) (*pb.DecryptContentResponse, error) {
	content, err := d.DecryptAes([]byte(d.Key), "7a75f705a142dc5e3978928d19febf1700000000000000000000000000000000000000000000000000000000")
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}
	return &pb.DecryptContentResponse{Content: content}, nil
}

func (d *Decrypt) DecryptFile(ctx context.Context, request *pb.DecryptFileRequest) (*pb.DecryptFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Decrypt) DecryptAes(key []byte, text string) (string, error) {
	ciphertext, _ := hex.DecodeString(text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	return s, nil
}

func NewDecrypt(key string) *Decrypt {
	return &Decrypt{Key: key}
}
