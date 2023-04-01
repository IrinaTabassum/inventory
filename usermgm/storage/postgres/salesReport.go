package postgres

import (
	"codemen.org/inventory/usermgm/storage"
)

const salesReportQuery = `WITH tot AS 
(select count(*) as total FROM products WHERE deleted_at IS NULL AND (name ILIKE '%%' || $1 || '%%'))
SELECT distinct(P.id) as product_id , P.name, P.deleted_at,
(select COALESCE(sum(quantity),0) from purchases where purchases.product_id=P.id  group by purchases.product_id) as purchase_quantity,
(select COALESCE(sum(quantity),0) from sell_products where sell_products.product_id=P.id group by sell_products.product_id) as sell_quantity,
COALESCE(stocks.quantity, 0) as stock_quantity,
tot.total as total
FROM products as P
LEFT JOIN stocks ON P.id = stocks.product_id
LEFT JOIN tot ON TRUE
WHERE P.deleted_at IS NULL AND (P.name ILIKE '%%' || $1 || '%%')
	ORDER BY P.id DESC 
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfSalesRepost(sf storage.SaleReportFilter) ([]storage.SaleReport, error) {
	var listSaleReport []storage.SaleReport
	if sf.Limit == 0 {
		sf.Limit = 10
	}

	if err := s.DB.Select(&listSaleReport, salesReportQuery, sf.SearchTerm, sf.Offset, sf.Limit); err != nil {
		return nil, err
	}

	return listSaleReport, nil
}
