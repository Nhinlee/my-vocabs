export interface Word {
    vocab_id: string;
    word: string;
    image_urls: string[] | null;
    next_review: string;
    reviewed_time: number;
    created_at: string;
    updated_at: string;
    deleted_at: string | null;
  }