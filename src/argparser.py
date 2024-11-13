import argparse


def parse_arguments() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=" Yajto is a tool for converting data between multiple structured file formats like YAML and JSON. For now, it only supports these two extensions.")

    parser.add_argument(
        'input', 
        help="Path to the input file."
    )

    parser.add_argument(
        '-o',
        '--output',
        help='Output file path or name. If omitted, the output will use the input file\'s name.'
    )

    return parser.parse_args()