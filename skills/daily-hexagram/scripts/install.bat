@echo off
REM install.bat - Wrapper for install.ps1
REM This script simply calls the PowerShell install script.
REM If execution policy blocks .ps1, it temporarily bypasses it.

setlocal

set SCRIPT_DIR=%~dp0
set PS_SCRIPT=%SCRIPT_DIR%install.ps1

if not exist "%PS_SCRIPT%" (
    echo Error: install.ps1 not found in %SCRIPT_DIR%
    exit /b 1
)

REM Pass all arguments to the PowerShell script
powershell -ExecutionPolicy Bypass -File "%PS_SCRIPT%" %*

exit /b %ERRORLEVEL%
