<script lang="ts">
	import type { SearchResult } from '$lib/types';

	export let result: SearchResult;
</script>

<a href={result.page.url} target="_blank">
	<div class="p-4 hover:bg-surface-300-600-token">
		<p class="text-tertiary-900-50-token">{result.page.breadcrumb}</p>
		<p class="text-2xl my-2">{result.page.title}</p>
		<!-- eslint-disable-next-line svelte/no-at-html-tags -->
		<p>{@html result.highlight}</p>
	</div>
</a>
