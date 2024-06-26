# Full Stack Job Scheduler Assignment - Go, React, and WebSockets

## Overview

This project is a simplified job scheduler with a React frontend that visualizes job statuses and allows users to submit new jobs. The scheduler prioritizes jobs using the Shortest Job First (SJF) algorithm. Real-time UI updates are achieved using WebSockets.

## Table of Contents

1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Design Choices](#design-choices)
6. [Models and structures used](#data-structures)
7. [Demo](#demo)

## Features

- **Go Backend**:

  - [x] Job representation with name and duration.
  - [x] SJF algorithm implementation for job prioritization.
  - [x] REST API to submit and retrieve jobs. (/jobs GET and POST)
  - [x] WebSocket server for real-time job status updates.

- **React Frontend**:
  - [x] Display job list with status indicators.
  - [x] Form to submit new jobs.
  - [x] WebSocket client for real-time UI updates.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [Node.js](https://nodejs.org/) (version 14 or higher) and npm
- [Git](https://git-scm.com/)

## Installation

### Backend (Go)

1. Clone the repository:

   ```sh
   git clone https://github.com/atarax665/Job-Scheduler-Go-React-Websocket.git
   cd Job-Scheduler-Go-React-Websocket/backend
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the backend server:
   ```sh
   make
   ```

### Frontend (React)

1. Navigate to the frontend directory:

   ```sh
   cd Job-Scheduler-Go-React-Websocket/frontend
   ```

2. Install dependencies:

   ```sh
   npm install
   ```

3. Run the frontend application:
   ```sh
   npm start
   ```

## Usage

### Submitting a New Job

- Use the form on the React frontend to submit a new job with a name and duration.
- The job will be added to the scheduler and prioritized based on the SJF algorithm.

### Viewing Job Statuses

- The React frontend displays the list of jobs with their statuses: pending, running, and completed.
- Job statuses are updated in real-time through WebSockets.

## Design Choices

### Backend (Go)

- **Job Representation**: Jobs are represented using a Go struct with fields for name and duration. I have added ID and Status field to the Job model to update status as the run proceeds. I am using an in-memory data structure `map[string]model.Job` to store jobs. Here I am using Job UUID as key in the map. I am using a map to make status update and specific job retrival simpler.
- **SJF Algorithm**: Implemented using Priority Queue (container/heap) data structures to prioritize jobs based on their duration. Implementation can be found in (`utils/priority_queue.go`)
- **REST API**: Provides endpoints to submit new jobs (`POST /jobs`) and retrieve the current list of jobs (`GET /jobs`).
- **WebSocket Server**: Uses the gorilla/websocket library to set up a WebSocket endpoint (`/ws`). Job status updates are broadcasted to all connected clients. The Clients connected to the backend are stored in a map[*Client]bool map and are deleted from the map once they disconnect. If this project is to be run in a distributed environment then we will need `Redis` to maintain the list of clients connected to the backend so that we can broadcast updates to them through websocket.
- **Error Handeling**: All the custom errors can be added to `backend/pkg/message/error.go`.

### Frontend (React)

- **Job Display**: A React component fetches and displays the list of jobs from the backend, with visual indicators for different job statuses.
- **Job Submission**: A simple form allows users to submit new jobs. The form data is sent to the backend via a POST request.
- **WebSocket Client**: Establishes a WebSocket connection to receive real-time job updates and update the UI accordingly. The websocket listener is implemented as a custom hook in `frontend/src/hooks/useWebSocket.js`.
- **Build Tool**: Used `parcel` to create optimal build of the frontend.

## Data Structures

## Job

```go
type Job struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Duration time.Duration        `json:"duration"`
	Status   job_status.JobStatus `json:"status"`
}

var JobsDataDb map[string]Job

```

## Status

```go
type Status struct {
	ID     string               `json:"id"`
	Name   string               `json:"name"`
	Status job_status.JobStatus `json:"status"`
}

```

Note - I have added .env in frontend folder to make it easier to run the project end to end without much effort. Though the best practice is to not include .env file in your project.

## Demo

### Homepage
<img src="https://github.com/atarax665/Job-Scheduler-Go-React-Websocket/blob/master/assets/homepage.png" alt="Homepage" width="800" height="500" />

### Working Gif
<img src="https://github.com/atarax665/Job-Scheduler-Go-React-Websocket/blob/master/assets/Demo.gif" alt="Homepage" width="800" height="500" />
Feel free to reach out if you have any questions !!
