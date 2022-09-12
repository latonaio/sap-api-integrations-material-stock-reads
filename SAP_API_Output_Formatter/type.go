package sap_api_output_formatter

type MaterialStockReads struct {
	 ConnectionKey string `json:"connection_key"`
	 Result        bool   `json:"result"`
	 RedisKey      string `json:"redis_key"`
	 Filepath      string `json:"filepath"`
	 APISchema     string `json:"api_schema"`
	 Material      string `json:"material_code"`
	 Deleted       bool   `json:"deleted"`
}

type MaterialStock struct {
     Material                     string `json:"Material"`
     Plant                        string `json:"Plant"`
     StorageLocation              string `json:"StorageLocation"`
     Batch                        string `json:"Batch"`
     Supplier                     string `json:"Supplier"`
     Customer                     string `json:"Customer"`
     WBSElementInternalID         string `json:"WBSElementInternalID"`
     SDDocument                   string `json:"SDDocument"`
     SDDocumentItem               string `json:"SDDocumentItem"`
     InventorySpecialStockType    string `json:"InventorySpecialStockType"`
     InventoryStockType           string `json:"InventoryStockType"`
     MaterialBaseUnit             string `json:"MaterialBaseUnit"`
     MatlWrhsStkQtyInMatlBaseUnit string `json:"MatlWrhsStkQtyInMatlBaseUnit"`
}
