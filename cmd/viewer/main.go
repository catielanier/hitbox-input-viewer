package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/catielanier/hitbox-input-viewer/internal/controller"
	"github.com/catielanier/hitbox-input-viewer/internal/hub"
)

// Build the Svelte frontend first.
//
//go:embed webdist/*
var frontend embed.FS

func main() {
	listen := flag.String(
		"listen",
		":8080",
		"LAN HTTP address",
	)

	index := flag.Uint(
		"controller",
		0,
		"XInput controller slot, 0 through 3",
	)

	flag.Parse()

	if *index > 3 {
		log.Fatal("controller index must be 0..3")
	}

	reader, err := controller.NewReader(*index)
	if err != nil {
		log.Fatal(err)
	}

	h := hub.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", h.ServeWS)

	assets, err := fs.Sub(frontend, "webdist")
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", http.FileServer(http.FS(assets)))

	go func() {
		ticker := time.NewTicker(4 * time.Millisecond)
		defer ticker.Stop()

		last := controller.Empty()

		for range ticker.C {
			state := reader.Read()

			if !controller.Equal(state, last) {
				h.Publish(state)
				last = state
			}
		}
	}()

	fmt.Printf(
		"Controller viewer serving http://localhost%s/\n",
		*listen,
	)

	log.Fatal(http.ListenAndServe(*listen, mux))
}
