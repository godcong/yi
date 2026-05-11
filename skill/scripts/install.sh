#!/usr/bin/env bash
# install.sh - Download yi binary + skill files from GitHub Release
set -euo pipefail

REPO="godcong/yi"
VERSION="${1:-latest}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SKILL_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
INSTALL_DIR="${SKILL_DIR}/bin"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$OS" in
  darwin*) OS="darwin" ;;
  linux*)  OS="linux" ;;
  mingw*|msys*|cygwin*|windows_nt*) OS="windows" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  i386|i686)     ARCH="386" ;;
  *) echo "Unsupported ARCH: $ARCH"; exit 1 ;;
esac

echo "Platform: ${OS}/${ARCH}"

if [ "$VERSION" = "latest" ]; then
  VERSION=$(curl -sfL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
  if [ -z "$VERSION" ]; then
    echo "Failed to resolve latest version"
    exit 1
  fi
fi

echo "Version: ${VERSION}"

EXT=""
ARCHIVE_EXT="tar.gz"
if [ "$OS" = "windows" ]; then
  EXT=".exe"
  ARCHIVE_EXT="zip"
fi

FILENAME="yi-${OS}-${ARCH}.${ARCHIVE_EXT}"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"
CHECKSUMS_URL="https://github.com/${REPO}/releases/download/${VERSION}/checksums.txt"

echo "=== Downloading binary ==="
echo "URL: ${DOWNLOAD_URL}"

mkdir -p "$INSTALL_DIR"
TMPDIR=$(mktemp -d)
trap "rm -rf $TMPDIR" EXIT

curl -sfL "$DOWNLOAD_URL" -o "$TMPDIR/archive.${ARCHIVE_EXT}"
curl -sfL "$CHECKSUMS_URL" -o "$TMPDIR/checksums.txt" 2>/dev/null || true

# Verify checksum
if [ -s "$TMPDIR/checksums.txt" ]; then
  echo "Verifying checksum..."
  EXPECTED=$(grep "$FILENAME" "$TMPDIR/checksums.txt" | awk '{print $1}')
  if [ -n "$EXPECTED" ]; then
    ACTUAL=$(sha256sum "$TMPDIR/archive.${ARCHIVE_EXT}" | awk '{print $1}')
    if [ "$EXPECTED" != "$ACTUAL" ]; then
      echo "Checksum mismatch! Expected $EXPECTED, got $ACTUAL"
      exit 1
    fi
    echo "Checksum OK"
  fi
fi

if [ "$ARCHIVE_EXT" = "zip" ]; then
  unzip -o "$TMPDIR/archive.zip" -d "$TMPDIR/out"
else
  tar xzf "$TMPDIR/archive.tar.gz" -C "$TMPDIR"
fi

BINARY_PATH=$(find "$TMPDIR" -name "yi${EXT}" -type f | head -1)
if [ -z "$BINARY_PATH" ]; then
  echo "Binary not found in archive"
  exit 1
fi

TARGET="${INSTALL_DIR}/yi${EXT}"
mv "$BINARY_PATH" "$TARGET"
chmod +x "$TARGET"
echo "Binary installed: ${TARGET}"

# Download skill files from release
echo ""
echo "=== Updating skill files ==="
SKILL_FILES_URL="https://github.com/${REPO}/releases/download/${VERSION}"

for f in SKILL.md references/data-format.md scripts/install.sh scripts/install.bat; do
  DEST="${SKILL_DIR}/${f}"
  DEST_DIR=$(dirname "$DEST")
  mkdir -p "$DEST_DIR"
  echo -n "  ${f} ... "
  HTTP_CODE=$(curl -sfL -w "%{http_code}" -o "$DEST" "${SKILL_FILES_URL}/skill/${f}" 2>/dev/null || echo "000")
  if [ "$HTTP_CODE" = "200" ]; then
    echo "OK"
  else
    echo "SKIP (HTTP ${HTTP_CODE}, using local copy)"
  fi
done

echo ""
echo "=== Done ==="
echo "Version: ${VERSION}"
echo "Binary:  ${TARGET}"
"$TARGET" --version 2>/dev/null || true
