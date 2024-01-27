package main

import (
	"github.com/GeovaneCavalcante/temperatura-cep/configs"
	"github.com/GeovaneCavalcante/temperatura-cep/internal/infra/web/handler"
	"github.com/GeovaneCavalcante/temperatura-cep/internal/infra/web/webserver"
	"github.com/GeovaneCavalcante/temperatura-cep/internal/usecase/temperature"
	"github.com/GeovaneCavalcante/temperatura-cep/pkg/address/viacep"
	"github.com/GeovaneCavalcante/temperatura-cep/pkg/requester/resty"
	"github.com/GeovaneCavalcante/temperatura-cep/pkg/temperature/weather"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	requester := resty.New()
	viaCepFetcher := viacep.New(configs.ViaCepApiUrl, requester)
	weatherFetcher := weather.New(configs.WeatherApiUrl, configs.WeatherApiKey, requester)

	findTemperature := temperature.NewFindTemperatureUseCase(viaCepFetcher, weatherFetcher)

	ws := webserver.New(configs.WebServerPort)
	tH := handler.NewWebTemperatureHandler(findTemperature)
	ws.AddHandler("/temperature", tH.TemperatureByCepHandler)
	ws.Start()
}
