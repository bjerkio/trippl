package project

import (
	"net/http"
	"regexp"
	"time"

	"github.com/bjerkio/tripletex-go/client/project"
	"github.com/bjerkio/tripletex-go/models"
	"github.com/bjerkio/trippl/internal/pkg/server"
)

func setNoneMatch(inner http.RoundTripper, versionDigest string) http.RoundTripper {
	return &addNM{
		inner:         inner,
		VersionDigest: versionDigest,
	}
}

type addNM struct {
	inner         http.RoundTripper
	VersionDigest string
}

func (nm *addNM) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("If-None-Match", nm.VersionDigest)
	resp, err := nm.inner.RoundTrip(r)
	if resp.StatusCode == 304 {
		return resp, nil
	}
	return resp, err
}

func GetProjects(opts server.Arguments) (*Projects, *models.ListResponseProject, error) {
	// 1. Get sync-key-value (where we store checksum and ids)
	syncedProjects := Projects{
		VersionDigest: "unknown",
		Projects:      map[int32]Project{},
	}
	err := opts.DB.GetObject([]byte("projects-sync"), &syncedProjects)
	if err != nil {
		if err.Error() != "Key not found" {
			return nil, nil, err
		}
	}

	// 2. Get all projects from Tripletex
	projectQuery := project.NewProjectSearchParams()
	projectQuery.HTTPClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: setNoneMatch(http.DefaultTransport, syncedProjects.VersionDigest),
	}
	falsy := false
	projectQuery.IsClosed = &falsy
	projectQuery.IsOffer = &falsy

	res, err := opts.TripletexClient.Project.ProjectSearch(projectQuery, opts.TripletexAuth)

	if err != nil {
		re := regexp.MustCompile(`(?m)\(status 304\)`)
		if re.Match([]byte(err.Error())) {
			// There are no changes
			return &syncedProjects, nil, nil
		} else {
			return nil, nil, err
		}
	}

	return &syncedProjects, res.Payload, nil
}
