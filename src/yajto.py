import os
import pathlib
import json
import yaml

from argparser import parse_arguments


allowed_extension = ['.json', '.yaml', '.yml']

def get_extension_file(input_file: str) -> str:
    extension = pathlib.Path(input_file).suffix
    if extension not in allowed_extension:
        raise NotImplementedError('The program does not handle this file extension yet.')
    return extension

def convert_file(input_file: str, output_file: str):
    extension = get_extension_file(input_file)
    
    match extension:
        case '.json':
            # Convert JSON to YAML
            with open(input_file, "r") as input_fp, open(output_file, "w") as output_fp:
                json_structure = json.load(input_fp)
                yaml_structure = yaml.dump(json_structure, sort_keys=False, default_flow_style=False)
                output_fp.write(yaml_structure)
        case '.yaml' | '.yml':
            # Convert YAML to JSON
            with open(input_file, "r") as input_fp, open(output_file, "w") as output_fp:
                yaml_structure = yaml.safe_load(input_fp)
                json_structure = json.dumps(yaml_structure, indent=4)
                output_fp.write(json_structure)


def main():
    args = parse_arguments()
    input_file = args.input
    if args.output:
        output_file = args.output
    else:
        output_file = pathlib.Path(input_file).stem
    convert_file(input_file, output_file)


if __name__ == '__main__':
    main()