package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Job struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JobStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func execute(job Job) JobStatus {
	log.Println("executing ", job.Name)
	return JobStatus{job.Id, "SUCCESS"}
}

type agent struct{}

func (a agent) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		data, _ := io.ReadAll(request.Body)
		var jobs []Job
		_ = json.Unmarshal(data, &jobs)
		ch := make(chan JobStatus)
		for _, job := range jobs {
			go func(job Job) {
				ch <- execute(job)
			}(job)
		}

		var jobStatus []JobStatus
		for range jobs {
			jobStatus = append(jobStatus, <-ch)
		}
		out, _ := json.Marshal(jobStatus)
		writer.WriteHeader(200)
		_, _ = writer.Write(out)
	}
}

func main() {
	http.Handle("/jobs", &agent{})
	_ = http.ListenAndServe("localhost:8080", nil)
}
