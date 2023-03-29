package handler

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/proto/foo"
	"io"
	"strconv"
	"time"
)

// SaveStreamHttpLog is method
func (c *Handler) SaveStreamHttpLog(stream foo.LogService_SaveStreamHttpLogServer) error {
	var i int32
	startTime := time.Now()
	for {
		_, err := stream.Recv()

		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&foo.HttpLogStreamResponse{
				Total:    i,
				Duration: int32(endTime.Sub(startTime).Seconds()),
			})
		}

		i++

		if err != nil {
			return err
		}
	}
}

// SaveHttpLog is a function
func (c *Handler) SaveHttpLog(ctx context.Context, r *foo.SaveHttpLogRequest) (*foo.HttpLog, error) {
	lg := entity.HttpLog{
		Ip:     r.GetIp(),
		Path:   r.GetPath(),
		Method: r.GetMethod(),
	}

	err := c.dep.repo.HttpLog.Save(ctx, &lg)
	if err != nil {
		return nil, err
	}

	return &foo.HttpLog{
		Id:        int64(lg.ID),
		Ip:        lg.Ip,
		Path:      lg.Path,
		Method:    lg.Method,
		CreatedAt: lg.CreatedAt.String(),
		UpdatedAt: lg.UpdatedAt.String(),
		DeletedAt: lg.DeletedAt.Time.String(),
	}, nil
}

// FindHttpLog is a function to retrieve single data
func (c *Handler) FindHttpLog(ctx context.Context, r *foo.FindHttpLogRequest) (*foo.HttpLog, error) {
	id, errC := strconv.ParseInt(r.GetId(), 10, 64)
	if errC != nil {
		return nil, errC
	}

	lg, err := c.dep.repo.HttpLog.Find(ctx, int(id))
	if err != nil {
		return nil, err
	}

	return &foo.HttpLog{
		Id:        int64(lg.ID),
		Ip:        lg.Ip,
		Path:      lg.Path,
		Method:    lg.Method,
		CreatedAt: lg.CreatedAt.String(),
		UpdatedAt: lg.UpdatedAt.String(),
		DeletedAt: lg.DeletedAt.Time.String(),
	}, nil
}

// GetHttpLog is a function to retrieve list of data
func (c *Handler) GetHttpLog(ctx context.Context, _ *foo.GetHttpLogRequest) (*foo.HttpLogs, error) {
	r, err := c.dep.repo.HttpLog.Get(ctx)
	if err != nil {
		return nil, err
	}

	var logs []*foo.HttpLog
	for _, lg := range r {
		logs = append(logs, &foo.HttpLog{
			Id:        int64(lg.ID),
			Ip:        lg.Ip,
			Path:      lg.Path,
			Method:    lg.Method,
			CreatedAt: lg.CreatedAt.String(),
			UpdatedAt: lg.UpdatedAt.String(),
			DeletedAt: lg.DeletedAt.Time.String(),
		})
	}

	return &foo.HttpLogs{
		Logs: logs,
	}, nil
}
