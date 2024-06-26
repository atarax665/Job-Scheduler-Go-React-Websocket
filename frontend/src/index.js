import React from "react";
import ReactDOM from "react-dom/client";
import App from "./Components/App";

const AppLayout = () => {
  return (
    <div>
      <App />
    </div>
  );
};

// create root using createRoot
const root = ReactDOM.createRoot(document.getElementById("root"));
// passing react element inside root
root.render(<AppLayout />);
