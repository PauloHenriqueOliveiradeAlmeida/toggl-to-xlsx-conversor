package togglResponse

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"toggl_to_xlsx/httpclient"
	"toggl_to_xlsx/integration/toggl/responses"
)

type Toggl struct {
	Email    string
	Password string
}

func GetWorkspaces(toggl *Toggl) ([]togglResponse.WorkspaceResponse, error) {
	response, error := httpclient.Request[[]togglResponse.WorkspaceResponse]("https://api.track.toggl.com/api/v9/me/workspaces", http.MethodGet, nil, httpclient.Authentication{Email: toggl.Email, Password: toggl.Password})
	if error != nil {
		return response, fmt.Errorf("Error while get workspaces %s", error)
	}
	return response, nil
}

func GetTimeEntries(initialDate string, finalDate string, toggl *Toggl) ([]togglResponse.TimeEntriesResponse, error) {
	response, error := httpclient.Request[[]togglResponse.TimeEntriesResponse]("https://api.track.toggl.com/api/v9/me/time_entries?start_date="+initialDate+"&end_date="+finalDate, http.MethodGet, nil, httpclient.Authentication{Email: toggl.Email, Password: toggl.Password})

	if error != nil {
		return response, fmt.Errorf("Error while get time entries %s", error)
	}

	var timeEntries []togglResponse.TimeEntriesResponse
	for _, entry := range response {
		entry.Duration = entry.Stop.Sub(entry.Start)
		timeEntries = append(timeEntries, entry)
	}
	sort.Slice(timeEntries, func(i, j int) bool {
		return timeEntries[i].Start.Before(timeEntries[j].Stop)
	})
	return timeEntries, nil

}

func GetProjects(workspaceId int, toggl *Toggl) ([]togglResponse.ProjectResponse, error) {
	response, error := httpclient.Request[[]togglResponse.ProjectResponse]("https://api.track.toggl.com/api/v9/workspaces/"+strconv.Itoa(workspaceId)+"/projects", http.MethodGet, nil, httpclient.Authentication{Email: toggl.Email, Password: toggl.Password})
	if error != nil {
		return response, fmt.Errorf("Error while get time entries %s", error)
	}
	return response, nil

}
