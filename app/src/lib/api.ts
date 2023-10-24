import type { SearchResult } from './types';

export async function getSearchResults(query: string): Promise<SearchResult[]> {
	const res = await fetch(`http://localhost:8080/search?q=${query}`);
	const pages = await res.json();

	return pages;
}
