package buildmonitor

import (
	"errors"
	"sync"
	"time"

	"github.com/level11consulting/orbitalci/models"
	"github.com/shankj3/go-til/log"
)

func NewBuildReaper() *BuildReaper {
	return &BuildReaper{contexts: make(map[string]*models.BuildContext)}
}

// BuildContext is responsible for managing all of the cancellable build contexts, and calling
// their cancel func. It will also un-track the builds that have completed
type BuildReaper struct {
	contexts map[string]*models.BuildContext
}

func (kv *BuildReaper) ListenForKillRequests(hashKillChan chan string) {
	for {
		time.Sleep(time.Millisecond)
		hash := <-hashKillChan
		kv.Kill(hash)
	}
}

// Kill will pull the cancelable context from from the context map, and call the CancelFunc() on it.
// it will then delete the hash out of the context map
func (kv *BuildReaper) Kill(killHash string) error {
	ctx, active := kv.contexts[killHash]
	if !active {
		log.Log().Warning("hash was already complete, ", killHash)
		return errors.New("hash " + killHash + " was already complete")
	}
	ctx.CancelFunc()
	delete(kv.contexts, killHash)
	return nil
}

func (kv *BuildReaper) ListenBuilds(buildsChan chan *models.BuildContext, mapLock sync.Mutex) {
	for newBuild := range buildsChan {
		mapLock.Lock()
		log.Log().Debug("got new build context for ", newBuild.Hash)
		kv.contexts[newBuild.Hash] = newBuild
		mapLock.Unlock()

	}
}

func (kv *BuildReaper) contextCleanup(buildCtx *models.BuildContext, mapLock sync.Mutex) {
	select {
	case <-buildCtx.Context.Done():
		log.Log().Debugf("build for hash %s is complete", buildCtx.Hash)
		mapLock.Lock()
		defer mapLock.Unlock()
		if _, ok := kv.contexts[buildCtx.Hash]; ok {
			delete(kv.contexts, buildCtx.Hash)
		}
		// should this be unlock?
		//mapLock.Lock()
		return
	}
}
