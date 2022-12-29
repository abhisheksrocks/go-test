package main

import (
	// "encoding/json"

	// "io"
	"log"
	"net/http"
	// "os"
	// "time"
	// "github.com/joho/godotenv"
	// "github.com/abhisheksrocks/github-graphql-test/services"
)

// const GithubGraphQLurl = "https://api.github.com/graphql"

func main() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic(err)
	// }
	// githubAuthToken := os.Getenv("githubAuthToken")
	// var _ services.RequestModel = services.NewRequestModel(
	// 	`
	// 					{
	// 						user(login: "abhisheksrocks") {
	// 							repositories(
	// 								orderBy: {field: PUSHED_AT, direction: DESC}
	// 								first: 10
	// 								privacy: PUBLIC
	// 							) {
	// 								nodes {
	// 									name
	// 									url
	// 									isPrivate
	// 									isArchived
	// 									parent {
	// 										url
	// 									}
	// 									description
	// 									languages(first: 1, orderBy: {field: SIZE, direction: DESC}) {
	// 										nodes {
	// 											name
	// 											color
	// 										}
	// 									}
	// 									stargazerCount
	// 									forkCount
	// 								}
	// 							}
	// 						}
	// 					}`,
	// 	map[string]string{
	// 		"Authorization": "bearer " + githubAuthToken,
	// 	},
	// )

	// var requestDataObj services.RequestModel = services.NewRequestModel(
	// 	`
	// 					{
	// 						repository(name: "pdf_viewer", owner: "abhisheksrocks") {
	// 							name
	// 							url
	// 							isPrivate
	// 							isArchived
	// 							parent {
	// 								url
	// 							}
	// 							description
	// 							languages(first: 1, orderBy: {field: SIZE, direction: DESC}) {
	// 								nodes {
	// 									name
	// 									color
	// 								}
	// 							}
	// 							stargazerCount
	// 							forkCount
	// 						}
	// 					}
	// 	`, map[string]string{
	// 		"Authorization": "bearer " + githubAuthToken,
	// 	},
	// )

	// var _ services.RequestModel = services.NewRequestModel(
	// 	`
	// 		{
	// 			viewer{
	// 				login
	// 			}
	// 		}
	// 	`,
	// 	map[string]string{
	// 		"Authorization": "bearer " + githubAuthToken,
	// 	},
	// )

	// requestObj, err := requestDataObj.CreateRequestObj(GithubGraphQLurl)
	// if err != nil {
	// 	panic(err)
	// }

	// client := http.Client{
	// 	Timeout: 15 * time.Second,
	// }

	// responseObj, err := client.Do(requestObj)

	// if err != nil {
	// 	panic(err)
	// }
	// defer responseObj.Body.Close()

	// result, err := io.ReadAll(responseObj.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// // var data = struct {
	// // 	Data struct {
	// // 		Viewer struct {
	// // 			Login string
	// // 		}
	// // 	}
	// // }{}

	// var data = struct {
	// 	Data struct {
	// 		Repository struct {
	// 			Name        string
	// 			Description string
	// 			Url         string
	// 			IsPrivate   bool
	// 			IsArchived  bool
	// 			Parent      struct {
	// 				Url string
	// 			}
	// 			Languages struct {
	// 				Nodes []struct {
	// 					Name  string
	// 					Color string
	// 				}
	// 			}
	// 			StargazerCount int
	// 			ForkCount      int
	// 		}
	// 	}
	// }{}

	// if err := json.Unmarshal(result, &data); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Name:", data.Data.Repository.Name)
	// fmt.Println("URL:", data.Data.Repository.Url)
	// fmt.Println("Parent URL:", data.Data.Repository.Parent.Url)
	// fmt.Println("Description:", data.Data.Repository.Description)
	// fmt.Println("Primary Language:", data.Data.Repository.Languages.Nodes[0])
	// fmt.Println("Star Count:", data.Data.Repository.StargazerCount)
	// fmt.Println("Fork Count:", data.Data.Repository.ForkCount)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
		http.ServeFile(w, r, "test3.svg")
	})

	mux.HandleFunc("/mdn_logo_only_color.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png; charset=utf-8")
		http.ServeFile(w, r, "https://go-test-s68v.onrender.com/mdn_logo_only_color.png")
	})

	// go func() {

	// 	// log.Fatal(http.ListenAndServe(":8080", nil))
	// }()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
	println("Started")

}
