package runtime

import (
	stringbuilder "github.com/level11consulting/orbitalci/build/helpers/stringbuilder/accountrepo"
	"github.com/level11consulting/orbitalci/build/notifiers"
	"github.com/level11consulting/orbitalci/build/notifiers/slack"
	"github.com/level11consulting/orbitalci/models"
	"github.com/level11consulting/orbitalci/models/pb"
	"github.com/pkg/errors"
)

func getNotifiers() []notifiers.Notifier {
	return []notifiers.Notifier{slack.Create()}
}

// doNotifications will notify everything you want it to. should be called at the end of a build
func (w *launcher) doNotifications(werk *pb.WerkerTask) error {
	accountName, _, err := stringbuilder.GetAcctRepo(werk.FullName)
	if err != nil {
		return errors.Wrap(err, "unable to split full name into acct/repo")
	}
	notifys := getNotifiers()
	stageResults, err := w.Store.RetrieveStageDetail(werk.Id)
	if err != nil {
		return err
	}
	buildSum, err := w.Store.RetrieveSumByBuildId(werk.Id)
	if err != nil {
		return err
	}
	fullResult := models.ParseStagesByBuildId(buildSum, stageResults)
	// if the status of this build doesn't match up with the notifications' on, then don't run a notification
	for _, notify := range notifys {
		if !notify.IsRelevant(werk.BuildConf, buildSum.Status) {
			continue
		}
		credz, err := w.RemoteConf.GetCredsBySubTypeAndAcct(w.Store, notify.SubType(), accountName, false)
		if err != nil {
			return err
		}

		err = notify.RunIntegration(credz, fullResult, werk.BuildConf.Notify)
		if err != nil {
			return err
		}
	}
	return nil
}
