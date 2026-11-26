.PHONY: web windows test

web:
	cd web && npm install && npm run build
	cp -r web/dist/. cmd/viewer/webdist/

windows:
	GOOS=windows GOARCH=amd64 go build -o hitbox-input-viewer.exe ./cmd/viewer

test:
	go test ./...
