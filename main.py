#!/usr/bin/env python
# A script to download railway apn sketches for all stations from trassenfinder.de
# Made by: AdriDoesThings <contact@adridoesthings.com>

from argparse import ArgumentParser
from os.path import join, exists
from os import mkdir

from progressbar import progressbar
import requests

parser = ArgumentParser(
    "download-db-apn",
    description="Script to download railway apn sketches for all stations from trassenfinder.de",
)
parser.add_argument(
    "-i",
    "--infrastructure-id",
    type=int,
    help="ID of the infrastructure otherwise the the first infrastructure found will be used",
)
parser.add_argument(
    "-o",
    "--output-dir",
    default="sketches",
    help="Directory where the apn sketches will be saved.",
)
parser.add_argument(
    "--overwrite-existing",
    action="store_true",
    help="Overwrite already existing files instead of skipping these files.",
)


def get_infrastructures():
    """Get a list of all trassenfinder infrastructures."""
    response = requests.get("https://trassenfinder.de/api/web/infrastrukturen")
    response.raise_for_status()
    return response.json()


def get_infrastructure(infrastructure_id: int):
    """Get information about an infrastructure."""
    response = requests.get(
        f"https://trassenfinder.de/api/web/infrastrukturen/{infrastructure_id}"
    )
    response.raise_for_status()
    return response.json()


def get_documents(infrastructure_id: int, ds100: str):
    """Get a list of documents of a station."""
    response = requests.get(
        f"https://trassenfinder.de/api/web/infrastrukturen/{infrastructure_id}/dokumente?ds100={ds100}"
    )
    response.raise_for_status()
    return response.json()["dokumente"]


def download_document(
    infrastructure_id: int, filename: str, output_dir: str, overwrite_existing: bool
):
    """Download a document `filename` to a directory `output_dir`. If the file already exist the download will be skipped if `overwrite_existing` is `False`."""
    output_path = join(output_dir, filename)
    if not overwrite_existing and exists(output_path):
        return
    response = requests.get(
        f"https://trassenfinder.de/api/web/infrastrukturen/{infrastructure_id}/dokumente/{filename}",
        stream=True,
    )
    response.raise_for_status()
    with open(output_path, "wb") as file:
        for chunk in response.iter_content(1024):
            file.write(chunk)


if __name__ == "__main__":
    args = parser.parse_args()
    infrastructure_id = args.infrastructure_id
    output_dir = args.output_dir
    overwrite_existing = args.overwrite_existing

    if not exists(output_dir):
        print("Output directory %s created." % output_dir)
        mkdir(output_dir)

    if not infrastructure_id:
        print("Loading list of infrastructures...")
        infrastructures = get_infrastructures()
        print("Selected infrastructure '%s'." % infrastructures[0]["anzeigename"])
        infrastructure_id = infrastructures[0]["id"]

    print("Loading infrastructure information...")
    infrastructure = get_infrastructure(infrastructure_id)
    stations = infrastructure["ordnungsrahmen"]["betriebsstellen"]

    print("Downloading documents...")
    for i in progressbar(range(len(stations))):
        station = stations[i]
        documents = get_documents(infrastructure_id, station["ds100"])
        for document in documents:
            if "typ" in document and document["typ"] == "apn_skizze":
                download_document(
                    infrastructure_id,
                    document["dateiname"],
                    output_dir,
                    overwrite_existing,
                )
