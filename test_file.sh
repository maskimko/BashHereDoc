#!/bin/bash
echo "Example of vim command"
cat <<x23LimitStringx23
i
This is line 1 of the example file.
This is line 2 of the example file.
^[
ZZ
x23LimitStringx23

echo "just echo heredoc"
cat <<TKN
some text
TKN

exit 0
