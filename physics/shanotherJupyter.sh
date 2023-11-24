#!/bin/bash

if ! command -v jupyter-lab &> /dev/null; then
    echo "Installing JupyterLab"
    pip3 install jupyterlab
else
    echo "JupyterLab is already installed"
fi

# for the simpleplot
pip3 install matplotlib

# scientific
pip3 install scipy

book_dir="dirac_delta"
mkdir -p "$book_dir"

# call py script to create a notebook
python3 pyanotherJupyter.py "$book_dir/$book_dir.ipynb"

cd "$book_dir"

# Create Makefile
cat <<EOF >Makefile
BOOK=dirac_delta
run:
	jupyter-lab \$(BOOK).ipynb
EOF

# launch JupyterLab with it
echo "Launching JupyterLab..."
make
