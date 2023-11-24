import nbformat as nbf
import sys

# path for the new notebook
notebook_path = sys.argv[1]

# create notebook
nb = nbf.v4.new_notebook()

# add code cell with the plot
code = """
%matplotlib inline
import matplotlib.pyplot as plt
import numpy as np

x = np.linspace(0, 10, 100)
y = np.sin(x)

plt.plot(x, y)
plt.title('Simple Plot')
plt.xlabel('x')
plt.ylabel('sin(x)')
"""

nb['cells'] = [nbf.v4.new_code_cell(code)]

# save the notebook
with open(notebook_path, 'w') as f:
    nbf.write(nb, f)
