import React from 'react';

interface SearchBarProps {
  searchQuery: string; // Pass the query from the parent
  onSearch: (query: string) => void; // Callback to update query in the parent
}

const SearchBar: React.FC<SearchBarProps> = ({ searchQuery, onSearch }) => {
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    onSearch(e.target.value); // Update the parent state directly
  };

  return (
    <div style={{ display: 'flex', alignItems: 'center', padding: '1rem 2rem' }}>
      <input
        type="text"
        placeholder="Search"
        value={searchQuery} // Controlled by the parent component
        onChange={handleInputChange} // Update parent on every input change
        style={{ flex: 1, padding: '0.5rem', border: '1px solid #ddd', borderRadius: '4px' }}
      />
      <button
        onClick={() => onSearch(searchQuery)} // Trigger search explicitly if needed
        style={{ marginLeft: '0.5rem', padding: '0.5rem 1rem', border: 'none', background: '#000', color: '#fff', cursor: 'pointer' }}
      >
        🔍
      </button>
    </div>
  );
};

export default SearchBar;