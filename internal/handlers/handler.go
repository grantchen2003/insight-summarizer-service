package handlers

import (
	"context"
	"log"

	db "github.com/grantchen2003/insight/summarizer/internal/database"
	fileComponentsService "github.com/grantchen2003/insight/summarizer/internal/services/filecomponentsservice"
	"github.com/grantchen2003/insight/summarizer/internal/utils"

	pb "github.com/grantchen2003/insight/summarizer/internal/protobufs"
)

type SummarizerServiceHandler struct {
	pb.SummarizerServiceServer
}

func (s *SummarizerServiceHandler) CreateFileComponentSummaries(ctx context.Context, req *pb.CreateFileComponentSummariesRequest) (*pb.FileComponentSummaries, error) {
	log.Println("received CreateFileComponentSummaries request")

	fileComponents, err := fileComponentsService.GetFileComponents(req.FileComponentIds)
	if err != nil {
		return nil, err
	}

	var fileComponentSummaryPayloads []db.FileComponentPayload

	for i, fileComponentId := range req.FileComponentIds {
		summary := utils.SummarizeSourceCode(fileComponents[i].Content)

		fileComponentSummaryPayloads = append(
			fileComponentSummaryPayloads,
			db.FileComponentPayload{
				FileComponentId: int(fileComponentId),
				Summary:         summary,
			},
		)
	}

	database := db.GetSingletonInstance()

	fileComponentSummaries, err := database.BatchSaveFileComponentSummaries(fileComponentSummaryPayloads)
	if err != nil {
		return nil, err
	}

	resp := &pb.FileComponentSummaries{
		FileComponentSummaries: castToPbFileComponentSummaries(fileComponentSummaries),
	}

	return resp, nil
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
