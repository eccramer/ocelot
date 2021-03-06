package circuitbreaker

import (
	"time"

	cred "github.com/level11consulting/orbitalci/common/credentials"
	"github.com/level11consulting/orbitalci/storage"
	"github.com/shankj3/go-til/nsqpb"

	//"github.com/nsqio/go-nsq"
	"testing"
)

func TestNsqWatch_MaintainHealths(t *testing.T) {
	rcHelathy := cred.NewHealthyMaintain()
	storeHealth := storage.NewHealthyStorage()
	consumer := nsqpb.NewDefaultProtoConsume()
	consumer.ConsumeMessages("testtesttesttest", "test")
	nsqw := &NsqWatch{
		interval:   1,
		pConsumers: []*nsqpb.ProtoConsume{consumer},
		remoteConf: rcHelathy,
		store:      storeHealth,
	}
	go nsqw.MaintainHealths()
	if nsqw.paused {
		t.Error("everything is up, nsq consumer  should not be paused")

	}
	rcHelathy.SetUnSuccessfulReconnect()
	rcHelathy.SetUnHealthy()
	//rcHelathy.IsHealthy = false
	//rcHelathy.SuccessfulReconnect = false
	time.Sleep(2 * time.Second)
	//if !nsqw.paused {
	//	t.Error("vault has been shut down, nsq consumer  should be paused")
	//	return
	//}
	//rcHelathy.SetHealthy()
	//rcHelathy.SetSuccessfulReconnect()
	//time.Sleep(2 * time.Second)
	//if nsqw.paused {
	//	t.Error("everything is up, nsq consumer  should not be paused")
	//}
	//storeHealth.IsHealthy = false
	//time.Sleep(2 * time.Second)
	//if !nsqw.paused {
	//	t.Error("postgres has been shut down, nsq consumer should be paused")
	//}
}
