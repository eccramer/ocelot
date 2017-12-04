package main

import (
	"bytes"
	"github.com/shankj3/ocelot/util/ocelog"
	"net/http"
	"time"
)

var JSONData = []byte(`{"actor":{"type":"user","username":"jessishank","display_name":"Jessi Shank","uuid":"{c8256c33-f8b2-417c-bc38-d5f60e07ecd7}","links":{"self":{"href":"https://api.bitbucket.org/2.0/users/jessishank"},"html":{"href":"https://bitbucket.org/jessishank/"},"avatar":{"href":"https://bitbucket.org/account/jessishank/avatar/32/"}}},"repository":{"type":"repository","links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules"},"avatar":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/avatar/32/"}},"uuid":"{e84d508f-5640-4c5b-8e3b-d2a41b83e3cf}","project":{"type":"","project":"","uuid":"","links":{"html":{"href":""},"avatar":{"href":""}},"key":""},"full_name":"jessishank/ocelot_test","name":"ocelot_test","website":"","owner":{"type":"user","username":"jessishank","display_name":"Jessi Shank","uuid":"{c8256c33-f8b2-417c-bc38-d5f60e07ecd7}","links":{"self":{"href":"https://api.bitbucket.org/2.0/users/jessishank"},"html":{"href":"https://bitbucket.org/jessishank/"},"avatar":{"href":"https://bitbucket.org/account/jessishank/avatar/32/"}}},"scm":"git","is_private":false},"push":{"changes":[{"new":{"type":"branch","name":"master","target":{"type":"commit","hash":"47a959ed2e701e16e6b7ffb14ec3c125906d8778","author":{"type":"author","username":"","display_name":"","uuid":"","links":{"self":{"href":""},"html":{"href":""},"avatar":{"href":""}}},"message":"testing1234\n","date":"2017-10-25T20:38:03Z","parents":[{"type":"commit","hash":"cf2732f4e22c191d7d9f31906d114f9a4290e5f4","links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/cf2732f4e22c191d7d9f31906d114f9a4290e5f4"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/cf2732f4e22c191d7d9f31906d114f9a4290e5f4"}}}],"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a"}}},"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/refs/branches/master"},"commits":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commits/master"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/branch/master"}}},"old":{"type":"branch","name":"master","target":{"type":"commit","hash":"da51988cada044e92d919fbd3d20549cf5a4baed","author":{"type":"author","username":"","display_name":"","uuid":"","links":{"self":{"href":""},"html":{"href":""},"avatar":{"href":""}}},"message":"testing 2nd commit\n","date":"2017-10-25T20:36:23Z","parents":[{"type":"commit","hash":"877e74770f6af33f5d3ce996246aeae5922144c8","links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/877e74770f6af33f5d3ce996246aeae5922144c8"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/877e74770f6af33f5d3ce996246aeae5922144c8"}}}],"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/da51988cada044e92d919fbd3d20549cf5a4baed"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/da51988cada044e92d919fbd3d20549cf5a4baed"}}},"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/refs/branches/master"},"commits":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commits/master"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/branch/master"}}},"links":{"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/branches/compare/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a..da51988cada044e92d919fbd3d20549cf5a4baed"},"diff":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/diff/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a..da51988cada044e92d919fbd3d20549cf5a4baed"},"commits":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commits?include=388710a1cf5cba5bf59dcc7e2db9cebd46886c6a&exclude=da51988cada044e92d919fbd3d20549cf5a4baed"}},"created":false,"forced":false,"closed":false,"commits":[{"hash":"388710a1cf5cba5bf59dcc7e2db9cebd46886c6a","type":"commit","message":"testing1234\n","author":{"type":"author","username":"","display_name":"","uuid":"","links":{"self":{"href":""},"html":{"href":""},"avatar":{"href":""}}},"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/388710a1cf5cba5bf59dcc7e2db9cebd46886c6a"}}},{"hash":"cf2732f4e22c191d7d9f31906d114f9a4290e5f4","type":"commit","message":"testing\n","author":{"type":"author","username":"","display_name":"","uuid":"","links":{"self":{"href":""},"html":{"href":""},"avatar":{"href":""}}},"links":{"self":{"href":"https://api.bitbucket.org/2.0/repositories/jessishank/terraform-public-modules/commit/cf2732f4e22c191d7d9f31906d114f9a4290e5f4"},"html":{"href":"https://bitbucket.org/jessishank/terraform-public-modules/commits/cf2732f4e22c191d7d9f31906d114f9a4290e5f4"}}}],"truncated":false}]}}`)

func main(){
	ch := make(chan int)
	for i := 0; i < 141; i++ {
		time.Sleep(1)
		go func(){
			//ocelog.Log().Error(i)
			ocelog.Log().Error("ab to run")
			url := "http://localhost:8088/test"
			req, err := http.NewRequest("POST", url, bytes.NewReader(JSONData))
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			_, err = client.Do(req)
			if err != nil {
				panic(err)
			}
			ch <- 0
		}()
	}
	<-ch
}

func run() {
	//fmt.Println("ab to run")
	ocelog.Log().Error("ab to run")
	url := "http://localhost:8088/test"
	req, err := http.NewRequest("POST", url, bytes.NewReader(JSONData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}