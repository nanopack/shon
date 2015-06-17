#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
shon="${dir}/../shon"
payload='{"one": "1", "two": "2"}'
prefix="PRE_"

# turn JSON payload into local variables with a prefix
eval $(echo $payload | $shon | sed -e "s/^/${prefix}/")

# extract values
echo "one -> ${PRE_one_value}"
echo "two -> ${PRE_two_value}"