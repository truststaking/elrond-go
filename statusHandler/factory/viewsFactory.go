package factory

import (
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data/endProcess"
	"github.com/ElrondNetwork/elrond-go/statusHandler"
	"github.com/ElrondNetwork/elrond-go/statusHandler/view"
	"github.com/ElrondNetwork/elrond-go/statusHandler/view/termuic"
)

type viewsFactory struct {
	presenter                 view.Presenter
	chanNodeStop              chan endProcess.ArgEndProcess
	refreshTimeInMilliseconds int
}

// NewViewsFactory is responsible for creating a new viewers factory object
func NewViewsFactory(presenter view.Presenter, chanNodeStop chan endProcess.ArgEndProcess, refreshTimeInMilliseconds int) (*viewsFactory, error) {
	if check.IfNil(presenter) {
		return nil, statusHandler.ErrNilPresenterInterface
	}
	if chanNodeStop == nil {
		return nil, statusHandler.ErrNilNodeStopChannel
	}

	return &viewsFactory{
		presenter:                 presenter,
		chanNodeStop:              chanNodeStop,
		refreshTimeInMilliseconds: refreshTimeInMilliseconds,
	}, nil
}

// Create returns an view slice that will hold all views in the system
func (vf *viewsFactory) Create() ([]Viewer, error) {
	views := make([]Viewer, 0)

	termuiConsole, err := vf.createTermuiConsole()
	if err != nil {
		return nil, err
	}
	views = append(views, termuiConsole)

	return views, nil
}

func (vf *viewsFactory) createTermuiConsole() (*termuic.TermuiConsole, error) {
	termuiConsole, err := termuic.NewTermuiConsole(vf.presenter, vf.chanNodeStop, vf.refreshTimeInMilliseconds)
	if err != nil {
		return nil, err
	}

	return termuiConsole, nil
}
