#!/bin/bash
GOAL=1764
DIR="$(cd "$(dirname "../../${BASH_SOURCE[0]}")" &>/dev/null && pwd)"

total_subfolders=0

for d in "$DIR"/*/ ; do
    if [ -d "$d" ]; then
        count=$(find "$d" -mindepth 1 -maxdepth 1 -type d | wc -l)
        total_subfolders=$((total_subfolders + count))
    fi
done

echo "Snippets: $total_subfolders/$GOAL"
progress=$(echo "scale=2; $total_subfolders/$GOAL*100" | bc)
progress=${progress%.*} # Convert to integer
progress_bar=""
for ((i=0; i<progress; i+=2)); do
    progress_bar+="="
done
echo -ne "Progress: [${progress_bar:0:50}] $progress%\r"
echo ""