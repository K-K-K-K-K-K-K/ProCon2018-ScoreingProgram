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
	TilesA	[]Tile	`json:"tiles"`
	TilesB	[]Tile	`json:"tiles"`
}
    
type Status struct {
	Field	Field	`json:"field"`
	Teams	[]Team	`json:"teams"`
}

func GetJSONBytes(status Status) ([]byte, error) {
	statusBytes, err := json.Marshal(status)
	if err != nil {
		return statusBytes, errors.New("Status構造体のJSON文字列化に失敗")
	}
	return statusBytes, nil
}

func GetJSONString(status Status, indent bool) (string, error) {
	statusBytes, err := GetJSONBytes(status);
	if err != nil {
		return string(statusBytes), err
	}
	if !indent {
		return string(statusBytes), nil
	} else {
		indentedStatusBytes := new(bytes.Buffer)
		json.Indent(indentedStatusBytes, statusBytes, "", "    ")
		return indentedStatusBytes.String(), nil
	}
}

type Data struct {
	TilePoint		int		`json:"tile_point"`
	TerritoryPoint	int		`json:"territory_point"`
	TileArea		[]int	`json:"tile_area"`
}

type Response struct {
	Data		[]Data	`json:"data"`
	ResponseID	string	`json:"response_id"`
	Error		string	`json:"error"`
}

func GetResponse(status Status) (Response, error) {
	// JSON化
	statusBytes, err := GetJSONBytes(status)
	if err != nil {
		return Response{}, err
	}


	// POST実行
	req, err := http.NewRequest(
		"POST",
		Endpoint,
		bytes.NewBuffer([]byte(statusBytes)),
	)
	if err != nil {
		fmt.Println("POST要求の作成に失敗")
		return Response{}, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", "************")
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("POST実行に失敗")
		return Response{}, err
	}
	
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("応答解析に失敗")
		return Response{}, err
	}

	fmt.Println(string(b))


	// レスポンスをデコード
	return Response{}, err
}

// JSONの構造把握できてない -> 修正
// GetResponseが適当すぎる -> ほならね
// レスポンス解析なされていないんですがそれは -> 諦める
