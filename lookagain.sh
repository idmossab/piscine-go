#! bin/bash

for file in *.sh
do
	echo "$file" | sed 's/\.sh$//' |  sort -r
done