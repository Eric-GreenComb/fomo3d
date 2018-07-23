package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
)

var (
	g errgroup.Group

	// ServerMode ServerMode
	ServerMode = flag.String("sm", "release", "server listen port")
	// ServerPort server port
	ServerPort = flag.String("sp", "8080", "server listen port")
	// EthHost eth host
	EthHost = flag.String("eh", "http://10.0.21.112:18545", "eth host")
)

// Clients global client connection
var Clients = &Conn{}

// Conn 链的链接走这里
type Conn struct {
	Eth *rpc.Client
}

// Init 初始化链接
func Init() error {
	fmt.Println(">>>>>>>>> start dial ethereum >>>>>>>>>>>")
	eth, err := rpc.Dial(*EthHost)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(">>>>>>>>> dialed ethereum >>>>>>>>>>>")

	Clients.Eth = eth

	return nil
}

// Cors Cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {

	gin.SetMode(*ServerMode)

	// Init 初始化链接
	Init()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", Index)
		r0.GET("health", Health)
	}

	rfomo3d := router.Group("/fomo3d")
	{
		rfomo3d.GET("/get_timeleft/:conaddr", GetTimeLeft)
	}

	server := &http.Server{
		Addr:         ":" + *ServerPort,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

// Index Index
func Index(c *gin.Context) {
	c.String(http.StatusOK, "eth-server index")
}

// Health Health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

// GetTimeLeft GetTimeLeft
func GetTimeLeft(c *gin.Context) {
	_conAddress := c.Params.ByName("conaddr")

	_client := ethclient.NewClient(Clients.Eth)
	defer _client.Close()

	tc, err := NewFoMo3DlongCaller(common.HexToAddress(_conAddress), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_value, err := tc.GetTimeLeft(&bind.CallOpts{Pending: true})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_timeNow := time.Now().UnixNano()/1e6 + _value.Int64()*1000

	c.JSON(http.StatusOK, gin.H{"now": _timeNow})
}
