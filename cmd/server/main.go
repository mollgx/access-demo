package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {

	// Hello world, the web server

	newConn := func(w http.ResponseWriter, req *http.Request) {
		//io.WriteString(w, "Hello, world!\n")
		stackID := "arn:aws:cloudformation:us-east-1:098922347644:stack/runtime/11f44320-b323-11eb-962e-0a678302cb6b"
		ctx := context.Background()
		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		e2 := ec2.NewFromConfig(cfg)
		sm := ssm.NewFromConfig(cfg)

		out, err := e2.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("tag:aws:cloudformation:stack-id"),
					Values: []string{stackID},
				},
			},
		})
		if err != nil {
			log.Fatalf("unable to get stack, %v", err)
		}

		sso, err := sm.StartSession(ctx, &ssm.StartSessionInput{
			Target: out.Reservations[0].Instances[0].InstanceId,
		})
		if err != nil {
			log.Fatalf("unable to start session, %v", err)
		}

		fmt.Printf("%#v\n", *sso.StreamUrl)

		data := map[string]string{
			"url":   *sso.StreamUrl,
			"token": *sso.TokenValue,
		}
		by, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("unable to Marshal data, %v", err)
		}

		fmt.Printf("%s\n", string(by))
		fmt.Fprint(w, string(by))
	}

	http.HandleFunc("/newConn", newConn)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
