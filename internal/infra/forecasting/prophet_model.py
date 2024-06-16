import pandas as pd
from prophet import Prophet

def train_and_forecast(data_path, periods):
    data = pd.read_csv(data_path)
    data.columns = ['ds', 'y']

    model = Prophet()
    model.fit(data)

    future = model.make_future_dataframe(periods=periods)
    forecast = model.predict(future)
    forecast[['ds', 'yhat', 'yhat_lower', 'yhat_upper']].to_csv('../data/prophet_forecast.csv', index=False)
    print("Prophet forecasting complete. Results saved to prophet_forecast.csv.")

if __name__ == "__main__":
    import sys
    train_and_forecast(sys.argv[1], int(sys.argv[2]))
