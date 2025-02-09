import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import DataTable from './components/DataTable';
import Preloader from './components/Preloader';

const App = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [button1Disabled, setButton1Disabled] = useState(false);
  const [button2Disabled, setButton2Disabled] = useState(true);

  const handleData = async (action) => {
    setLoading(true);
    if (action === 'fetch') {
      setButton1Disabled(true);
    }
    try {
      const response = await fetch('http://localhost:8080/logs', { method: 'GET' });
      const result = await response.json();
      setData(result);
      if (action === 'fetch') {
        setButton2Disabled(false);
      }
    } catch (error) {
      console.error(`Error ${action === 'fetch' ? 'fetching' : 'updating'} data:`, error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
      <h1 className="mb-4">Мониторинг Docker контейнеров</h1>
      {!loading && data.length === 0 && (
        <div className="mb-3">
          <button className="btn btn-primary" onClick={() => handleData('fetch')} disabled={button1Disabled}>
          Получить данные
          </button>
        </div>
      )}
      {loading && <Preloader />}
      {!loading && data.length > 0 && (
        <>
          <button className="btn btn-secondary mb-3" onClick={() => handleData('update')} disabled={button2Disabled}>
          Обновить данные
          </button>
          <DataTable data={data} />
        </>
      )}
    </div>
  );
};

export default App;
