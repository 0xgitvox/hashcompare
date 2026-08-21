"""Minimal example for HashCompare."""

from hashcompare import hashcompare


def main():
 runner = hashcompare({"name": "HashCompare", "dry_run": False})
 result = runner.execute()
 print(result)


if __name__ == "__main__":
 main()