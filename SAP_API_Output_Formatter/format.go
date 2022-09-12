package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-stock-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToMaterialStock(raw []byte, l *logger.Logger) ([]MaterialStock, error) {
	pm := &responses.MaterialStock{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialStock. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	materialStock := make([]MaterialStock, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		materialStock = append(materialStock, MaterialStock{
        Material:                     data.Material,
        Plant:                        data.Plant,
        StorageLocation:              data.StorageLocation,
        Batch:                        data.Batch,
        Supplier:                     data.Supplier,
        Customer:                     data.Customer,
        WBSElementInternalID:         data.WBSElementInternalID,
        SDDocument:                   data.SDDocument,
        SDDocumentItem:               data.SDDocumentItem,
        InventorySpecialStockType:    data.InventorySpecialStockType,
        InventoryStockType:           data.InventoryStockType,
        MaterialBaseUnit:             data.MaterialBaseUnit,
        MatlWrhsStkQtyInMatlBaseUnit: data.MatlWrhsStkQtyInMatlBaseUnit,
		})
	}

	return materialStock, nil
}
