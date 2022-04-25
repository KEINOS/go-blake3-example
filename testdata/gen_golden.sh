#!/bin/sh
# =============================================================================
#  Golden Data Generator
# =============================================================================
#  This script generates the golden data for the test cases via the Docker
#  container which includes the official Blake3 written in Rust.

tr -dc 'a-zA-Z0-9' </dev/urandom | fold -w 32 | head -n 32 | sort | uniq >list1

# -----------------------------------------------------------------------------
#  Generate 32 byte golden data
# -----------------------------------------------------------------------------
echo "{" >list2

while IFS= read -r i; do
    origin="$i"
    hash=$(echo -n "$origin" | b3sum --no-names -)
    printf "\"%s\": \"%s\",\n" "$origin" "$hash" >>list2
done <list1

origin="This is a string"
hash="$(echo -n "$origin" | b3sum --no-names -)"
printf "\"%s\": \"%s\"\n" "$origin" "$hash" >>list2

echo "}" >>list2

jq . <list2 >golden_data32.json

# -----------------------------------------------------------------------------
#  Generate 64 byte golden data
# -----------------------------------------------------------------------------
echo "{" >list2

while IFS= read -r i; do
    origin="$i"
    hash=$(echo -n "$origin" | b3sum --length 64 --no-names -)
    printf "\"%s\": \"%s\",\n" "$origin" "$hash" >>list2
done <list1

origin="This is a string"
hash="$(echo -n "$origin" | b3sum --length 64 --no-names -)"
printf "\"%s\": \"%s\"\n" "$origin" "$hash" >>list2

echo "}" >>list2

jq . <list2 >golden_data64.json

rm -rf list*
