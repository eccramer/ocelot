package sshhelper

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/shankj3/go-til/test"
)

func Test_splitEnvs(t *testing.T) {
	envs := []string{
		"HERE=encodeddata==",
		"ANOTHERLINE=;alisdf8xsa8dfalw3jnv8dsaa;sdkfne82,vxcug74-a;lsn",
		"ASUPER_LONG_LINE=jjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjj",
		"ABASICLINE=84,xmc83b!!!!!]n\\n\n\n\n\\n\\n=32j218vne6yyyyyyayyayayyaya",
	}
	expected := [][2]string{
		{"HERE", "encodeddata=="},
		{"ANOTHERLINE", ";alisdf8xsa8dfalw3jnv8dsaa;sdkfne82,vxcug74-a;lsn"},
		{"ASUPER_LONG_LINE", "jjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjj"},
		{"ABASICLINE", "84,xmc83b!!!!!]n\\n\n\n\n\\n\\n=32j218vne6yyyyyyayyayayyaya"},
	}
	split := splitEnvs(envs)
	if diff := deep.Equal(split, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContextConnection_CheckConnection(t *testing.T) {
	cleanup, ctx := CreateSSHDockerContainer(t, "2222")
	defer cleanup()
	cnxn := InitContextConnect("./test-fixtures/docker_id_rsa", "", "root", "localhost", 2222)
	err := cnxn.Connect(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if err := cnxn.CheckConnection(); err != nil {
		t.Error(err)
		return
	}
}


func TestContextConnection_RunAndLog(t *testing.T) {
	cleanup, ctx := CreateSSHDockerContainer(t, "2223")
	defer cleanup()
	cnxn := InitContextConnect("./test-fixtures/docker_id_rsa", "","root", "localhost", 2223)
	err := cnxn.Connect(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer cnxn.Close()
	here := make(chan []byte, 1000)
	err = cnxn.RunAndLog("echo 'OH YEAH OH YEAH OH YEAH\nohnoohnoohno'", []string{}, here, testPipeHandler)
	if err != nil {
		t.Error(err)
	}
	close(here)
	var totallist []string
	for i := range here {
		totallist = append(totallist, string(i))
	}
	total := strings.Join(totallist, "\n")
	if total != "OH YEAH OH YEAH OH YEAH\nohnoohnoohno" {
		t.Error(test.StrFormatErrors("channel data", "OH YEAH OH YEAH OH YEAH\nohnoohnoohno", total))
	}
}

func testPipeHandler(rc io.Reader, logout chan[]byte, donechan chan int) {
	defer close(donechan)
	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		//fmt.Println("bytes")
		logout <- scanner.Bytes()
	}
	//fmt.Println("DOING PIPE WERK")
	//fmt.Println("closing done chan")
}


func TestContextConnection_Setenvs(t *testing.T) {
	cleanup, ctx := CreateSSHDockerContainer(t, "2224")
	defer cleanup()
	cnxn := InitContextConnect("./test-fixtures/docker_id_rsa", "","root", "localhost", 2224)
	err := cnxn.Connect(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer cnxn.Close()
	cnxn.SetGlobals([]string{"IVORYTRADE=BAD", "GIT_HASH=nd8sb29"})
	logout := make(chan []byte, 1000)
	if err = cnxn.RunAndLog("echo $RUNTIME && echo $IVORYTRADE && echo $GIT_HASH", []string{"RUNTIME=1", "LONG="+ SUPERLONGLINE}, logout, testPipeHandler); err != nil {
		t.Error(err)
		return
	}
	close(logout)
	var totallist []string
	for i := range logout {
		totallist = append(totallist, string(i))
	}
	if len(totallist) != 3 {
		t.Error("something went awry, list is: ", totallist)
		return
	}
	if totallist[0] != "1" {
		t.Error(test.StrFormatErrors("RUNTIME value", "1", totallist[0]))
	}
	if totallist[1] != "BAD" {
		t.Error(test.StrFormatErrors("IVORYTRADE value", "BAD", totallist[1]))
	}
	if totallist[2] != "nd8sb29" {
		t.Error(test.StrFormatErrors("GIT_HASH value", "nd8sb29", totallist[2]))
	}


}

const SUPERLONGLINE=`jjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j
=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjj
jjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j
=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj
=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj
=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=
j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjj
jjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjj
jj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=
jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjj
jjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjj
jjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjj
jjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjj
jjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjjjjjjjjjj=j=jjjjj`