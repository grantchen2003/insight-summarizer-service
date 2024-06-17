package filechunksservice

import (
	"context"
	"os"
	"strings"

	"github.com/grantchen2003/insight/summarizer/internal/services/filechunksservice/pb"
	"google.golang.org/grpc"
)

type UserFilePath struct {
	UserId   string
	Filepath string
}

func getFileContent(userFilePath UserFilePath) (string, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_CHUNKS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewFileChunksServiceClient(conn)
	response, err := client.GetSortedFileChunksContent(context.Background(), &pb.GetSortedFileChunksContentRequest{
		UserId:   userFilePath.UserId,
		FilePath: userFilePath.Filepath,
	})
	if err != nil {
		return "", err
	}

	var fileChunkContents []string

	for _, fileChunkContent := range response.FileChunksContent {
		fileChunkContents = append(fileChunkContents, string(fileChunkContent.Content))
	}

	return strings.Join(fileChunkContents, ""), nil
}

func GetFileContents(userFilePaths []UserFilePath) ([]string, error) {
	var fileContents []string

	for _, userFilePath := range userFilePaths {
		fileContent, err := getFileContent(userFilePath)
		if err != nil {
			return nil, err
		}
		fileContents = append(fileContents, fileContent)
	}

	return fileContents, nil
}
