import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Container.css';
import TableForSingle from './TableForSingle';
import Preloader from './Preloader';
import { goToMainPage } from './utils';

const Container = () => {
  const { id } = useParams();
  const [options, setOptions] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedOption, setSelectedOption] = useState(null);
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    fetchOptions();
  }, []);

  useEffect(() => {
    if (options.length > 0) {
      const option = options[id];
      if (option) {
        setSelectedOption(option);
        fetchData(option);
      }
    }
  }, [id, options]);

  const fetchOptions = async () => {
    try {
      const response = await fetch('/list/'); 
      const data = await response.json();
      setOptions(data);
    } catch (error) {
      console.error('Error fetching options:', error);
    }
  };

  const fetchData = async (option) => {
    setLoading(true);
    try {
      const response = await fetch(`/container/${encodeURIComponent(option)}`); 
      const data = await response.json();
      setData(data); 
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleDropdownToggle = () => {
    setShowDropdown(!showDropdown);
  };


  const handleOptionSelect = (option, index) => {
    setSelectedOption(option);
    setShowDropdown(false);
    fetchData(option);
    navigate(`/container/${index}`);
  };

  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
       {loading && <Preloader />}
      {!loading && !selectedOption && (
        <div className="mb-3">
          <button className="btn btn-primary" onClick={handleDropdownToggle}>
            Получить данные одного контейнера
           </button>
          {showDropdown && (
            <select onChange={(e) => handleOptionSelect(e.target.value, e.target.selectedIndex)} defaultValue="">
              <option value="" disabled>Выбрать контейнер</option>
              {options.map((option, index) => (
                <option key={index} value={option}>
                  {option}
                </option>
              ))}
            </select>
          )}
        </div>
      )}
      { data && (
        <>
         <button className="btn btn-red" onClick={() => goToMainPage(navigate)}>
            Назад
          </button>

          <h2>Данные для контейнера "{selectedOption}":</h2>
          <h3>Дата последней успешной попытки: {data.SuccessPingTime}</h3>
          <button className="btn btn-secondary mb-3" onClick={() => fetchData(selectedOption)}>
            Обновить данные
          </button>
          <TableForSingle data={data} />
        </>
      )}
    </div>
  );
};

export default Container;