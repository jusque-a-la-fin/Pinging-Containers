import React from 'react';
import './Preloader.css'; 

const Preloader = () => {
  return (
    <div className="text-center">
      <div className="spinner-border" role="status">
      </div>
      <div className="loading-text">
        Loading
        <span className="dot">.</span>
        <span className="dot">.</span>
        <span className="dot">.</span>
      </div>
    </div>
  );
};

export default Preloader;

