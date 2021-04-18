@echo off
SET GOOS=linux
SET GOARCH=arm64
SET CGO_ENABLED=0
REM BUILDINFO_FLAGS !? blank

go build -ldflags "-extldflags "-static"" -trimpath -o cybervector-proxy cmd/main.go

echo Done.
pause > nul