package main

type WeatherService struct {
}

func (s *WeatherService) GetWeatherSummary(lat float64, lon float64) map[string]map[string]string {
	summary := make(map[string]map[string]string)

	summary["summary"] = map[string]string{"greeting": "폭우가 내리고 있어요.",
		"temperature": "어제보다 4도 가량 춥습니다. 최고기온은 10도 최저기온은 3도 입니다.",
		"heads_up":    "내일 폭설이 내릴 수도 있으니 외출 시 주의하세요."}

	return summary
}
