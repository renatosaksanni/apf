import pandas as pd
import matplotlib.pyplot as plt

def visualize_forecast(model):
    if model == 'prophet':
        forecast = pd.read_csv('../data/prophet_forecast.csv')
        plt.plot(forecast['ds'], forecast['yhat'], label='Prophet Forecast')
        plt.fill_between(forecast['ds'], forecast['yhat_lower'], forecast['yhat_upper'], color='gray', alpha=0.2)
    elif model == 'garch':
        forecast = pd.read_csv('../data/garch_forecast.csv')
        plt.plot(forecast.index, forecast['yhat'], label='GARCH Forecast')

    plt.xlabel('Date')
    plt.ylabel('Transactions')
    plt.title(f'{model.capitalize()} Transaction Volume Forecast')
    plt.legend()
    plt.savefig(f'../data/{model}_forecast_plot.png')
    plt.show()

if __name__ == "__main__":
    import sys
    visualize_forecast(sys.argv[1])
