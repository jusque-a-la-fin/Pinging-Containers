import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Container.css';
import Preloader from './Preloader';
import { handleData } from './utils';

const Many = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [hasNavigated, setHasNavigated] = useState(false);
  const navigate = useNavigate();

  
  const fetchData = async () => {
    await handleData('fetch', setLoading, setData, navigate, !hasNavigated);
    if (!hasNavigated) {
        setHasNavigated(true);
    }
  };

  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
      {!loading && data.length === 0 && (
        <div className="mb-3">
          <button className="btn btn-primary" onClick={fetchData}>
          Получить данные всех контейнеров
          </button>
        </div>
      )}
      {loading && <Preloader />}
    </div>
  );
};

export default Many;