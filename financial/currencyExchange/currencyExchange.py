import requests

# Define the API endpoint
api_endpoint = 'https://api.exchangerate-api.com/v4/latest/'

# Define the base currencies
base_currencies = ['EUR', 'USD', 'EUR']

# Define the target currencies
target_currencies = ['USD', 'AED', 'AED']

# Fetch exchange rate data for each currency pair
for base_currency, target_currency in zip(base_currencies, target_currencies):
    response = requests.get(api_endpoint + base_currency)
    data = response.json()
    exchange_rate = data['rates'][target_currency]
    print(f"Latest exchange rate for {base_currency}/{target_currency}: {exchange_rate}")
