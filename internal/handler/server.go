package handler

import (
	"context"

	pb "github.com/devlogger/kpradipta/api/proto/logger"
	"github.com/devlogger/kpradipta/internal/db"
	"github.com/devlogger/kpradipta/internal/metrics"
)

type LogServer struct {
	pb.UnimplementedLogServiceServer
	// Add your DB/Redis clients here
}

func (s *LogServer) SendLog(ctx context.Context, entry *pb.LogEntry) (*pb.Empty, error) {
	metrics.LogsReceived.Inc()
	_, err := db.Pool.Exec(ctx,
		`INSERT INTO logs (service, level, message, timestamp)
		 VALUES ($1, $2, $3, $4)`,
		entry.Service, entry.Level, entry.Message, entry.Timestamp,
	)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *LogServer) GetLogs(ctx context.Context, req *pb.QueryRequest) (*pb.LogList, error) {
	// Query PostgreSQL or Redis and return log entries
	metrics.LogsQueried.Inc()
	rows, err := db.Pool.Query(ctx,
		`SELECT service, level, message, timestamp
         FROM logs
         WHERE service = $1 AND level = $2`,
		req.Service, req.Level,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*pb.LogEntry
	for rows.Next() {
		var e pb.LogEntry
		err := rows.Scan(&e.Service, &e.Level, &e.Message, &e.Timestamp)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &e)
	}

	return &pb.LogList{Entries: entries}, nil
}
