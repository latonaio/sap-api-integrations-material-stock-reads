package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-material-stock-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetMaterialStock(material, plant, storageLocation string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialStock":
			func() {
				c.MaterialStock(material, plant, storageLocation)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) MaterialStock(material, plant, storageLocation string) {
	materialStockData, err := c.callMaterialStockSrvAPIRequirementMaterialStock("A_MatlStkInAcctMod", material, plant, storageLocation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialStockData)
}

func (c *SAPAPICaller) callMaterialStockSrvAPIRequirementMaterialStock(api, material, plant, storageLocation string) ([]sap_api_output_formatter.MaterialStock, error) {
	url := strings.Join([]string{c.baseURL, "API_MATERIAL_STOCK_SRV", api}, "/")
	param := c.getQueryWithMaterialStock(map[string]string{}, material, plant, storageLocation)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialStock(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithMaterialStock(params map[string]string, material, plant, storageLocation string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s' and StorageLocation eq '%s'", material, plant, storageLocation)
	return params
}
