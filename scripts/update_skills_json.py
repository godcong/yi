#!/usr/bin/env python3
"""Update skills.json with sha256 checksums for built .skill files.

Usage: python scripts/update_skills_json.py --dist dist/ --skills-json skills.json --version v1.2.1
"""
import argparse
import hashlib
import json
import os


def sha256_file(path):
    h = hashlib.sha256()
    with open(path, "rb") as f:
        for chunk in iter(lambda: f.read(8192), b""):
            h.update(chunk)
    return h.hexdigest()


def update_skills_json(dist_dir, skills_json_path, version):
    with open(skills_json_path, "r", encoding="utf-8") as f:
        data = json.load(f)

    for skill in data.get("skills", []):
        slug = skill.get("slug")
        if slug != "daily-hexagram":
            continue

        # Update generic fallback zip_url sha256
        generic_path = os.path.join(dist_dir, "daily-hexagram.skill")
        if os.path.exists(generic_path):
            skill["sha256"] = sha256_file(generic_path)
            print(f"Updated {slug} sha256 (generic): {skill['sha256']}")

        # Update platform-specific sha256s
        platforms = skill.get("platforms", {})
        for platform_key in platforms:
            filename = f"daily-hexagram-{platform_key}.skill"
            filepath = os.path.join(dist_dir, filename)
            if os.path.exists(filepath):
                checksum = sha256_file(filepath)
                platforms[platform_key]["sha256"] = checksum
                print(f"Updated {slug} {platform_key} sha256: {checksum}")
            else:
                print(f"Warning: {filename} not found in {dist_dir}")

    with open(skills_json_path, "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
        f.write("\n")

    print(f"Updated {skills_json_path}")


def main():
    parser = argparse.ArgumentParser(description="Update skills.json checksums")
    parser.add_argument("--dist", default="dist", help="Directory containing .skill files")
    parser.add_argument("--skills-json", default="skills.json", help="Path to skills.json")
    parser.add_argument("--version", default="", help="Release version (optional, for logging)")
    args = parser.parse_args()

    update_skills_json(args.dist, args.skills_json, args.version)


if __name__ == "__main__":
    main()
