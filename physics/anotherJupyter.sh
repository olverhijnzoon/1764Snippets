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

book_dir="anotherJupyterBook"
mkdir -p "$book_dir"

# call py script to create a notebook
python3 anotherJupyter.py "$book_dir/basic_plot.ipynb"

# launch JupyterLab with it
echo "Launching JupyterLab..."
jupyter-lab "$book_dir"
