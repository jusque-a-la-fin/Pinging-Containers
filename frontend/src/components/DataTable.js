import React from 'react';

const DataTable = ({ data }) => {
  return (
    <table className="table table-bordered mt-3">
      <thead className="thead-light">
        <tr>
          <th className="text-center"></th>
          <th className="text-center">IP адрес</th>
          <th className="text-center">Время пинга</th>
          <th className="text-center">Дата последней успешной попытки</th>
        </tr>
      </thead>
      <tbody>
        {data.map(container => (
          <tr key={container.ID}>
            <td>{container.ID}</td>
            <td>{container.IPv4}</td>
            <td>{container.PingDuration}</td>
            <td>{container.SuccessPingTime}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default DataTable;

