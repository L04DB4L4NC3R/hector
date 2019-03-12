package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/angadsharma1016/hector/controller"
)

type server struct {
	port        string
	host        string
	controllers map[string]*http.Handler
}

func (s *server) Startup() {
	log.Println("Listening...")
	controller.StartClient()
	log.Fatal(http.ListenAndServe(s.port, nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := ethclient.Dial(os.Getenv("INFURA_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
	log.Println(address)
	// instance, err := store.NewStore(address, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("INSTANCE \n\n\n")
	// log.Println(instance)
	s := server{port: ":3000", host: "0.0.0.0", controllers: nil}
	s.Startup()
}
