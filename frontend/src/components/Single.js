import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Container.css';
import TableForSingle from './TableForSingle';
import Preloader from './Preloader';

const Container = ({ onSelectOption }) => {
  const [options, setOptions] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedOption, setSelectedOption] = useState(null);
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(false);

  const fetchOptions = async () => {
    try {
      const response = await fetch('/list/'); 
      const data = await response.json();
      setOptions(data);
    } catch (error) {
      console.error('Error fetching options:', error);
    }
  };

  const fetchdata = async (option) => {
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
    if (!showDropdown) {
      fetchOptions();
    }
    setShowDropdown(!showDropdown);
  };

  const handleOptionSelect = (option) => {
    setSelectedOption(option);
    setShowDropdown(false);
    fetchdata(option); 
    onSelectOption();
  };

  const handleRefreshPage = () => {
    window.location.reload(); 
  };
  
  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
      {!loading && !selectedOption && (
        <div className="mb-3">
          <button className="btn btn-primary" onClick={handleDropdownToggle}>
            Получить данные одного контейнера
           </button>
          {showDropdown && (
            <select onChange={(e) => handleOptionSelect(e.target.value)} defaultValue="">
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
      {loading && <Preloader />}
      {!loading && data && (
        <>
         <button className="btn-custom-red" onClick={handleRefreshPage}>
            Назад
          </button>
          <h2>Данные для контейнера "{selectedOption}":</h2>
          <h3>Дата последней успешной попытки: {data.SuccessPingTime}</h3>
          <button className="btn btn-secondary mb-3" onClick={() => fetchdata(selectedOption)}>
            Обновить данные
          </button>
          <TableForSingle data={data} />
        </>
      )}
    </div>
  );
};

export default Container;