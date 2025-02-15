import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home';
import Many from './components/Many';
import Data from './components/Data';
import Container from './components/Single';

const App = () => {
  return (
      <Router>
          <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/data/" element={<Many />} />
              <Route path="/containers/" element={<Data />} />
              <Route path="/container/:id" element={<Container />} />
          </Routes>
      </Router>
  );
};

export default App;
