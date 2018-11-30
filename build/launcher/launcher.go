package launcher

import (
<<<<<<< HEAD
	"github.com/level11consulting/ocelot/build/basher"
	"github.com/level11consulting/ocelot/build/integrations"
	"github.com/level11consulting/ocelot/build/valet"
	"github.com/level11consulting/ocelot/common/credentials"
	"github.com/level11consulting/ocelot/models"
	"github.com/level11consulting/ocelot/storage"
=======
	"github.com/shankj3/go-til/nsqpb"
	"github.com/shankj3/ocelot/build/basher"
	"github.com/shankj3/ocelot/build/integrations"
	"github.com/shankj3/ocelot/build/valet"
	"github.com/shankj3/ocelot/common/credentials"
	"github.com/shankj3/ocelot/models"
	"github.com/shankj3/ocelot/storage"
>>>>>>> trying to revert everything
)

// create a struct that is werker msg handler creates when it receives a new nsq message

type launcher struct {
	*models.WerkerFacts
	RemoteConf   credentials.CVRemoteConfig
	infochan     chan []byte
	StreamChan   chan *models.Transport
	BuildCtxChan chan *models.BuildContext
	Basher       *basher.Basher
	Store        storage.OcelotStorage
	BuildValet   *valet.Valet
	handler      models.VCSHandler
	integrations []integrations.StringIntegrator
	binaryIntegs []integrations.BinaryIntegrator
	producer     nsqpb.Producer
}

func NewLauncher(facts *models.WerkerFacts,
	remoteConf credentials.CVRemoteConfig,
	streamChan chan *models.Transport,
	BuildCtxChan chan *models.BuildContext,
	bshr *basher.Basher,
	store storage.OcelotStorage,
	bv *valet.Valet,
	producer nsqpb.Producer) *launcher {
	return &launcher{
		WerkerFacts:  facts,
		RemoteConf:   remoteConf,
		StreamChan:   streamChan,
		BuildCtxChan: BuildCtxChan,
		Basher:       bshr,
		Store:        store,
		BuildValet:   bv,
		infochan:     make(chan []byte),
		integrations: getIntegrationList(),
		binaryIntegs: getBinaryIntegList(bshr.LoopbackIp, facts.ServicePort),
		producer:     producer,
	}
}
