
package main

import (
    "user-management-service/config"
    "user-management-service/httpserver"
    "log"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Start HTTP server
    server := httpserver.NewServer(cfg)
    if err := server.Start(); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
