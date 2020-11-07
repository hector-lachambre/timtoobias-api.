package presenters

import (
	"gitlab.com/timtoobias-projects/timtoobias-api/viewmodels"
	"gitlab.com/timtoobias-projects/timtoobias-core/usecases"
)

type LiveNotifierPresenter struct {
	model *viewmodels.LiveNotifierViewModel
	err   *viewmodels.ErrorViewModel
}

func (p *LiveNotifierPresenter) GetViewModel() interface{} {
	if p.err != nil {
		return p.err
	}

	return p.model
}

func (p *LiveNotifierPresenter) Error(e error) {
	p.err = &viewmodels.ErrorViewModel{Message: e.Error()}
}

func (p *LiveNotifierPresenter) Success(output *usecases.GetMediaInfosOutputMessage) {
	p.model = &viewmodels.LiveNotifierViewModel{
		StreamContainer: viewmodels.StreamContainer{
			Stream: output.Status,
		},
		VideosContainer: viewmodels.VideoContainer{
			Videos: viewmodels.Videos{
				Main:      output.LastVideoOnMainChannel,
				Secondary: output.LastVideoOnSecondaryChannel,
			},
		},
	}
}
