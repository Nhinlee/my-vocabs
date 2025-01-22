import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import WordList from './pages/WordList';
import WordDetail from './pages/WordDetail';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<WordList />} />
        <Route
          path="/word/:id"
          element={
            <WordDetail/>
          }
        />
      </Routes>
    </Router>
  );
};

export default App;
