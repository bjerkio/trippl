package syncprojects

import (
	"fmt"
	"log"
	"strings"

	"github.com/bjerkio/trippl/internal/pkg/project"
	"github.com/bjerkio/trippl/internal/pkg/server"
	"github.com/bjerkio/trippl/internal/pkg/toggl"
	gtogglProject "github.com/dougEfresh/gtoggl-api/gtproject"
)

// func removeIndexFromProjects(s []project.Project, i int) []project.Project {
// 	if len(s) < 1 {
// 		return s
// 	}
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

func Run(opts server.Arguments) error {
	log.Println("Synchronizing Projects")

	// Get projects
	syncedProjects, tripletexProjects, err := project.GetProjects(opts)
	if err != nil {
		log.Panic(err)
		return err
	}

	togglProjects, err := toggl.GetProjects(&opts.TogglClient, opts.Config.TogglWorkspaceID)
	if err != nil {
		return err
	}

	var versionDigest string

	if tripletexProjects != nil {
		versionDigest = tripletexProjects.VersionDigest
	} else {
		versionDigest = syncedProjects.VersionDigest
	}

	cProjects := syncedProjects.Projects

	for i, sp := range cProjects {

		// Remove from synced if it's removed from Toggl. (lost sync)
		deletedFromToggl := true
		for _, p := range *togglProjects {
			if p.Id == sp.TogglId {
				deletedFromToggl = false
			}
		}

		if deletedFromToggl {
			log.Printf("«%s»(%d) was deleted from Toggl", sp.Name, i)
			versionDigest = "unknown"
			delete(cProjects, i)
		} else if tripletexProjects != nil {
			exists := false
			for _, p := range tripletexProjects.Values {
				if p.ID == sp.TripletexId {
					exists = true
				}
			}

			// If deleted, delete it from Toggl
			if !exists {
				log.Printf("«%s» does not exists anymore, deleting", sp.Name)
				err := opts.TogglClient.ProjectClient.Delete(sp.TogglId)

				if err != nil {
					return err
				}

				delete(cProjects, i)
			}
		}
	}

	if tripletexProjects != nil {
		log.Printf("new updated arrived from Tripletex (versionDigest: %s)", versionDigest)
		for _, p := range tripletexProjects.Values {

			displayName := strings.ReplaceAll(*&p.DisplayName, p.Number+" ", "")
			projectName := fmt.Sprintf("(%s) %s", p.Number, displayName)

			// Check if we have it already (check key-value)
			createIt := true
			updateIt := false
			var sP project.Project
			for i, sProject := range cProjects {
				if sProject.TripletexId == p.ID {
					if sProject.Name == projectName {
						createIt = false
					} else {
						sP = sProject
						sP.Name = projectName
						cProjects[i] = sP
						createIt = false
						updateIt = true
					}
				}
			}

			if createIt {
				log.Printf("«%s» is not created at Toggl, creating now", projectName)
				res, err := opts.TogglClient.ProjectClient.Create(&gtogglProject.Project{
					Name: projectName,
					WId:  opts.Config.TogglWorkspaceID,
				})

				if err != nil {
					return err
				}

				cProjects[p.ID] = project.Project{
					Name:        res.Name,
					TripletexId: p.ID,
					TogglId:     res.Id,
				}
			} else if updateIt {
				log.Printf("«%s» needs to be updated at Toggl, updating now", projectName)
				_, err := opts.TogglClient.ProjectClient.Update(&gtogglProject.Project{
					Id:   sP.TogglId,
					Name: projectName,
					WId:  opts.Config.TogglWorkspaceID,
				})

				if err != nil {
					return err
				}
			}
		}
	}

	err = opts.DB.SetObject([]byte("projects-sync"), &project.Projects{
		VersionDigest: versionDigest,
		Projects:      cProjects,
	})
	if err != nil {
		return err
	}

	log.Println("Changes are saved.")

	return nil
}
