#!/bin/sh

for file in gitlander-*; do
  cat $file | sed 's/\r//' > temp.tevs
  cat temp.tevs > $file
done

rm temp.tevs
