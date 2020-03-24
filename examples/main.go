package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/xiaogaozi/tikv-proxy/pkg/serverpb"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Must provide server address")
	}
	addr := os.Args[1]
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		log.Fatalf("Cannot connect to %s: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewTikvProxyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := []byte("Hello")
	val := []byte("World")

	// RawPut
	if _, err := c.RawPut(ctx, &pb.RawPutRequest{Key: key, Value: val}); err != nil {
		log.Fatalf("Call RawPut failed: %v", err)
	}
	log.Printf("Successfully put: key: %s, value: %s\n", key, val)

	// RawGet
	if value, err := c.RawGet(ctx, &pb.RawGetRequest{Key: key}); err != nil {
		log.Fatalf("Call RawGet failed: %v", err)
	} else {
		log.Printf("Get value of key \"%s\": %s\n", key, value.Value)
	}

	// RawDelete
	if _, err := c.RawDelete(ctx, &pb.RawDeleteRequest{Key: key}); err != nil {
		log.Fatal("Call RawDelete failed: %v", err)
	}
	log.Printf("Successfully delete key \"%s\"", key)

	// RawGet again after delete key
	if value, err := c.RawGet(ctx, &pb.RawGetRequest{Key: key}); err != nil {
		log.Fatalf("Call RawGet failed: %v", err)
	} else {
		log.Printf("Get value of key \"%s\": %s\n", key, value.Value)
	}

	keys := [][]byte{[]byte("k1"), []byte("k2"), []byte("k3")}
	vals := [][]byte{[]byte("v1"), []byte("v2"), []byte("v3")}

	// RawBatchPut
	if _, err := c.RawBatchPut(ctx, &pb.RawBatchPutRequest{Keys: keys, Values: vals}); err != nil {
		log.Fatalf("Call RawBatchPut failed: %v", err)
	}
	log.Printf("Successfully batch put: keys: %s, values: %s", keys, vals)

	// RawBatchGet
	if values, err := c.RawBatchGet(ctx, &pb.RawBatchGetRequest{Keys: keys}); err != nil {
		log.Fatalf("Call RawBatchGet failed: %v", err)
	} else {
		log.Printf("Batch get: keys: %s, values: %s", keys, values.Values)
	}

	// RawBatchDelete
	if _, err := c.RawBatchDelete(ctx, &pb.RawBatchDeleteRequest{Keys: keys}); err != nil {
		log.Fatalf("Call RawBatchDelete failed: %v", err)
	}
	log.Printf("Successfully batch delete: keys: %s", keys)

	// RawBatchGet again after delete keys
	if values, err := c.RawBatchGet(ctx, &pb.RawBatchGetRequest{Keys: keys}); err != nil {
		log.Fatalf("Call RawBatchGet failed: %v", err)
	} else {
		log.Printf("Batch get: keys: %s, values: %s", keys, values.Values)
	}

	// RawBatchPut
	if _, err := c.RawBatchPut(ctx, &pb.RawBatchPutRequest{Keys: keys, Values: vals}); err != nil {
		log.Fatalf("Call RawBatchPut failed: %v", err)
	}
	log.Printf("Successfully batch put: keys: %s, values: %s", keys, vals)

	startKey := []byte("k1")
	endKey := []byte("k4")

	// RawScan
	if res, err := c.RawScan(ctx, &pb.RawScanRequest{StartKey: startKey, EndKey: endKey, Limit: 10}); err != nil {
		log.Fatalf("Call RawScan failed: %v", err)
	} else {
		log.Printf("Scan: keys: %s, values: %s", res.Keys, res.Values)
	}

	// RawReverseScan
	if res, err := c.RawReverseScan(ctx, &pb.RawReverseScanRequest{StartKey: endKey, EndKey: startKey, Limit: 10}); err != nil {
		log.Fatalf("Call RawReverseScan failed: %v", err)
	} else {
		log.Printf("Reverse scan: keys: %s, values: %s", res.Keys, res.Values)
	}

	// RawDeleteRange
	if _, err := c.RawDeleteRange(ctx, &pb.RawDeleteRangeRequest{StartKey: startKey, EndKey: endKey}); err != nil {
		log.Fatalf("Call RawDeleteRange failed: %v", err)
	}
	log.Printf("Successfully delete range: start key: %s, end key: %s", startKey, endKey)

	// RawScan again after delete range
	if res, err := c.RawScan(ctx, &pb.RawScanRequest{StartKey: startKey, EndKey: endKey, Limit: 10}); err != nil {
		log.Fatalf("Call RawScan failed: %v", err)
	} else {
		log.Printf("Scan: keys: %s, values: %s", res.Keys, res.Values)
	}
}
