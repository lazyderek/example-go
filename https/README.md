## how to https



## 证书与密钥
#### 生成CA证书
```bash
openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr
openssl x509 -req -in ca.csr -signkey ca.key -out  ca.crt -days 365
```

#### 生成服务器密钥和证书
#### 
```bash
# 生成服务端私钥
openssl genrsa -out server.key 2048
# 生成服务端公钥
openssl rsa -in server.key -pubout -out server.pem
# 生成请求证书文件
openssl req -new -key server.key -out server.csr
# 用ca证书和key生成服务端证书
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt -days 365
```


#### 生成客户端密钥和证书 
```bash
# 生成客户端私钥
openssl genrsa -out client.key 2048 
# 生成客户端公钥
openssl rsa -in client.key -pubout -out client.pem
# 生成请求证书文件
openssl req -new -key client.key -out client.csr
# 用ca证书和key生成客户端证书
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr -out client.crt -days 365
```

#### ps
生成证书请求文件时，填写的common name是你的服务器域名。
在本地调试的话，可以写localhost。


在上面，我们所需要的证书和文件就已经生成了，接着可以开始我们tls
## tls双向验证
- 交互流程
 1. 客户端向服务端发起"client hello"和一串随机数A
 2. 服务端响应"server hello", 一段随机数B的密文以及服务端证书，证书里面包含了服务端公钥
 3. 客户端用ca证书检验服务端证书的有效性，并使用服务端证书的公钥解析密文;用客户端的私钥加密一串随机数B，并将其和客户端证书发送给服务端
 4. 服务端用ca证书检验服务端证书的有效性，并使用客户端的公钥解析密文，得到一串随机数。此时服务端/客户端均得到三串随机数，讲其用算法加密后，
  得到双方通信的密钥。验证通过，服务器响应结束，双方使用生成的密钥进行通信。
### 服务端代码
```
import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "http server ...........................................")
}

type Server struct {
	caCert    string
	serverCrt string
	serverKey string

	server *http.Server

	port string
}

func loadCA(caCertPath string) *x509.CertPool {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil
	}
	pool.AppendCertsFromPEM(caCrt)

	return pool
}

func newServer() *Server {
	return &Server{
		caCert:    "../etc/ca/ca.crt",
		serverCrt: "../etc/ca/server.crt",
		serverKey: "../etc/ca/server.key",
		port:      ":9091",
	}
}

func (s *Server) Run() error {
	pool := loadCA(s.caCert)

	server := http.Server{
		Addr:    s.port,
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	s.server = &server

	return s.server.ListenAndServeTLS(s.serverCrt, s.serverKey)
}

func (s *Server) Close() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		fmt.Printf("http server close, err: %v\n", err)
		return
	}
}

func main() {
	s := newServer()

	if err := s.Run(); err != nil {
		fmt.Printf("http servr run, err: %v", err)
		return
	}

	s.Close()
}
```

### 客户端代码
```
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	caCrt     string
	clientCrt string
	clientKey string

	url    string
	client *http.Client
}

func loadCA(caCertPath string) *x509.CertPool {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil
	}
	pool.AppendCertsFromPEM(caCrt)

	return pool
}

func newClient() *Client {
	client := Client{
		url:       "https://localhost:9091",
		caCrt:     "../etc/ca/ca.crt",
		clientCrt: "../etc/ca/client.crt",
		clientKey: "../etc/ca/client.key",
	}

	pool := loadCA(client.caCrt)
	cliCrt, err := tls.LoadX509KeyPair(client.clientCrt, client.clientKey)
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return nil
	}

	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			//InsecureSkipVerify: true,
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	client.client = &http.Client{Transport: &tr}

	return &client
}

func (c *Client) Do() (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func main() {
	cli := newClient()

	rsp, err := cli.Do()
	if err != nil {
		fmt.Printf("cli.Do, err: %v\n", err)
		return
	}
	defer rsp.Body.Close()

	fmt.Println("statusCode", rsp.StatusCode)

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll, err: %v\n", err)
		return
	}
	fmt.Println("body", string(b))
}
```

## tls单向验证

单向和双向验证的流程是相似的，但单向验证只需要客户端验证服务端的身份，服务端不需要验证客户端的身份。无论使用哪种验证，都是由服务端决定的

- 交互流程
 1. 客户端向服务端发起"client hello"和一串随机数A
 2. 服务端响应"server hello", 一段随机数B的密文以及服务端证书，证书里面包含了服务端公钥
 3. 客户端用ca证书检验服务端证书的有效性，并使用服务端证书的公钥解析密文;用客户端的证书的公钥加密一串随机数B，并将其和客户端证书发送给服务端
 4. 服务端用私钥解析密文，得到一串随机数。此时服务端/客户端均得到三串随机数，讲其用算法加密后，
  得到双方通信的密钥。验证通过，服务器响应结束，双方使用生成的密钥进行通信。

代码略，注释掉server的clientAuth即可
```
	server := http.Server{
		Addr:    s.port,
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs: pool,
			//ClientAuth: tls.RequireAndVerifyClientCert,	
		},
	}
```


## 参考帖子
[HTTPS实战之单向验证和双向验证](https://www.jianshu.com/p/119c4dbb7225)

[数字证书原理,公钥私钥加密原理](https://www.jianshu.com/p/072a8283c257)

