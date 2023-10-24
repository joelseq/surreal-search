<script lang="ts">
	import debounce from 'lodash/debounce';
	import { getSearchResults } from '$lib/api';
	import Result from '$lib/components/Result.svelte';
	import type { SearchResult, FormEventHandler } from '$lib/types';

	let searchResults: SearchResult[] | null = null;

	const handleInput: FormEventHandler<HTMLInputElement> = async (event) => {
		const query = (event.target as HTMLInputElement).value;

		if (query === '' || query == null) {
			searchResults = null;
		} else {
			const results = await getSearchResults(query);
			searchResults = results;
		}
	};

	const handleInputDebounced = debounce(handleInput, 200);
</script>

<div class="flex h-full justify-center overflow-y-scroll py-16">
	<div class="container">
		<h1 class="h1 mb-4">Surreal Search</h1>
		<p class="text-xl mb-2">An unofficial search engine for the SurrealDB docs</p>
		<label class="label">
			<input
				class="input rounded"
				type="text"
				placeholder="Search"
				on:input={handleInputDebounced}
			/>
		</label>
		{#if searchResults != null}
			<p class="text-surface-500-400-token my-4">{searchResults.length} results</p>
			<div class="border border-surface-600 rounded">
				{#each searchResults as result, i}
					<div>
						{#if i !== 0}
							<hr />
						{/if}
						<Result {result} />
					</div>
				{/each}
			</div>
		{/if}
		<div class="h-16" />
	</div>
</div>
