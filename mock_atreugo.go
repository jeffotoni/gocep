package main

// import (
// 	"context"
// 	"fmt"
// 	"github.com/savsgio/atreugo/v11"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type End struct {
// 	Source string
// 	Url    string
// }

// type Result struct {
// 	Body string
// }

// var endpoints = []End{
// 	{"viacep", "https://viacep.com.br/ws/%s/json/"},
// 	{"postmon", "https://api.postmon.com.br/v1/cep/%s"},
// 	{"republicavirtual", "https://republicavirtual.com.br/web_cep.php?cep=%s&formato=json"},
// }

// func main() {

// 	config := atreugo.Config{
// 		Addr: "0.0.0.0:8084",
// 	}

// 	server := atreugo.New(config)

// 	v1 := server.NewGroupPath("/v1")
// 	v1.GET("/{cep:[0-9]{8}}", func(ctx *atreugo.RequestCtx) error {

// 		cep := fmt.Sprintf("%s", ctx.UserValue("cep"))

// 		var mux sync.Mutex
// 		var result Result
// 		var wg sync.WaitGroup
// 		wg.Add(3)
// 		for _, e := range endpoints {
// 			endpoint := fmt.Sprintf(e.Url, cep)
// 			go func(endpoint string, result *Result) {
// 				defer wg.Done()
// 				ctx2, cancel := context.WithCancel(context.TODO())
// 				afterFuncTimer := time.AfterFunc(50*time.Millisecond, func() {
// 					cancel()
// 				})
// 				defer afterFuncTimer.Stop()

// 				req, err := http.NewRequest("GET", endpoint, nil)
// 				if err != nil {
// 					log.Println("Error NewRequest:", err)
// 					return
// 				}

// 				req = req.WithContext(ctx2)

// 				response, err := http.DefaultClient.Do(req)
// 				if err != nil {
// 					log.Println("Error ClientDo:", err)
// 					return
// 				}

// 				defer response.Body.Close()

// 				requestContent, err := ioutil.ReadAll(response.Body)
// 				if err != nil {
// 					log.Println("Error ioutil.ReadAll:", err)
// 					return
// 				}

// 				if len(string(requestContent)) > 0 &&
// 					response.StatusCode == http.StatusOK {
// 					println("entrei....::")
// 					mux.Lock()
// 					defer mux.Unlock()
// 					result.Body = string(requestContent)
// 					return
// 				}
// 			}(endpoint, &result)
// 		}
// 		wg.Wait()

// 		return ctx.TextResponse(result.Body)
// 	})

// 	if err := server.ListenAndServe(); err != nil {
// 		fmt.Println(err)
// 	}
// }
