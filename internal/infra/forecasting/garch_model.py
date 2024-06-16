import pandas as pd
from arch import arch_model

def train_and_forecast(data_path, periods):
    data = pd.read_csv(data_path)
    data.columns = ['ds', 'y']
    data['returns'] = data['y'].pct_change().dropna()

    model = arch_model(data['returns'], vol='Garch', p=1, q=1)
    model_fit = model.fit(disp='off')

    forecast = model_fit.forecast(horizon=periods)
    forecast_df = forecast.mean.dropna()
    forecast_df.columns = ['yhat']
    forecast_df.to_csv('../data/garch_forecast.csv', index=False)
    print("GARCH forecasting complete. Results saved to garch_forecast.csv.")

if __name__ == "__main__":
    import sys
    train_and_forecast(sys.argv[1], int(sys.argv[2]))
