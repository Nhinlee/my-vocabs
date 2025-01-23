import axios from 'axios';
import { Word } from '../models/Word';

const API_BASE_URL = 'https://my-vocabs-821554367069.us-central1.run.app/api/v1/vocabs';

// Define a type for the filter options
interface VocabFilters {
  word?: string;           // Optional word filter
  categories?: string[];   // Optional array of categories
}

export const getVocabList = async (filters: VocabFilters = {}): Promise<Word[]> => {
  try {
    // Convert filters into query parameters
    const params = new URLSearchParams();

    if (filters.word) {
      params.append('word', filters.word);
    }

    if (filters.categories && filters.categories.length > 0) {
      filters.categories.forEach((category) => params.append('categories', category));
    }

    // Make the GET request with query parameters
    const response = await axios.get(`${API_BASE_URL}/list`, {
      params,
    });

    return response.data as Word[];
  } catch (error) {
    console.error('Error fetching vocab list:', error);
    throw error;
  }
};

export const getWordDetail = async (id: string): Promise<Word> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/${id}`);
    return response.data as Word;
  } catch (error) {
    console.error('Error fetching word details:', error);
    throw error;
  }
}
