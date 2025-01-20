import axios from 'axios';
import { Word } from '../models/Word';

const API_BASE_URL = 'http://localhost:8080/api/v1/vocabs';

export const getVocabList = async (): Promise<Word[]> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/list`);
    return response.data as Word[];
  } catch (error) {
    console.error('Error fetching vocab list:', error);
    throw error;
  }
};
