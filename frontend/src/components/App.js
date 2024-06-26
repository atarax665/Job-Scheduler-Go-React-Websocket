import React, { useEffect, useState } from "react";
import JobList from "./JobList";
import JobForm from "./JobForm";
import useWebSocket from "../hooks/useWebSocket";

const App = () => {
  const [jobs, setJobs] = useWebSocket(process.env.REACT_APP_WEBSOCKET_URL);

  useEffect(() => {
    // Function to fetch initial job list
    const fetchJobs = async () => {
      try {
        const response = await fetch(process.env.REACT_APP_BASE_URL);
        if (!response.ok) {
          throw new Error("Failed to fetch jobs");
        }
        const data = await response.json();
        setJobs(data);
      } catch (error) {
        console.error("Error fetching jobs:", error);
      }
    };

    fetchJobs();
  }, []);

  const addJob = async (newJob) => {
    try {
      const response = await fetch(process.env.REACT_APP_BASE_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newJob),
      });
      if (!response.ok) {
        throw new Error("Failed to add job");
      }
      const data = await response.json();

      if (jobs === null || jobs.length === 0) {
        setJobs([data]);
      } else if (data && data.name && data.duration && data.status) {
        setJobs((prevJobs) => [...prevJobs, data]);
      } else {
        console.error("Invalid job data received:", data);
      }
    } catch (error) {
      console.error("Error adding job:", error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Job Scheduler</h1>
      <JobForm addJob={addJob} />
      <JobList jobs={jobs} />
    </div>
  );
};

export default App;
