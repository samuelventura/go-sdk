package main

import (
	"context"
	"fmt"
	"os"

	"github.com/plutov/paypal/v4"
)

//export PAYPAL_CLIENTID=
//export PAYPAL_SECRETID=
func main() {
	clientID := os.Getenv("PAYPAL_CLIENTID")
	secretID := os.Getenv("PAYPAL_SECRETID")
	c, err := paypal.NewClient(clientID, secretID, paypal.APIBaseSandBox)
	if err != nil {
		fmt.Println("paypal.NewClient", err)
	}
	c.SetLog(os.Stdout) // Set log to terminal stdout
	accessToken, err := c.GetAccessToken(context.Background())
	if err != nil {
		fmt.Println("c.GetAccessToken", err)
	}
	fmt.Println("accessToken", accessToken)
}
