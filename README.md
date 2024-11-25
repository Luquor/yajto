# yajto
🐡 Yajto is a versatile tool designed to convert JSON files into YAML or TOML formats.

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

### What is the yajto-python repository?
This repository, https://github.com/Luquor/yajto-python, represents the first iteration of the yajto tool. In the v1.0.0 release, I outlined plans to create a more efficient and faster version of the tool, given its CLI nature. Rather than overhauling the existing codebase and transitioning from Python to Golang within the same repository, I decided to preserve this repository to maintain accessibility for those who might still find the Python version useful.

## Features

- **Automatic Format Detection**: Detects input file format based on the extension (`.json`, `.yaml`, `.yml`).
- **Preserves Structure**: Maintains data integrity and formatting during conversions.
- **Ease of Use**: Convert files with a simple function call, ensuring minimal setup.


## How to Use

### Arguments & flags
#### Arguments
- `input`: path to the input file.
#### Optional
- `-h`: help info.
- `-o`: precise the output file.