package filecomponentservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/summarizer/internal/services/filecomponentsservice/pb"
	"google.golang.org/grpc"
)

type FileComponentIds []int32

type FileComponent struct {
	Id        int
	UserId    string
	FilePath  string
	StartLine int
	EndLine   int
	Content   string
}

func GetFileComponents(fileComponentIds FileComponentIds) ([]FileComponent, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.FileComponentIds{
		FileComponentIds: fileComponentIds,
	}

	resp, err := client.GetFileComponents(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var fileComponents []FileComponent
	for _, pbFileComponent := range resp.FileComponents {
		fileComponents = append(fileComponents, FileComponent{
			Id:        int(pbFileComponent.Id),
			UserId:    pbFileComponent.UserId,
			FilePath:  pbFileComponent.FilePath,
			StartLine: int(pbFileComponent.StartLine),
			EndLine:   int(pbFileComponent.EndLine),
			Content:   pbFileComponent.Content,
		})
	}

	return fileComponents, nil
}
