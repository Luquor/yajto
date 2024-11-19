# Yajto

🐡 Yajto is a tool for converting data between multiple structured file formats like **YAML** and **JSON**.


## What It Does

**Yajto** allows you to:
- Convert JSON files into well-structured YAML files.
- Convert YAML files (including `.yml` extensions) into properly formatted JSON files.

This tool ensures that your data remains consistent and readable during format conversions, preserving the structure and integrity of the input files.


## Why This Project?

In modern development workflows, both YAML and JSON are extensively used for configurations, data exchange, and application settings. Each format has its own strengths:
- **JSON**: Lightweight, widely supported, and ideal for APIs and data interchange.
- **YAML**: Human-readable, flexible, and perfect for configuration files.

However, switching between these formats manually can be time-consuming and error-prone. This tool was created to streamline the process, saving time and reducing errors during file conversions.


## Features

- **Automatic Format Detection**: Detects input file format based on the extension (`.json`, `.yaml`, `.yml`).
- **Preserves Structure**: Maintains data integrity and formatting during conversions.
- **Ease of Use**: Convert files with a simple function call, ensuring minimal setup.


## How to Use

To use the program you can either execute the binary file in the `bin`folder or install the dependencies and run the program with the following :
```python
source .venv/bin/activate
pip install -r requirements.txt
python3 src/yajto.py
```

### Arguments & flags
#### Arguments
- `input`: path to the input file.
#### Optional
- `-h`: help info.
- `-o`: precise the output file.
