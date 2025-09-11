package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const binanceURL = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

// GetBinanceP2PAdsPage makes a single page request to Binance P2P API
func GetBinanceP2PAdsPage(req Request) (*Response, error) {

	// Encode payload as JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", binanceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Decode JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &response, nil
}

// GetBinanceP2PAds loads all rows using paging
func GetBinanceP2PAds(config *Config) (*Response, error) {
	// Default request parameters
	baseReq := Request{
		Asset:     config.Asset,
		Fiat:      config.Fiat,
		TradeType: "BUY",
		Page:      1,
		Rows:      config.PageSize,
	}

	// Get first page to determine total count
	firstPageResp, err := GetBinanceP2PAdsPage(baseReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get first page: %v", err)
	}

	// If we got all data in one page, return it
	if len(firstPageResp.Data) >= firstPageResp.Total {
		return firstPageResp, nil
	}

	// Calculate total pages needed
	totalPages := (firstPageResp.Total + baseReq.Rows - 1) / baseReq.Rows

	// Collect all data
	allData := make([]AdvertisementItem, 0, firstPageResp.Total)
	allData = append(allData, firstPageResp.Data...)

	// Fetch remaining pages
	for page := 2; page <= totalPages; page++ {
		pageReq := baseReq
		pageReq.Page = page
		
		pageResp, err := GetBinanceP2PAdsPage(pageReq)
		if err != nil {
			return nil, fmt.Errorf("failed to get page %d: %v", page, err)
		}
		
		allData = append(allData, pageResp.Data...)
	}

	// Create combined response
	combinedResp := &Response{
		Code:          firstPageResp.Code,
		Message:       firstPageResp.Message,
		MessageDetail: firstPageResp.MessageDetail,
		Data:          allData,
		Total:         firstPageResp.Total,
		Success:       firstPageResp.Success,
	}

	return combinedResp, nil
}