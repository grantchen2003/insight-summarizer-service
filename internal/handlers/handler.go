package handlers

import (
	"context"
	"log"
	"strings"

	db "github.com/grantchen2003/insight/summarizer/internal/database"
	fileChunksService "github.com/grantchen2003/insight/summarizer/internal/services/filechunksservice"
	fileComponentsService "github.com/grantchen2003/insight/summarizer/internal/services/filecomponentsservice"
	"github.com/grantchen2003/insight/summarizer/internal/utils/filecomponentsummarizer"

	pb "github.com/grantchen2003/insight/summarizer/internal/protobufs"
)

type SummarizerServiceHandler struct {
	pb.SummarizerServiceServer
}

func (s *SummarizerServiceHandler) CreateFileComponentSummaries(ctx context.Context, req *pb.FileComponentIds) (*pb.FileComponentSummaries, error) {
	log.Println("received CreateFileComponentSummaries request")

	database := db.GetSingletonInstance()

	fileComponents, err := fileComponentsService.GetFileComponents(req.FileComponentIds)
	if err != nil {
		return nil, err
	}

	var userFilePaths []fileChunksService.UserFilePath
	for _, fileComponent := range fileComponents {
		userFilePaths = append(userFilePaths, fileChunksService.UserFilePath{
			UserId:   fileComponent.UserId,
			Filepath: fileComponent.FilePath,
		})
	}

	fileContents, err := fileChunksService.GetFileContents(userFilePaths)
	if err != nil {
		return nil, err
	}

	var fileComponentSummaryPayloads []db.FileComponentPayload

	for i, fileComponentId := range req.FileComponentIds {

		fileComponent := fileComponents[i]
		fileContent := fileContents[i]

		fileComponentContent := getFileComponentContent(fileComponent, fileContent)
		// TODO: actually implement this
		summary := filecomponentsummarizer.SummarizeFileComponentContent(fileComponentContent)

		fileComponentSummaryPayloads = append(
			fileComponentSummaryPayloads,
			db.FileComponentPayload{
				FileComponentId: int(fileComponentId),
				Summary:         summary,
			},
		)
	}

	fileComponentSummaries, err := database.BatchSaveFileComponentSummaries(fileComponentSummaryPayloads)
	if err != nil {
		return nil, err
	}

	resp := &pb.FileComponentSummaries{
		FileComponentSummaries: castToPbFileComponentSummaries(fileComponentSummaries),
	}

	return resp, nil
}

func getFileComponentContent(fileComponent fileComponentsService.FileComponent, fileContent string) string {
	fileContentLines := strings.Split(fileContent, "\n")

	fileComponentContent := strings.Join(fileContentLines[fileComponent.StartLine-1:fileComponent.EndLine], "\n")

	return fileComponentContent
}

func castToPbFileComponentSummaries(fileComponentSummaries []db.FileComponentSummary) []*pb.FileComponentSummary {
	var pbFileComponentSummaries []*pb.FileComponentSummary

	for _, fileComponentSummary := range fileComponentSummaries {
		pbFileComponentSummaries = append(pbFileComponentSummaries, &pb.FileComponentSummary{
			Id:              int32(fileComponentSummary.Id),
			FileComponentId: int32(fileComponentSummary.FileComponentId),
			Summary:         fileComponentSummary.Summary,
		})
	}

	return pbFileComponentSummaries
}
