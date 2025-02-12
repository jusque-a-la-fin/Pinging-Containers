import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Container from './components/Single';
import Containers from './components/Many';

const App = () => {
  const [showManyComponent, setManyComponent] = useState(true);
  const [showSingleComponent, setSingleComponent] = useState(true);

  const handleDropdownToggle = () => {
    setManyComponent(false); 
  };

  const handleFetchData = () => {
    setSingleComponent(false); 
  };

  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
       <h1 className="mb-4">Мониторинг Docker контейнеров</h1>
      {showManyComponent && <Containers onFetchData={handleFetchData} />}
      {showSingleComponent && <Container onSelectOption={handleDropdownToggle} />}
    </div>
  );
};

export default App;