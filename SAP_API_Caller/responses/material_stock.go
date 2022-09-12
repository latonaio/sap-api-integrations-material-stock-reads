package responses

type MaterialStock struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
		Material                     string   `json:"Material"`
		Plant                        string   `json:"Plant"`
		StorageLocation              string   `json:"StorageLocation"`
		Batch                        string   `json:"Batch"`
		Supplier                     string   `json:"Supplier"`
		Customer                     string   `json:"Customer"`
		WBSElementInternalID         string   `json:"WBSElementInternalID"`
		SDDocument                   string   `json:"SDDocument"`
		SDDocumentItem               string   `json:"SDDocumentItem"`
		InventorySpecialStockType    string   `json:"InventorySpecialStockType"`
		InventoryStockType           string   `json:"InventoryStockType"`
		MaterialBaseUnit             string   `json:"MaterialBaseUnit"`
		MatlWrhsStkQtyInMatlBaseUnit string   `json:"MatlWrhsStkQtyInMatlBaseUnit"`
		} `json:"results"`
	} `json:"d"`
}
