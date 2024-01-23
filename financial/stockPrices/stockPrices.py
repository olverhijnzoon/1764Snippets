import yfinance as yf
from datetime import datetime

# Define the stock symbol and the time period
stock_symbol = 'TSLA'
start_date = datetime(2010, 6, 29)
end_date = datetime(2024, 1, 23)

# Fetch historical stock prices
stock_data = yf.download(stock_symbol, start=start_date, end=end_date)

# Calculate the daily returns
stock_data['Daily_Return'] = stock_data['Adj Close'].pct_change()

# Calculate the cumulative returns
stock_data['Cumulative_Return'] = (1 + stock_data['Daily_Return']).cumprod()

print(f"Daily Return for {stock_symbol}: {stock_data['Daily_Return']}")
print(f"Cumulative Return for {stock_symbol}: {stock_data['Cumulative_Return']}")

