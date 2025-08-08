package pingkratos

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/orzkratos/pingkratos/pingpong"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type PingService struct {
	pb.UnimplementedPingServer
	slog *log.Helper
}

func NewPingService(logger log.Logger) *PingService {
	slog := log.NewHelper(logger)
	return &PingService{
		slog: slog,
	}
}

func (s *PingService) Ping(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	s.slog.Infof("receive-ping-message: %s", req.GetValue())
	return wrapperspb.String(req.GetValue()), nil
}
