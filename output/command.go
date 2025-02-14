// Copyright 2018 github.com/andesli/gossh Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Author: andes
// Email: email.tata@qq.com

package output

import (
	//	"context"
	// "fmt"

	"github.com/andesli/gossh/log"
	"github.com/andesli/gossh/machine"

	//	"strings"
	"sync"
	"time"
)

const (
	TIMEOUT = 4500
)

//new print result
func Print(res machine.Result) {
	// log.Info("ip=%s", res.Ip)
	//index := strings.Index(cmd, ";")
	//newcmd := cmd[index+1:]
	//fmt.Printf("ip=%s|command=%s\n", ip, cmd)
	// log.Info("command=%s", res.Cmd)
	if res.Err != nil {

		log.Error("return=1\tip=%s\tcommand=%s", res.Ip, res.Cmd)
		log.Error("%s", res.Err)
	} else {
		log.Info("return=0\tip=%s\tcommand=%s", res.Ip, res.Cmd)
		log.Info("%s", res.Result)
	}
	log.Info("----------------------------------------------------------")
}

func PrintResults2(crs chan machine.Result, ls int, wt *sync.WaitGroup, ccons chan struct{}, timeout int) {
	if timeout == 0 {
		timeout = TIMEOUT
	}

	for i := 0; i < ls; i++ {
		select {
		case rs := <-crs:
			//PrintResult(rs.Ip, rs.Cmd, rs.Result)
			Print(rs)
		case <-time.After(time.Second * time.Duration(timeout)):
			log.Info("getSSHClient error,SSH-Read-TimeOut,Timeout=%ds", timeout)
		}
		wt.Done()
		<-ccons
	}

}

//print push file result
func PrintPushResult(ip, src, dst string, err error) {
	// fmt.Println("ip=", ip)
	// fmt.Println("command=", "scp "+src+" root@"+ip+":"+dst)
	if err != nil {
		log.Error("return=1\tip=%s\tPush %s to %s", ip, src, dst)
		log.Error("%s", err)
	} else {
		// fmt.Printf("return=0\n")
		// fmt.Printf("Push %s to %s ok.\n", src, dst)
		log.Info("return=0\tip=%s\tPush %s to %s ok.", ip, src, dst)
	}
	log.Info("----------------------------------------------------------")
}

//print pull result
func PrintPullResult(ip, src, dst string, err error) {
	// fmt.Println("ip=", ip)
	// fmt.Println("command=", "scp "+" root@"+ip+":"+dst+" "+src)
	if err != nil {
		// fmt.Printf("return=1\n")
		// fmt.Println(err)
		log.Error("return=1\tip=%1s\tPull from %s:%s to %s", ip, ip, dst, src)
		log.Error("%s", err)
	} else {
		// fmt.Printf("return=0\n")
		// fmt.Printf("Pull from %s to %s ok.\n", dst, src)
		log.Info("return=0\tip=%s\tPull from %s:%s to %s ok.", ip, ip, dst, src)
	}
	log.Info("----------------------------------------------------------")
}
