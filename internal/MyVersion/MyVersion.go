package MyVersion

import (
	"context"
	"fmt"

	"github.com/apavanello/gRPC-Calculato/pkg/calculatorpb"
	"github.com/apavanello/gRPC-Calculato/versioninfo"
)

type Server struct{}

func (s *Server) GetVersion(ctx context.Context, req *calculatorpb.VersionRequest) (res *calculatorpb.VersionResponse, err error) {

	fmt.Println("start Version Service")
	res = &calculatorpb.VersionResponse{Version: &calculatorpb.Version{}}

	res.Version.ProjectName = versioninfo.ProjectName
	res.Version.Revision = versioninfo.Revision
	res.Version.Version = versioninfo.Version

	err = nil
	return
}
