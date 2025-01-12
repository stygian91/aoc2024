#!/bin/bash

if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <day-number> <part-number> <input-file-path>"
  exit 1
fi

day=$(printf %02d $1)
part=$2
input_path=$3

build_path="bin/q$day"
bin_path="$build_path/part$part"

echo "Running: mkdir -p $build_path"
echo '---------------------------'
mkdir -p $build_path

echo "Running: odin run "./src/q$day/part$part" "-out:$bin_path" -- $input_path"
echo '---------------------------'
odin run "./src/q$day/part$part" "-out:$bin_path" -- $input_path
