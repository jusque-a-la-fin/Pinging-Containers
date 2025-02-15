import React, { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import TableForMany from './TableForMany';
import Preloader from './Preloader';
import { handleData, goToMainPage } from './utils';

const Data = () => {
    const location = useLocation();
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();
    const [data, setData] = useState(location.state?.data || []);

    useEffect(() => {
        setLoading(false);
    }, [data]);

      const fetchData = async () => {
        await handleData('update', setLoading, setData, navigate, false);
    };

    return (
      <div className="container mt-5 d-flex flex-column align-items-center">
        {loading ? (
          <Preloader />
        ) : (
          <>
          <button className="btn btn-red" onClick={() => goToMainPage(navigate)}>
            Назад
          </button>
          <div className="mb-3">
          <button className="btn btn-secondary mb-3" onClick={fetchData}>
          Обновить данные
          </button>
          </div>
            {data.length > 0 ? (
              <TableForMany data={data} />
            ) : null}
          </>
        )}
      </div>
    );
};

export default Data;
