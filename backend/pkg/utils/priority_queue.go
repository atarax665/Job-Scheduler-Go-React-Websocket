package utils

import "scheduler-service/pkg/model"

type JobPriorityQueue []*model.Job

func (pq JobPriorityQueue) Len() int { return len(pq) }

func (pq JobPriorityQueue) Less(i, j int) bool {
	return pq[i].Duration < pq[j].Duration
}

func (pq JobPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *JobPriorityQueue) Push(x interface{}) {
	job := x.(*model.Job)
	*pq = append(*pq, job)
}

func (pq *JobPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	job := old[n-1]
	*pq = old[0 : n-1]
	return job
}
