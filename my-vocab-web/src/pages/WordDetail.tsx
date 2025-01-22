import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getWordDetail } from '../services/vocabService';

interface WordDetailProps {
  word: string;
  imageUrl?: string;
  subheading?: string;
  description?: string;
}

const WordDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [wordDetail, setWordDetail] = useState<WordDetailProps | null>(null);

  useEffect(() => {
    if (!id) {
      console.error('No ID provided');
      return;
    }
    getWordDetail(id).then((data) => setWordDetail(data));
  }, [id]);

  if (!wordDetail) {
    return <p>Loading...</p>;
  }

  const { word, imageUrl, subheading, description } = wordDetail;

  return (
    <div style={{ padding: '2rem', display: 'flex', gap: '2rem', alignItems: 'flex-start' }}>
      <div style={{ flex: '1 1 auto', maxWidth: '400px' }}>
        {imageUrl ? (
          <img
            src={imageUrl}
            alt={word}
            style={{ width: '100%', borderRadius: '8px', objectFit: 'cover' }}
          />
        ) : (
          <div
            style={{
              width: '100%',
              height: '300px',
              backgroundColor: '#f0f0f0',
              borderRadius: '8px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              fontSize: '1.5rem',
              color: '#aaa',
            }}
          >
            No Image
          </div>
        )}
      </div>

      <div style={{ flex: '2 1 auto' }}>
        <h1 style={{ marginBottom: '0.5rem' }}>{word}</h1>
        {subheading && <h2 style={{ margin: '0 0 1rem', color: '#666' }}>{subheading}</h2>}
        {description && <p style={{ lineHeight: '1.6', color: '#333' }}>{description}</p>}
      </div>
    </div>
  );
};

export default WordDetail;