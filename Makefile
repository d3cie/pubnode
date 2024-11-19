dev:
	air --build.cmd "go build -o bin/web.exe ./cmd/web" --build.bin ".\bin\web.exe"

PHONY: dev/tw
dev/tw:
	pnpm tailwind:watch


