package salesreport

import (
	"codemen.org/inventory/usermgm/storage"
)



type SalesStoreReport interface {
	ListOfSalesRepost(storage.SaleReportFilter) ([]storage.SaleReport, error)
}
type CoreSalesReport struct { 
	store SalesStoreReport
}

func NewCoreSalesReport(ss SalesStoreReport) *CoreSalesReport {
	return &CoreSalesReport{
		store: ss,
	}
}

func (csr CoreSalesReport) ListOfSalesRepost(sf storage.SaleReportFilter) ([]storage.SaleReport, error){
	sales, err := csr.store.ListOfSalesRepost(sf)
	if err != nil {
		return nil, err
	}

	return sales, nil
}


