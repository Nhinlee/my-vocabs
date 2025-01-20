import React from 'react';

const SearchBar: React.FC = () => {
  return (
    <div style={{ display: 'flex', alignItems: 'center', padding: '1rem 2rem' }}>
      <input
        type="text"
        placeholder="Search"
        style={{ flex: 1, padding: '0.5rem', border: '1px solid #ddd', borderRadius: '4px' }}
      />
      <button style={{ marginLeft: '0.5rem', padding: '0.5rem 1rem', border: 'none', background: '#000', color: '#fff', cursor: 'pointer' }}>
        🔍
      </button>
    </div>
  );
};

export default SearchBar;
