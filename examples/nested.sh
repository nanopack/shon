#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
shon="${dir}/../shon"
payload='{"one": {"a": "ay"}, "two": {"b": "bee"}}'

# turn JSON payload into local variables
eval $(echo $payload | $shon)

# extract values
echo "one -> ${one_a_value}"
echo "two -> ${two_b_value}"