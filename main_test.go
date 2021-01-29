package pullRequests_milestones_reminder

import (
	"github.com/bnn-tk/pullRequests-milestones-reminde/repository"
	"testing"
)

func TestFetchAllMileStones(t *testing.T) {
	// TODO: API叩くところはmock可したい
	repo := repository.NewFetchMileStonesHttpClientRepository()
	mileStones, err := repo.FetchAllMileStones()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if len(mileStones) != 1 {
		t.Errorf("mileStones length is not 1 got: %b", len(mileStones))
	}

	t.Log("Success")
}
