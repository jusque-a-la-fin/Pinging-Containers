import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Container.css';
import TableForMany from './TableForMany';
import Preloader from './Preloader';

const Containers = ({ onFetchData }) => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);

  const handleData = async (action) => {
    setLoading(true);
    if (action === 'fetch') {
      onFetchData();
    }
    try {
      const response = await fetch('/get/', { method: 'GET' });
      const result = await response.json();
      setData(result);
    } catch (error) {
      console.error(`Error ${action === 'fetch' ? 'fetching' : 'updating'} data:`, error);
    } finally {
      setLoading(false);
    }
  };


  const handleRefreshPage = () => {
    window.location.reload(); 
  };

  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
      {!loading && data.length === 0 && (
        <div className="mb-3">
          <button className="btn btn-primary" onClick={() => handleData('fetch')}>
          Получить данные всех контейнеров
          </button>
        </div>
      )}
      {!loading && data.length > 0 && (
        <>
          <button className="btn-custom-red" onClick={handleRefreshPage}>
            Назад
          </button>
          <button className="btn btn-secondary mb-3" onClick={() => handleData('update')}>
          Обновить данные
          </button>
          <TableForMany data={data} />
        </>
      )}
      {loading && <Preloader />}
    </div>
  );
};

export default Containers;