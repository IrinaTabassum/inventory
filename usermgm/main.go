package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"strings"

	categorypb "codemen.org/inventory/gunk/v1/category"
	ccat "codemen.org/inventory/usermgm/core/category"
	"codemen.org/inventory/usermgm/service/category"

	customerpb "codemen.org/inventory/gunk/v1/customer"
	ccus "codemen.org/inventory/usermgm/core/customer"
	"codemen.org/inventory/usermgm/service/customer"
    
	productpb "codemen.org/inventory/gunk/v1/product"
	cpro "codemen.org/inventory/usermgm/core/product"
	"codemen.org/inventory/usermgm/service/product"

	purchasepb "codemen.org/inventory/gunk/v1/purchase"
	cpas "codemen.org/inventory/usermgm/core/puschase"
	"codemen.org/inventory/usermgm/service/purchase"

	sellpb "codemen.org/inventory/gunk/v1/sell"
	csell "codemen.org/inventory/usermgm/core/sell"
	csellpro "codemen.org/inventory/usermgm/core/sellProduct"
	"codemen.org/inventory/usermgm/service/sell"

	stockpb "codemen.org/inventory/gunk/v1/stock"
	cstock "codemen.org/inventory/usermgm/core/stock"
	"codemen.org/inventory/usermgm/service/stock"
	
	supplierpb "codemen.org/inventory/gunk/v1/supplier"
	csup "codemen.org/inventory/usermgm/core/supplier"
	"codemen.org/inventory/usermgm/service/supplier"

	salepb "codemen.org/inventory/gunk/v1/salesReport"
	csale "codemen.org/inventory/usermgm/core/salesReport"
	"codemen.org/inventory/usermgm/service/salesReport"

	userpb "codemen.org/inventory/gunk/v1/user"
	cu "codemen.org/inventory/usermgm/core/user"
	"codemen.org/inventory/usermgm/service/user"

	"codemen.org/inventory/usermgm/storage/postgres"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:embed migrations
var migrationFiles embed.FS

func main()  {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	port := config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen port: %v", err)
	}

	postgreStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}

	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err := goose.Up(postgreStorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	userCore := cu.NewCoreUser(postgreStorage)
	userSvc := user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer, userSvc)

	supplierCore := csup.NewCoreSupplier(postgreStorage)
	supplierSvc := supplier.NewSupplierSvc(supplierCore)
	supplierpb.RegisterSupplierServiceServer(grpcServer, supplierSvc)

	customerCore := ccus.NewCoreCustomer(postgreStorage)
	customerSvc := customer.NewCustomerSvc(customerCore)
	customerpb.RegisterCustomerServiceServer(grpcServer, customerSvc)
	
	categoryCore := ccat.NewCoreCategory(postgreStorage)
	categorySvc := category.NewCategorySvc(categoryCore)
	categorypb.RegisterCategoryServiceServer(grpcServer, categorySvc)

	productCore := cpro.NewCoreProduct(postgreStorage)
	prodictSvc := product.NewProductSvc(productCore)
	productpb.RegisterProductServiceServer(grpcServer, prodictSvc)

	purchaseCore := cpas.NewCorePurchase(postgreStorage)
	purchaseSvc := purchase.NewPurchaseSvc(purchaseCore)
	purchasepb.RegisterPurchaseServiceServer(grpcServer, purchaseSvc)
	
	stockCore := cstock.NewCoreStock(postgreStorage)
	stockSvc := stock.NewStockSvc(stockCore)
	stockpb.RegisterStockServiceServer(grpcServer, stockSvc)

	sellCore := csell.NewCoreSell(postgreStorage)
	sellProductCore := csellpro.NewCoreSellProduct(postgreStorage)
	sellSvc := sell.NewSellSvc(sellCore,stockCore,sellProductCore)
	sellpb.RegisterSellServiceServer(grpcServer, sellSvc)

	saleCore := csale.NewCoreSalesReport(postgreStorage)
	salekSvc := salesreport.NewPurchaseSvc(saleCore)
	salepb.RegisterSalesServiceServer(grpcServer, salekSvc)
	


	// start reflection server
	reflection.Register(grpcServer)

	fmt.Println("usermgm server running on: ", lis.Addr())
	if err := grpcServer.Serve(lis);err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}