import React from 'react';

interface WordCardProps {
  word: string;
  nextReview?: string;
}

const WordCard: React.FC<WordCardProps> = ({ word, nextReview }) => {
  return (
    <div style={{ border: '1px solid #ddd', borderRadius: '8px', padding: '1rem', textAlign: 'center' }}>
      <div style={{ width: '100%', height: '120px', background: '#f0f0f0', borderRadius: '4px', marginBottom: '1rem' }}></div>
      <p>{word}</p>
      {nextReview && <small>Next Review: {new Date(nextReview).toLocaleDateString()}</small>}
    </div>
  );
};

export default WordCard;
