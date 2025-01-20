import React from 'react';
import Header from '../components/Header';
import SearchBar from '../components/SearchBar';
import FilterBar from '../components/FilterBar';
import WordCard from '../components/WordCard';

const words = ['Retail', 'Retail', 'Retail', 'Retail', 'Retail', 'Retail', 'Retail', 'Retail'];

const WordList: React.FC = () => {
  return (
    <div>
      <Header />
      <SearchBar />
      {/* <FilterBar /> */}
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))', gap: '1rem', padding: '2rem' }}>
        {words.map((word, index) => (
          <WordCard key={index} word={word} />
        ))}
      </div>
    </div>
  );
};

export default WordList;
