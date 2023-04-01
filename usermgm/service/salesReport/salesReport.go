package salesreport

import (
	"context"

	salespb "codemen.org/inventory/gunk/v1/salesReport"
	"codemen.org/inventory/usermgm/storage"
)

type CoreSalesReport interface {
	ListOfSalesRepost(storage.SaleReportFilter) ([]storage.SaleReport, error)
}

type SalesReportSvc struct {
	salespb.UnimplementedSalesServiceServer
	core CoreSalesReport
}

func NewPurchaseSvc(cr CoreSalesReport) *SalesReportSvc {
	return &SalesReportSvc{
		core: cr,
	}
}

func (srs SalesReportSvc) ListSalesReport(ctx context.Context, r *salespb.ListSalesReportRequest) (*salespb.ListSalesReportResponse, error) {
	Salefilter := storage.SaleReportFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	srl, err := srs.core.ListOfSalesRepost(Salefilter)
	if err != nil {
		return nil, err
	}

	list := make([]*salespb.SalesReport, len(srl))
	for i, salePro := range srl {
		list[i] = &salespb.SalesReport{
			ProductId:        int32(salePro.ProductId),
			ProductName:      salePro.ProductName,
			PurchaseQuantity: int32(salePro.PurchaseQuantity),
			SellQuantity:     int32(salePro.SellQuantity),
			StockQuantity:    int32(salePro.StockQuantity),
		}
	}
	return &salespb.ListSalesReportResponse{
		SalesReports: list,
	}, nil
}
