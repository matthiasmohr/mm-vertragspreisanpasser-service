package handler

import (
	"net/http"

	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/labstack/echo/v4"
)

// AppStatus manages status related endpoints.
type AppStatus struct {
	version     string
	buildDate   string
	description string
	commitHash  string
	commitDate  string
	buildBranch string
}

// NewStatus creates a new status handler.
func NewStatus(version, buildDate, description, commitHash, commitDate, buildBranch string) *AppStatus {
	return &AppStatus{
		version:     version,
		buildDate:   buildDate,
		description: description,
		commitHash:  commitHash,
		commitDate:  commitDate,
		buildBranch: buildBranch,
	}
}

// Version provides the application version information.
// @Summary Return the service version.
// @Description Application version information.
// @Tags Service Status
// @Produce json
// @Success 201 {object} dto.AppStatus
// @Router /version [get].
func (s *AppStatus) Version(ctx echo.Context) error {
	status := &dto.AppStatus{
		Version:     s.version,
		BuildDate:   s.buildDate,
		Description: s.description,
		CommitHash:  s.commitHash,
		CommitDate:  s.commitDate,
		BuildBranch: s.buildBranch,
	}

	return ctx.JSON(http.StatusOK, status)
}
