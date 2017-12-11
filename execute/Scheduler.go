package execute

import (
	"github.com/hashicorp/nomad/api"
	"github.com/golang-plus/errors"
	"github.com/dataprism/dataprism-commons/utils"
)

type NomadScheduler struct {
	nomad *api.Client
	jobsDir string
}

type DataprismJob interface {

	ToJob() (*api.Job, error)

}

type DataprismJobStatus struct {
	Queued int `json:"queued"`
	Complete int `json:"complete"`
	Failed int `json:"failed"`
	Running int `json:"running"`
	Starting int `json:"starting"`
	Lost int `json:"lost"`
}

type ScheduleResponse struct {
	EvalId string `json:"eval_id"`
}

type UnscheduleResponse struct {
	EvalId string `json:"eval_id"`
}

func NewNomadScheduler(nomad *api.Client, jobsDir string) (*NomadScheduler) {
	return &NomadScheduler{nomad: nomad, jobsDir: jobsDir}
}

func (s *NomadScheduler) Schedule(job DataprismJob) (*ScheduleResponse, error) {
	nomadJob, err := job.ToJob()

	if err != nil {
		return nil, err
	}

	resp, _, err := s.nomad.Jobs().Register(nomadJob, &api.WriteOptions{})

	if err != nil {
		return nil, err
	}


	return &ScheduleResponse{resp.EvalID}, nil
}

func (s *NomadScheduler) Unschedule(kind string, id string) (*UnscheduleResponse, error) {
	nomadJobId := utils.ToNomadJobId(kind, id)

	res, _, err := s.nomad.Jobs().Deregister(nomadJobId, true, &api.WriteOptions{})

	if err != nil {
		return nil, err
	} else {
		return &UnscheduleResponse{EvalId: res}, nil
	}
}

func (m *NomadScheduler) GetJobStatus(kind string, id string) (*DataprismJobStatus, error) {
	summary, _, err := m.nomad.Jobs().Summary(id, &api.QueryOptions{})
	if err != nil {
		return &DataprismJobStatus{
			Complete: 0,
			Failed:   0,
			Lost:     0,
			Queued:   0,
			Running:  0,
			Starting: 0,
		}, nil
	}

	tgSummary, exists := summary.Summary[kind + "s"]

	if exists {
		return &DataprismJobStatus{
			Complete: tgSummary.Complete,
			Failed:   tgSummary.Failed,
			Lost:     tgSummary.Lost,
			Queued:   tgSummary.Queued,
			Running:  tgSummary.Running,
			Starting: tgSummary.Starting,
		}, nil
	} else {
		return nil, errors.New("no " + kind + "s element found in the job summary")
	}
}