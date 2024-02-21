package main

type WeatherService struct {
}

func (s *WeatherService) GetWeatherSummary(lat float64, lon float64) Summary {
	return Summary{
		SummaryStatement{
			"폭우가 내리고 있어요.",
			"어제보다 4도 가량 춥습니다. 최고기온은 10도 최저기온은 3도 입니다.",
			"내일 폭설이 내릴 수도 있으니 외출 시 주의하세요.",
		},
	}
}

type Summary struct {
	SummaryStatement `json:"summary"`
}

type SummaryStatement struct {
	Greeting    string `json:"greeting"`
	Temperature string `json:"temperature"`
	Heads_up    string `json:"heads_up"`
}
