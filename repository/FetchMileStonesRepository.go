package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MileStoneResponse struct {
	// TODO: 使うresponseの精査
	Url string `json:"html_url"`
	Due string `json:"due_on"`
}

type FetchMileStonesRepository interface {
	FetchAllMileStones() ([]MileStoneResponse, error)
}

// TODO: interfaceと実装分ける
type FetchMilesStoneHttpClientRepository struct {}

func NewFetchMileStonesHttpClientRepository() *FetchMilesStoneHttpClientRepository {
	return &FetchMilesStoneHttpClientRepository{}
}

func (repo *FetchMilesStoneHttpClientRepository) FetchAllMileStones() ([]MileStoneResponse, error) {
	// TODO: bnn-tk/test-auto-releaseのような organization/test-auto-releaseは可変にしたい
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/bnn-tk/test-auto-release/milestones", nil)

	// TODO: configオブジェクト作ってそこからenvは読み込みたい
	req.SetBasicAuth(os.Getenv("ORGANIZATION"), os.Getenv("TOKEN"))
	// TODO: timeoutの設定等 別途 clientオブジェクト作っていいかも
	client := http.DefaultClient
	res, err := client.Do(req); if err != nil {
		// TODO: status code ごとのエラーハンドリング
		_ = fmt.Errorf("error2: %v", err)
		return nil, err
	}
	var mileStones []MileStoneResponse
	_ = json.NewDecoder(io.TeeReader(res.Body, os.Stderr)).Decode(&mileStones)
	return mileStones, nil
}