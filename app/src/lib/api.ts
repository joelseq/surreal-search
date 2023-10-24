import type { SearchResult } from './types';
import { PUBLIC_API_URL } from '$env/static/public';

export async function getSearchResults(query: string): Promise<SearchResult[]> {
	const res = await fetch(`${PUBLIC_API_URL}/search?q=${query}`);
	const pages = await res.json();

	return pages;
}
