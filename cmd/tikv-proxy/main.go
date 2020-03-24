package main

import (
	"net"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/xiaogaozi/tikv-proxy/pkg/server"
	pb "github.com/xiaogaozi/tikv-proxy/pkg/serverpb"
)

const (
	defaultPort    = "7788"
	defaultPdAddrs = "localhost:2379"
)

var rootCmd = &cobra.Command{
	Use:   "tikv-proxy",
	Short: "A proxy for TiKV",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	// Flags
	rootCmd.PersistentFlags().String("port", defaultPort, "server port number")
	rootCmd.PersistentFlags().String("pd-addrs", defaultPdAddrs, "comma separated placement driver addresses")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug mode")
	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.AutomaticEnv()
}

func startServer() {
	// Logging
	var logger *log.Logger
	if viper.GetBool("debug") {
		logger, _ = log.NewDevelopment()
	} else {
		logger, _ = log.NewProduction()
	}
	log.ReplaceGlobals(logger)

	log.S().Infof("Welcome to TiKV proxy.")

	// Start server
	port := viper.GetString("port")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.S().Fatalf("Failed to listen: %v", err)
	}
	pdAddrs := strings.Split(viper.GetString("pd-addrs"), ",")
	s := grpc.NewServer()
	pb.RegisterTikvProxyServer(s, server.NewServer(pdAddrs))
	log.S().Infof("Server is listening on 0.0.0.0:%s", port)
	if err := s.Serve(lis); err != nil {
		log.S().Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	rootCmd.Execute()
}
