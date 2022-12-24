package handlers

import (
	"time"

	skewauth "github.com/Skewax/backend/internal/auth"
	"github.com/Skewax/backend/pkg/swagger/server/models"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/files"

	"github.com/go-openapi/runtime/middleware"
)

var _ = skewauth.Begin

func HandleGetFiles(files.GetFilesParams) middleware.Responder {
	returnedFiles := []*models.BasicFileObject{}
	returnedFiles = append(returnedFiles, &models.BasicFileObject{"TESTID1", "TESTNAME1"})
	returnedFiles = append(returnedFiles, &models.BasicFileObject{"TESTID2", "TESTNAME2"})
	returnedFiles = append(returnedFiles, &models.BasicFileObject{"TESTID3", "TESTNAME3"})

	body := files.GetFilesOKBody{Files: returnedFiles}
	return files.NewGetFilesOK().WithPayload(&body)
}

func HandleReadFile(files.ReadFileParams) middleware.Responder {
	body := models.FileDataResponse{"", "TESTID1", time.Now().Unix(), "TEST TEXT\n HI", true}
	return files.NewReadFileOK().WithPayload(&body)
}

func HandleUpdateFile(files.UpdateFileParams) middleware.Responder {
	body := models.BasicResponse{""}
	return files.NewUpdateFileOK().WithPayload(&body)
}

func HandleCreateFile(files.CreateFileParams) middleware.Responder {
	body := files.CreateFileOKBody{"", &models.BasicFileObject{"TESTIDNEW", "NEW FILE"}}
	return files.NewCreateFileOK().WithPayload(&body)
}

func HandleDeleteFile(files.DeleteFileParams) middleware.Responder {
	body := models.BasicResponse{""}
	return files.NewDeleteFileOK().WithPayload(&body)
}
