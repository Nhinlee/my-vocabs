import React from 'react';
import { useNavigate } from 'react-router-dom';

interface WordCardProps {
  word: string;
  vocabId: string;
  nextReview: string;
}

const WordCard: React.FC<WordCardProps> = ({ word, vocabId }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/word/${vocabId}`);
  };

  return (
    <div
      style={{
        border: '1px solid #ddd',
        borderRadius: '8px',
        padding: '1rem',
        textAlign: 'center',
        cursor: 'pointer',
      }}
      onClick={handleClick}
    >
      <div
        style={{
          width: '100%',
          height: '120px',
          background: '#f0f0f0',
          borderRadius: '4px',
          marginBottom: '1rem',
        }}
      ></div>
      <p>{word}</p>
    </div>
  );
};

export default WordCard;
