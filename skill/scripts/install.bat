@echo off
REM install.bat - Download yi binary + skill files from GitHub Release
setlocal enabledelayedexpansion

set REPO=godcong/yi
set VERSION=%1
if "%VERSION%"=="" set VERSION=latest
set SCRIPT_DIR=%~dp0
set SKILL_DIR=%SCRIPT_DIR%..
set INSTALL_DIR=%SKILL_DIR%\bin

if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
    set ARCH=amd64
) else if "%PROCESSOR_ARCHITECTURE%"=="x86" (
    set ARCH=386
) else (
    set ARCH=amd64
)

echo Platform: windows/!ARCH!

if "!VERSION!"=="latest" (
    echo Resolving latest version...
    for /f "tokens=2 delims=:," %%a in ('curl -sfL "https://api.github.com/repos/!REPO!/releases/latest" ^| findstr "tag_name"') do (
        set VERSION=%%~a
        set VERSION=!VERSION: =!
        set VERSION=!VERSION:"=!
    )
    if "!VERSION!"=="" (
        echo Failed to resolve latest version
        exit /b 1
    )
)

echo Version: !VERSION!

set FILENAME=yi-windows-!ARCH!.zip
set DOWNLOAD_URL=https://github.com/!REPO!/releases/download/!VERSION!/!FILENAME!
set CHECKSUMS_URL=https://github.com/!REPO!/releases/download/!VERSION!/checksums.txt

echo === Downloading binary ===
echo URL: !DOWNLOAD_URL!

if not exist "!INSTALL_DIR!" mkdir "!INSTALL_DIR!"

set TMPFILE=%TEMP%\yi-install-!RANDOM!.zip
curl -sfL "!DOWNLOAD_URL!" -o "!TMPFILE!"
if errorlevel 1 (
    echo Download failed
    del "!TMPFILE!" 2>nul
    exit /b 1
)

REM Verify checksum
set CHKFILE=%TEMP%\yi-checksums-!RANDOM!.txt
curl -sfL "!CHECKSUMS_URL!" -o "!CHKFILE!" 2>nul
if exist "!CHKFILE!" (
    echo Verifying checksum...
    for /f "tokens=1" %%c in ('findstr "!FILENAME!" "!CHKFILE!"') do set EXPECTED=%%c
    if defined EXPECTED (
        for /f "tokens=1" %%a in ('certutil -hashfile "!TMPFILE!" SHA256 ^| findstr /v ":" ^| findstr /v "CertUtil"') do set ACTUAL=%%a
        set ACTUAL=!ACTUAL: =!
        set EXPECTED=!EXPECTED: =!
        if /i "!ACTUAL!" NEQ "!EXPECTED!" (
            echo Checksum mismatch! Expected !EXPECTED!, got !ACTUAL!
            del "!TMPFILE!" 2>nul
            del "!CHKFILE!" 2>nul
            exit /b 1
        )
        echo Checksum OK
    )
    del "!CHKFILE!" 2>nul
)

powershell -Command "Expand-Archive -Path '!TMPFILE!' -DestinationPath '!INSTALL_DIR!' -Force"
del "!TMPFILE!" 2>nul

if exist "!INSTALL_DIR!\yi.exe" (
    echo Binary installed: !INSTALL_DIR!\yi.exe
) else if exist "!INSTALL_DIR!\yi-windows-!ARCH!\yi.exe" (
    move /y "!INSTALL_DIR!\yi-windows-!ARCH!\yi.exe" "!INSTALL_DIR!\yi.exe" >nul
    echo Binary installed: !INSTALL_DIR!\yi.exe
) else (
    echo Warning: yi.exe not found after extraction
)

echo.
echo === Updating skill files ===
set SKILL_FILES_URL=https://github.com/!REPO!/releases/download/!VERSION!

for %%f in (SKILL.md references\data-format.md scripts\install.sh scripts\install.bat) do (
    set DEST=!SKILL_DIR!\%%f
    set DEST_DIR=!DEST!\..
    if not exist "!DEST_DIR!" mkdir "!DEST_DIR!" 2>nul
    echo|set /p="  %%f ... "
    curl -sfL -o "!DEST!" "!SKILL_FILES_URL!/skill/%%f" 2>nul
    if errorlevel 1 (
        echo SKIP ^(using local copy^)
    ) else (
        echo OK
    )
)

echo.
echo === Done ===
echo Version: !VERSION!
echo Binary:  !INSTALL_DIR!\yi.exe
"!INSTALL_DIR!\yi.exe" --version 2>nul || true

endlocal
