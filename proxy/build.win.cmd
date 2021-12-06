@echo off
SET GOOS=windows
SET GOARCH=amd64
SET CGO_ENABLED=0
REM BUILDINFO_FLAGS !? blank

go build -ldflags "-extldflags "-static"" -trimpath -o cybervector-proxy cmd/main.go

echo Done.
pause > nul