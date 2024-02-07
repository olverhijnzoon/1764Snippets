#!/bin/bash

# for the simpleplot
pip3 install matplotlib

# scientific
pip3 install scipy

book_dir="anotherJupyterBook"
mkdir -p "$book_dir"

# call py script to create a notebook
python3 pyanotherJupyter.py "$book_dir/$book_dir.ipynb"

cd "$book_dir"

# Create Makefile
cat <<EOF >Makefile
BOOK=$book_dir
run:
	jupyter-lab \$(BOOK).ipynb

prep:
	pip3 install jupyterlab
	jupyter lab build

EOF

# launch JupyterLab with it
echo "Launching JupyterLab..."
make
