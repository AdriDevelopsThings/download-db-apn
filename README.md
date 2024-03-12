# download-db-apn
A script to download railway apn sketches for all stations from trassenfinder.de

## Installation dependencies
You need python version 3.5 or higher. Run `pip install -r requirements.txt` to install all requirements.

## Running the script
Just run `python main.py`.
### Usage
```
usage: download-db-apn [-h] [-i INFRASTRUCTURE_ID] [-o OUTPUT_DIR] [--overwrite-existing]

Script to download railway apn sketches for all stations from trassenfinder.de

options:
  -h, --help            show this help message and exit
  -i INFRASTRUCTURE_ID, --infrastructure-id INFRASTRUCTURE_ID
                        ID of the infrastructure otherwise the the first infrastructure found will be used
  -o OUTPUT_DIR, --output-dir OUTPUT_DIR
                        Directory where the apn sketches will be saved.
  --overwrite-existing  Overwrite already existing files instead of skipping these files.
```