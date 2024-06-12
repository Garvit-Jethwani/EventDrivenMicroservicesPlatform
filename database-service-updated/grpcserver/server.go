
package grpcserver

import (
    "context"
    "database-service/database"
    "database-service/proto"
    "database/sql"
)

type Server struct {
    proto.UnimplementedDatabaseServiceServer
    DB *sql.DB
}

func (s *Server) Query(ctx context.Context, req *proto.QueryRequest) (*proto.QueryResponse, error) {
    rows, err := database.Query(s.DB, req.Sql, req.Params...)
    if err != nil {
        return nil, err
    }

    response := &proto.QueryResponse{}
    for _, row := range rows {
        rowMap := make(map[string]string)
        for key, value := range row {
            rowMap[key] = value.(string)
        }
        response.Rows = append(response.Rows, rowMap)
    }
    return response, nil
}

func (s *Server) Command(ctx context.Context, req *proto.CommandRequest) (*proto.CommandResponse, error) {
    rowsAffected, err := database.Execute(s.DB, req.Sql, req.Params...)
    if err != nil {
        return nil, err
    }
    return &proto.CommandResponse{RowsAffected: rowsAffected}, nil
}
