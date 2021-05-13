package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mmmorris1975/ssm-session-client/ssmclient"
)

func main() {
	resp, err := http.Get("http://localhost:8080/newConn")
	if err != nil {
		log.Fatalf("unable to get ws, %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unable to read body, %v", err)
	}

	data := map[string]string{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("unable to unmarshal body, %v", err)
	}

	fmt.Printf("%#v\n", data)

	log.Fatal(ssmclient.ShellSessionWithURL(data["url"], data["token"]))

	/*
		sso, err := sm.StartSession(ctx, &ssm.StartSessionInput{
			Target: out.Reservations[0].Instances[0].InstanceId,
		})
		if err != nil {
			log.Fatalf("unable to start session, %v", err)
		}

		fmt.Printf("%#v\n", *sso.StreamUrl)

			c, _, err := websocket.DefaultDialer.Dial(*sso.StreamUrl, nil)
			if err != nil {
				log.Fatalf("unable to dial session, %v", err)
			}
			defer c.Close()

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			rl, err := readline.NewEx(&readline.Config{
				Prompt: "\033[32m»\033[0m ",
			})
			if err != nil {
				log.Fatalf("unable to create readline, %v", err)
			}
			defer rl.Close()

			go func() {
				for {
					_, message, err := c.ReadMessage()
					if err != nil {
						fmt.Printf("<<server: %s>>\n", err)
						close(interrupt)
						return
					}
					tsStr := ""

					io.WriteString(rl.Stdout(), fmt.Sprintf("%s\033[31m«\033[0m %s\n", tsStr, message))
				}
			}()

			go func() {
				for {
					line, err := rl.Readline()
					if err == readline.ErrInterrupt {
						interrupt <- os.Interrupt
						return
					} else if err != nil {
						close(interrupt)
						return
					}
					if len(line) == 0 {
						continue
					}
					err = c.WriteMessage(websocket.TextMessage, []byte(line))
					if err != nil {
						fmt.Println("err:", err)
						return
					}
				}
			}()

			select {
			case <-interrupt:
				c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				fmt.Println("<<client: sent websocket close frame>>")
				c.Close()
				os.Exit(0)
			}*/
}
