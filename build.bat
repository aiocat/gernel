if exist build (
    rmdir /Q /S build
) else (
    mkdir build
)

set GOOS=windows
set GOARCH=386
go build -ldflags "-s -w" .

echo f | xcopy /E /H /C /I /Y .\web .\build\web
echo f | xcopy /E /H /C /I /Y .\ckrnl .\build\ckrnl
echo f | xcopy /F /Y .\gernel.exe .\build\gernel.exe
mkdir .\build\scripts
