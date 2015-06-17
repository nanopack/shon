#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
shon="${dir}/../shon"
payload='{"one": "1", "two": "2"}'

# turn JSON payload into local variables
eval $(echo $payload | $shon)

# extract values
echo "one -> ${one_value}"
echo "two -> ${two_value}"