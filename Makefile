dev:
	air --build.cmd "go build -o bin/web.exe ./cmd/web" --build.bin ".\bin\web.exe" --build.exclude_dir "bin, tmp, node_modules, dist, data"

PHONY: dev/tw
dev/tw:
	pnpm tailwind:watch


