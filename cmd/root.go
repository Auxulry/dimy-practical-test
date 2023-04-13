package cmd

import (
	"context"
	"fmt"
	"github.com/MochamadAkbar/dimy-practical-test/common"
	"github.com/MochamadAkbar/dimy-practical-test/configs"
	"github.com/MochamadAkbar/dimy-practical-test/domain/handler"
	"github.com/MochamadAkbar/dimy-practical-test/domain/repository"
	"github.com/MochamadAkbar/dimy-practical-test/domain/usecase"
	"github.com/MochamadAkbar/dimy-practical-test/exception"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/router"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var (
	svPort  string
	rootCmd = &cobra.Command{
		Use:   "service",
		Short: "Running the gRPC service",
		Long:  "Used to run gRPC Service including rpc server, rpc client and gateway",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			ctx := context.Background()

			cfg, err := configs.LoadConfig(".")
			if err != nil {
				panic(err)
			}
			db, err := orm.NewPSQL(ctx,
				cfg.DBDNS,
				&orm.ConfigConnProvider{
					ConnMaxLifetime: time.Hour,
					ConnMaxIdleTime: time.Minute,
					MaxOpenConns:    10,
					MaxIdleConns:    10,
				}, &gorm.Config{})
			if err != nil {
				panic(err)
			}

			r := router.NewRouter()
			r.Use(middleware.DefaultLogger)
			orderRepository := repository.NewOrderRepository(db)
			productRepository := repository.NewProductRepository(db)
			paymentRepository := repository.NewPaymentRepository(db)
			customerRepository := repository.NewCustomerRepository(db)
			addressRepository := repository.NewCustomerAddressRepository(db)
			itemRepository := repository.NewOrderItemRepository(db)

			orderUsecase := usecase.NewOrderUsecase(orderRepository, productRepository, paymentRepository, customerRepository, addressRepository, itemRepository)
			err = handler.NewOrderHandler(orderUsecase, r)
			if err != nil {
				panic(err)
			}

			r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(200)
				writer.Write([]byte(string("welcome to service")))
			})

			server := common.NewServer(r, fmt.Sprintf(":%v", svPort))

			logger.Infoln("Serving on", svPort)
			if err := server.ListenAndServe(); err != nil {
				panic(exception.NewException("[stop] server failed to start : " + err.Error()))
			}
		},
	}
)

func Execute() {
	rootCmd.Flags().StringVarP(&svPort, "svport", "s", "", "define server port")
	rootCmd.MarkFlagsRequiredTogether("svport")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}