# install.ps1 - Download yi binary + skill files from GitHub Release
#Requires -Version 3.0
[CmdletBinding()]
param(
    [string]$Version = "latest"
)

$ErrorActionPreference = "Stop"

$Repo = "godcong/yi"
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$SkillDir = Split-Path -Parent $ScriptDir
$InstallDir = Join-Path $SkillDir "bin"

# Detect architecture
$Arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "amd64" }
    "x86"   { "386" }
    "ARM64" { "arm64" }
    default { "amd64" }
}

Write-Host "Platform: windows/$Arch"

$FileName = "yi-windows-$Arch.zip"
$DownloadOk = $false

function Try-Download {
    param([string]$Ver)

    if ($Ver -eq "latest") {
        Write-Host "Resolving latest version..."
        try {
            $Release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -UseBasicParsing
            $Ver = $Release.tag_name
        } catch {
            Write-Host "Failed to resolve latest version"
            return $null
        }
    }

    Write-Host "Version: $Ver"

    $DownloadUrl = "https://github.com/$Repo/releases/download/$Ver/$FileName"
    $ChecksumsUrl = "https://github.com/$Repo/releases/download/$Ver/checksums.txt"

    Write-Host "=== Downloading binary ==="
    Write-Host "URL: $DownloadUrl"

    if (-not (Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    }

    $TmpFile = Join-Path $env:TEMP "yi-install-$([System.IO.Path]::GetRandomFileName()).zip"

    try {
        Invoke-WebRequest -Uri $DownloadUrl -OutFile $TmpFile -UseBasicParsing
    } catch {
        Write-Host "Download failed: $_"
        if (Test-Path $TmpFile) { Remove-Item $TmpFile -Force }
        return $null
    }

    # Verify checksum
    $ChkFile = Join-Path $env:TEMP "yi-checksums-$([System.IO.Path]::GetRandomFileName()).txt"
    try {
        Invoke-WebRequest -Uri $ChecksumsUrl -OutFile $ChkFile -UseBasicParsing -ErrorAction SilentlyContinue
        if (Test-Path $ChkFile) {
            Write-Host "Verifying checksum..."
            $Checksums = Get-Content $ChkFile
            $ExpectedLine = $Checksums | Where-Object { $_ -match $FileName }
            if ($ExpectedLine) {
                $Expected = ($ExpectedLine -split '\s+')[0]
                $Actual = (Get-FileHash -Path $TmpFile -Algorithm SHA256).Hash
                if ($Expected -ne $Actual) {
                    Write-Host "Checksum mismatch! Expected $Expected, got $Actual"
                    Remove-Item $TmpFile -Force -ErrorAction SilentlyContinue
                    Remove-Item $ChkFile -Force -ErrorAction SilentlyContinue
                    return $null
                }
                Write-Host "Checksum OK"
            }
        }
    } catch {
        # Checksum verification is optional, continue if it fails
    } finally {
        if (Test-Path $ChkFile) { Remove-Item $ChkFile -Force -ErrorAction SilentlyContinue }
    }

    # Extract
    try {
        Expand-Archive -Path $TmpFile -DestinationPath $InstallDir -Force
    } catch {
        Write-Host "Extraction failed: $_"
        Remove-Item $TmpFile -Force -ErrorAction SilentlyContinue
        return $null
    } finally {
        if (Test-Path $TmpFile) { Remove-Item $TmpFile -Force -ErrorAction SilentlyContinue }
    }

    # Check for binary
    $BinaryPath = Join-Path $InstallDir "yi.exe"
    if (Test-Path $BinaryPath) {
        Write-Host "Binary installed: $BinaryPath"
        return $Ver
    }

    # Check nested directory
    $NestedPath = Join-Path $InstallDir "yi-windows-$Arch\yi.exe"
    if (Test-Path $NestedPath) {
        Move-Item -Path $NestedPath -Destination $BinaryPath -Force
        Write-Host "Binary installed: $BinaryPath"
        return $Ver
    }

    Write-Host "Warning: yi.exe not found after extraction"
    return $null
}

function Try-GoInstall {
    $GoCmd = Get-Command go -ErrorAction SilentlyContinue
    if (-not $GoCmd) {
        Write-Host "Go not found, skipping go install fallback"
        return $false
    }

    Write-Host "=== Building from source via Go ==="
    $env:GOBIN = $InstallDir

    try {
        & go install "github.com/$Repo/cmd/divine@latest"
    } catch {
        Write-Host "Go install failed: $_"
        return $false
    }

    $DivinePath = Join-Path $InstallDir "divine.exe"
    $BinaryPath = Join-Path $InstallDir "yi.exe"

    if (Test-Path $DivinePath) {
        Move-Item -Path $DivinePath -Destination $BinaryPath -Force
        Write-Host "Binary built and installed: $BinaryPath"
        return $true
    }

    Write-Host "Go build did not produce expected binary"
    return $false
}

# Main
$InstalledVersion = Try-Download -Ver $Version

if ($InstalledVersion) {
    Write-Host ""
    Write-Host "=== Done ==="
    Write-Host "Version: $InstalledVersion"
    $DownloadOk = $true
} elseif (Try-GoInstall) {
    Write-Host ""
    Write-Host "=== Done (built from source) ==="
    $DownloadOk = $false
} else {
    Write-Host ""
    Write-Host "=== Failed ==="
    Write-Host "Could not download or build yi binary."
    Write-Host "Please install manually:"
    Write-Host "  go install github.com/godcong/yi/cmd/divine@latest"
    exit 1
}

# Update skill files
Write-Host ""
Write-Host "=== Updating skill files ==="
$SkillFilesUrl = "https://github.com/$Repo/releases/download/$InstalledVersion"

$SkillFiles = @(
    "SKILL.md",
    "references\data-format.md",
    "scripts\install.sh",
    "scripts\install.ps1",
    "scripts\install.bat"
)

foreach ($File in $SkillFiles) {
    $Dest = Join-Path $SkillDir $File
    $DestDir = Split-Path -Parent $Dest
    if (-not (Test-Path $DestDir)) {
        New-Item -ItemType Directory -Path $DestDir -Force | Out-Null
    }
    Write-Host -NoNewline "  $File ... "
    try {
        Invoke-WebRequest -Uri "$SkillFilesUrl/skill/$File" -OutFile $Dest -UseBasicParsing -ErrorAction Stop
        Write-Host "OK"
    } catch {
        Write-Host "SKIP (using local copy)"
    }
}

$BinaryPath = Join-Path $InstallDir "yi.exe"
Write-Host ""
Write-Host "Binary:  $BinaryPath"
if (Test-Path $BinaryPath) {
    & $BinaryPath --version 2>$null
}
