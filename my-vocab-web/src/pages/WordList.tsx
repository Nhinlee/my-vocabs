import React, { useEffect, useState } from "react";
import Header from "../components/Header";
import SearchBar from "../components/SearchBar";
import WordCard from "../components/WordCard";
import { getVocabList } from "../services/vocabService";
import { Word } from "../models/Word";

const WordList: React.FC = () => {
  const [words, setWords] = useState<Word[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState<string>(""); // User's input
  const [debouncedQuery, setDebouncedQuery] = useState<string>(""); // Debounced value

  const fetchVocabList = async (word: string) => {
    try {
      setLoading(true);
      const vocabList = await getVocabList({ word });
      setWords(vocabList);
    } catch (err) {
      setError("Failed to fetch vocab list. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  // Debounce searchQuery updates
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedQuery(searchQuery);
    }, 500); // 500ms debounce time

    return () => clearTimeout(timer); // Cleanup the timer on unmount or input change
  }, [searchQuery]);

  // Fetch data when the debounced query changes
  useEffect(() => {
    fetchVocabList(debouncedQuery);
  }, [debouncedQuery]);

  if (loading) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>{error}</p>;
  }

  return (
    <div>
      <Header />
      <SearchBar searchQuery={searchQuery} onSearch={setSearchQuery} />
      {/* <FilterBar /> */}
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(auto-fit, minmax(200px, 1fr))",
          gap: "1rem",
          padding: "2rem",
        }}
      >
        {words.map((word) => (
          <WordCard
            key={word.vocab_id}
            vocabId={word.vocab_id}
            word={word.word}
            nextReview={word.next_review}
          />
        ))}
      </div>
    </div>
  );
};

export default WordList;
