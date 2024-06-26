import { useEffect, useState } from "react";

const useWebSocket = (url, initialJobs = []) => {
  const [jobs, setJobs] = useState(initialJobs);
  const [ws, setWs] = useState(null);

  useEffect(() => {
    const webSocket = new WebSocket(url);

    webSocket.onopen = () => {
      console.log("WebSocket connected");
    };

    webSocket.onmessage = (event) => {
      const updateJob = JSON.parse(event.data);
      setJobs((prevJobs) => {
        if (prevJobs === null || prevJobs.length === 0) {
          return prevJobs;
        }
        const jobIndex = prevJobs.findIndex((job) => job.id === updateJob.id);
        if (jobIndex !== -1) {
          // Update existing job
          const updatedJobs = [...prevJobs];
          updatedJobs[jobIndex]["status"] = updateJob.status;
          return updatedJobs;
        } else {
          // Add new job
          return prevJobs;
        }
      });
    };

    webSocket.onclose = () => {
      console.log("WebSocket closed");
    };

    setWs(webSocket);

    // Cleanup WebSocket connection on component unmount
    return () => {
      webSocket.close();
    };
  }, [url]);

  return [jobs, setJobs];
};

export default useWebSocket;
