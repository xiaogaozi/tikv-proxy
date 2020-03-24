package server

import (
	"context"

	tidbconfig "github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"

	pb "github.com/xiaogaozi/tikv-proxy/pkg/serverpb"
)

type Server struct {
	pb.UnimplementedTikvProxyServer
	tikvClient *tikv.RawKVClient
}

func NewServer(pdAddrs []string) *Server {
	tikvClient, err := tikv.NewRawKVClient(pdAddrs, tidbconfig.Security{})
	if err != nil {
		panic(err)
	}
	return &Server{
		tikvClient: tikvClient,
	}
}

func (s *Server) RawGet(ctx context.Context, req *pb.RawGetRequest) (*pb.RawGetResponse, error) {
	res := &pb.RawGetResponse{}
	if val, err := s.tikvClient.Get(req.GetKey()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		res.Value = val
		return res, nil
	}
}

func (s *Server) RawBatchGet(ctx context.Context, req *pb.RawBatchGetRequest) (*pb.RawBatchGetResponse, error) {
	res := &pb.RawBatchGetResponse{}
	if values, err := s.tikvClient.BatchGet(req.GetKeys()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		res.Values = values
		return res, nil
	}
}

func (s *Server) RawPut(ctx context.Context, req *pb.RawPutRequest) (*pb.RawPutResponse, error) {
	res := &pb.RawPutResponse{}
	if err := s.tikvClient.Put(req.GetKey(), req.GetValue()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		return res, nil
	}
}

func (s *Server) RawBatchPut(ctx context.Context, req *pb.RawBatchPutRequest) (*pb.RawBatchPutResponse, error) {
	res := &pb.RawBatchPutResponse{}
	if err := s.tikvClient.BatchPut(req.GetKeys(), req.GetValues()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		return res, nil
	}
}

func (s *Server) RawDelete(ctx context.Context, req *pb.RawDeleteRequest) (*pb.RawDeleteResponse, error) {
	res := &pb.RawDeleteResponse{}
	if err := s.tikvClient.Delete(req.Key); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		return res, nil
	}
}

func (s *Server) RawBatchDelete(ctx context.Context, req *pb.RawBatchDeleteRequest) (*pb.RawBatchDeleteResponse, error) {
	res := &pb.RawBatchDeleteResponse{}
	if err := s.tikvClient.BatchDelete(req.GetKeys()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		return res, nil
	}
}

func (s *Server) RawDeleteRange(ctx context.Context, req *pb.RawDeleteRangeRequest) (*pb.RawDeleteRangeResponse, error) {
	res := &pb.RawDeleteRangeResponse{}
	if err := s.tikvClient.DeleteRange(req.GetStartKey(), req.GetEndKey()); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		return res, nil
	}
}

func (s *Server) RawScan(ctx context.Context, req *pb.RawScanRequest) (*pb.RawScanResponse, error) {
	res := &pb.RawScanResponse{}
	if keys, values, err := s.tikvClient.Scan(req.GetStartKey(), req.GetEndKey(), int(req.GetLimit())); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		res.Keys = keys
		res.Values = values
		return res, nil
	}
}

func (s *Server) RawReverseScan(ctx context.Context, req *pb.RawReverseScanRequest) (*pb.RawReverseScanResponse, error) {
	res := &pb.RawReverseScanResponse{}
	if keys, values, err := s.tikvClient.ReverseScan(req.GetStartKey(), req.GetEndKey(), int(req.GetLimit())); err != nil {
		res.Error = err.Error()
		return res, err
	} else {
		res.Keys = keys
		res.Values = values
		return res, nil
	}
}
