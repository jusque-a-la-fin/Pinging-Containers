import React from 'react';

const TableForSingle = ({ data }) => {
  const { PingDurations } = data;
  const maxLength = Math.max(PingDurations.length);

  return (
    <table className="table table-bordered mt-3">
      <thead className="thead-light">
        <tr>
          <th className="text-center"></th>
          <th className="text-center">Время пинга</th>
        </tr>
      </thead>
      <tbody>
        {Array.from({ length: maxLength }).map((_, index) => (
          <tr key={index}>
            <td>{index + 1}</td>
            {PingDurations[index] && <td>{PingDurations[index]}</td>}
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default TableForSingle;
