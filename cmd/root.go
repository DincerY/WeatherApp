package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var lat string
var lon string
var apiKey = " "
var baseUri = "https://api.openweathermap.org/data/2.5/weather"

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type WeatherData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "uygulamanın kısa bilgisi",
	Long:  `uygulamanın uzun bilgilendirme metnini bu kısma girebilirsiniz.`,
	Run:   rootRun,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&lat, "lat", "t", "40.57", "--lat <enlem> veya -lt <enlem>")
	rootCmd.Flags().StringVarP(&lon, "lon", "n", "32.53", "--lon <boylam> veya -ln <boylam>")
}

func rootRun(cmd *cobra.Command, args []string) {
	var latQuery = "lat=" + lat
	var lonQuery = "lon=" + lon
	if lonQuery != "" && latQuery != "" {
		baseUri = baseUri + "?" + latQuery + "&" + lonQuery + "&" + "appid=" + apiKey
	}

	res, err := http.Get(baseUri)

	if err != nil {
		log.Fatal(err)
	}
	//ioutil nedir ve yerine ne kullanılmalıdır bak
	body, err := ioutil.ReadAll(res.Body)

	var weatherData WeatherData
	decoder := json.NewDecoder(bytes.NewReader(body))
	err_ := decoder.Decode(&weatherData)
	if err_ != nil {
		log.Fatal(err_)
	}
	fmt.Println("\n")
	color.Green("Enlem : %f\n", weatherData.Coord.Lat)
	color.Green("Boylam : %f\n", weatherData.Coord.Lon)
	color.Red("Açıklama : %s\n", weatherData.Weather[0].Description)
	color.Red("Genel : %s\n", weatherData.Weather[0].Main)
	color.Blue("Derece : %f\n", weatherData.Main.Temp/10)
	color.Cyan("Ülke Kodu : %s\n", weatherData.Sys.Country)
	color.Cyan("İl : %s\n", weatherData.Name)

	//buradan gelen veriyi okumaya çalışıcaz.
	//cmd ye verileri renkli bir şekilde sunucaz
	//istek sonucunda ki verinin nasıl alındığını öğrenicez
}
