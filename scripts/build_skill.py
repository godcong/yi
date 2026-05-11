#!/usr/bin/env python3
"""Generate platform-specific .skill files from GoReleaser dist/ binaries.

Usage: python scripts/build_skill.py --dist dist/ --skill-dir skill/ --output dist/
"""
import argparse
import os
import shutil
import tempfile
import zipfile
import tarfile
import glob as glob_mod


def find_binaries(dist_dir):
    """Find all GoReleaser archive files and map them to platform info."""
    archives = []
    for f in sorted(glob_mod.glob(os.path.join(dist_dir, "yi-*.tar.gz")) +
                    glob_mod.glob(os.path.join(dist_dir, "yi-*.zip"))):
        basename = os.path.basename(f)
        # yi-windows-amd64.zip, yi-linux-arm64.tar.gz, etc.
        name = basename.rsplit(".", 1)[0]
        if name.endswith(".tar"):
            name = name[:-4]
        parts = name.split("-", 1)  # "yi", "windows-amd64"
        if len(parts) == 2:
            archives.append((f, parts[1]))
    return archives


def extract_binary(archive_path, platform):
    """Extract the yi binary from a GoReleaser archive, return path or None."""
    tmpdir = tempfile.mkdtemp()
    try:
        if archive_path.endswith(".tar.gz"):
            with tarfile.open(archive_path, "r:gz") as tf:
                tf.extractall(tmpdir)
        elif archive_path.endswith(".zip"):
            with zipfile.ZipFile(archive_path, "r") as zf:
                zf.extractall(tmpdir)

        # Look for yi or yi.exe
        for root, dirs, files in os.walk(tmpdir):
            for name in files:
                if name in ("yi", "yi.exe"):
                    return os.path.join(root, name)
        return None
    except Exception as e:
        print(f"  Warning: failed to extract {archive_path}: {e}")
        return None


def build_skill(platform, binary_path, skill_dir, output_dir):
    """Create a .skill file for given platform."""
    is_windows = "windows" in platform
    skill_name = f"daily-hexagram-{platform}.skill"
    output_path = os.path.join(output_dir, skill_name)

    skill_inner = "daily-hexagram"
    with tempfile.TemporaryDirectory() as td:
        inner_dir = os.path.join(td, skill_inner)
        os.makedirs(inner_dir)

        # Copy SKILL.md
        shutil.copy2(os.path.join(skill_dir, "SKILL.md"), inner_dir)

        # Copy references/
        ref_src = os.path.join(skill_dir, "references")
        ref_dst = os.path.join(inner_dir, "references")
        if os.path.exists(ref_src):
            shutil.copytree(ref_src, ref_dst)

        # Copy scripts/
        scripts_src = os.path.join(skill_dir, "scripts")
        scripts_dst = os.path.join(inner_dir, "scripts")
        if os.path.exists(scripts_src):
            shutil.copytree(scripts_src, scripts_dst)

        # Copy binary as bin/yi
        bin_dir = os.path.join(inner_dir, "bin")
        os.makedirs(bin_dir)
        if is_windows:
            shutil.copy2(binary_path, os.path.join(bin_dir, "yi.exe"))
        else:
            shutil.copy2(binary_path, os.path.join(bin_dir, "yi"))
            # Ensure executable
            os.chmod(os.path.join(bin_dir, "yi"), 0o755)

        # Create zip
        with zipfile.ZipFile(output_path, "w", zipfile.ZIP_DEFLATED) as zf:
            for root, dirs, files in os.walk(td):
                for fname in files:
                    full = os.path.join(root, fname)
                    arcname = os.path.relpath(full, td)
                    zf.write(full, arcname)

    size_kb = os.path.getsize(output_path) / 1024
    print(f"  {skill_name} ({size_kb:.0f} KB)")
    return output_path


def main():
    parser = argparse.ArgumentParser(description="Build platform .skill files")
    parser.add_argument("--dist", default="dist", help="GoReleaser dist/ dir")
    parser.add_argument("--skill-dir", default="skill", help="Skill source dir")
    parser.add_argument("--output", default="dist", help="Output dir for .skill files")
    args = parser.parse_args()

    archives = find_binaries(args.dist)
    if not archives:
        print("No GoReleaser archives found in dist/")
        return

    print(f"Found {len(archives)} platform(s)")

    for archive_path, platform in archives:
        print(f"Building {platform}...")
        binary = extract_binary(archive_path, platform)
        if not binary:
            print(f"  Skipping {platform}: no binary found in archive")
            continue
        build_skill(platform, binary, args.skill_dir, args.output)

    print("Done.")


if __name__ == "__main__":
    main()