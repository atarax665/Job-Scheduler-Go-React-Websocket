import React from "react";

const JobList = ({ jobs }) => {
  return (
    <div className="job-list bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
      <h2 className="text-2xl mb-4">Job List</h2>
      <ul className="list-none p-0">
        {jobs === null || jobs.length === 0 ? (
          <p className="text-red-500">No jobs scheduled</p>
        ) : (
          jobs.map((job) => (
            <li
              key={job.id}
              className={`job ${
                job.status
              } p-4 mb-4 border-2 rounded shadow ${getStatusClass(job.status)}`}
            >
              <div className="flex justify-between items-center">
                <span className="font-bold text-lg">Job Name : {job.name}</span>
                <span className="italic text-gray-600">
                  Status : {job.status}
                </span>
              </div>
              <div className="mt-2">
                <span>Duration : {job.duration / 1000000000} seconds</span>
              </div>
            </li>
          ))
        )}
      </ul>
    </div>
  );
};

const getStatusClass = (status) => {
  switch (status) {
    case "PENDING":
      return "bg-gray-100";
    case "RUNNING":
      return "bg-yellow-100";
    case "COMPLETED":
      return "bg-green-100";
    default:
      return "";
  }
};

export default JobList;
