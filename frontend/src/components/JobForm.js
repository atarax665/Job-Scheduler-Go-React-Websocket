import React, { useState } from "react";

const JobForm = ({ addJob }) => {
  const [name, setName] = useState("");
  const [duration, setDuration] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    const newJob = {
      name,
      duration: parseInt(duration, 10),
    };
    addJob(newJob);
    setName("");
    setDuration("");
  };

  return (
    <div className="job-form mb-4">
      <h2 className="text-xl font-semibold mb-2">Add Job</h2>
      <form onSubmit={handleSubmit} className="flex flex-col space-y-2">
        <label className="flex flex-col">
          Name:
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="p-2 border rounded"
            required
          />
        </label>
        <label className="flex flex-col">
          Duration (seconds):
          <input
            type="number"
            value={duration}
            onChange={(e) => setDuration(e.target.value)}
            className="p-2 border rounded"
            required
          />
        </label>
        <button type="submit" className="p-2 bg-blue-500 text-white rounded">
          Add Job
        </button>
      </form>
    </div>
  );
};

export default JobForm;
