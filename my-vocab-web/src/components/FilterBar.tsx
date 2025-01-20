import React from 'react';

const FilterBar: React.FC = () => {
  return (
    <div style={{ display: 'flex', gap: '1rem', padding: '1rem 2rem' }}>
      <button style={{ border: '1px solid #000', padding: '0.5rem 1rem', background: '#000', color: '#fff', cursor: 'pointer' }}>New</button>
      <button style={{ border: '1px solid #ddd', padding: '0.5rem 1rem', background: 'transparent', cursor: 'pointer' }}>Price ascending</button>
      <button style={{ border: '1px solid #ddd', padding: '0.5rem 1rem', background: 'transparent', cursor: 'pointer' }}>Price descending</button>
      <button style={{ border: '1px solid #ddd', padding: '0.5rem 1rem', background: 'transparent', cursor: 'pointer' }}>Rating</button>
    </div>
  );
};

export default FilterBar;
