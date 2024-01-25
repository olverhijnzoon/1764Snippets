import pandas as pd
import numpy as np
from scipy.optimize import curve_fit
import matplotlib.pyplot as plt

# Reading data from a CSV file
data = pd.read_csv('data.csv')
x = data['Year']
y = data['Revenue (in million USD)']

# Normalize the data
x_norm = (x - np.min(x)) / (np.max(x) - np.min(x))
y_norm = (y - np.min(y)) / (np.max(y) - np.min(y))

# Defining the exponential function
def exponential_func(x, a, b):
    return a * np.exp(b * x)

# Fitting the exponential model to the normalized data
params, covariance = curve_fit(exponential_func, x_norm, y_norm)

# Extracting the parameters
a, b = params

# Calculate the residuals and the error
residuals = y_norm - exponential_func(x_norm, a, b)
fitted_error = np.sqrt(np.mean(residuals**2))

# Plotting the normalized data, the fit, and the residuals
plt.figure(figsize=(10, 6))

plt.subplot(2, 1, 1)
plt.scatter(x_norm, y_norm, label='Data')
plt.plot(x_norm, exponential_func(x_norm, a, b), label='Fitted function', color='red')
plt.legend()
plt.xlabel('Year')
plt.ylabel('Revenue (in million USD)')
plt.title('Exponential Fit')
plt.text(0.5, 0.5, f'a = {a:.2f}, b = {b:.2f}, error = {fitted_error:.2f}', horizontalalignment='center', verticalalignment='center', transform=plt.gca().transAxes)

plt.subplot(2, 1, 2)
plt.scatter(x_norm, residuals, label='Residuals', color='orange')
plt.xlabel('Year')
plt.ylabel('Residuals')
plt.title('Residuals')
plt.axhline(0, color='black', linewidth=0.6)
plt.legend()

plt.tight_layout()
plt.show()

print(f"Fitted parameters: a = {a}, b = {b}")
