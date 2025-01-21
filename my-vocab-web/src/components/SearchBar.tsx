import React, { useState } from 'react';

interface SearchBarProps {
  onSearch: (query: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [searchQuery, setSearchQuery] = useState<string>('');

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

  const handleSearch = () => {
    onSearch(searchQuery);
  };

  return (
    <div style={{ display: 'flex', alignItems: 'center', padding: '1rem 2rem' }}>
      <input
        type="text"
        placeholder="Search"
        value={searchQuery}
        onChange={handleInputChange}
        style={{ flex: 1, padding: '0.5rem', border: '1px solid #ddd', borderRadius: '4px' }}
      />
      <button
        onClick={handleSearch}
        style={{ marginLeft: '0.5rem', padding: '0.5rem 1rem', border: 'none', background: '#000', color: '#fff', cursor: 'pointer' }}
      >
        🔍
      </button>
    </div>
  );
};

export default SearchBar;