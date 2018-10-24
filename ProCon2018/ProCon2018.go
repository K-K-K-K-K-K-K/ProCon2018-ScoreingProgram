package ProCon2018

import (
	"errors"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"

	"fmt"
)

const Endpoint string = "https://42isf6z498.execute-api.ap-northeast-1.amazonaws.com/dev"

type Field struct {
	Scores	[]int	`json:"scores"`
	Height	int		`json:"height"`
	Width	int		`json:"width"`
}

type Tile struct {
	Y		int		`json:"y"`
	X		int		`json:"x"`
}

type Team struct {
	Tiles	[]Tile	`json:"tiles"`
}
    
type Status struct {
	Field	Field	`json:"field"`
	Teams	[]Team	`json:"teams"`
}

func GenerateJSONBytes(status Status) ([]byte, error) {
	stBs, err := json.Marshal(status)
	if err != nil {
		return stBs, errors.New("Status構造体のバイト列化に失敗")
	}
	return stBs, nil
}

func GenerateJSONString(status Status) (string, error) {
	stBs, err := GenerateJSONBytes(status);
	if err != nil {
		return string(stBs), err
	}
	indStBs := new(bytes.Buffer)
	json.Indent(indStBs, stBs, "", "    ")
	return indStBs.String(), nil
}

type Data struct {
	TilePoint		int		`json:"tile_point"`
	TerritoryPoint	int		`json:"territory_point"`
	TileArea		[]int	`json:"tile_area"`
}

type Result struct {
	Data		[]Data	`json:"data"`
	ResponseID	string	`json:"response_id"`
	Error		string	`json:"error"`
}

func BuildResult(resBs []byte) (Result, error) {
	return Result{}, nil
}

func SendRequest(stBs []byte) (Result, error) {
	req, err := http.NewRequest(
		"POST",
		Endpoint,
		bytes.NewBuffer([]byte(stBs)),
	)
	if err != nil {
		fmt.Println("POST要求の作成に失敗")
		return Result{}, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", "*****")
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("POST実行に失敗")
		return Result{}, err
	}
	
	resBs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("応答解析に失敗")
		return Result{}, err
	}

	fmt.Println(string(resBs)) // 験
	return BuildResult(resBs)
}

func GetResult(status Status) (Result, error) {
	stBs, err := GenerateJSONBytes(status)
	if err != nil {
		return Result{}, err
	}

	return SendRequest(stBs)
}

// JSONの構造把握できてない -> 修正
// レスポンス解析なされていないんですがそれは -> 諦める
