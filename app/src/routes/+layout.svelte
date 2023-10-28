<script>
	import '../app.postcss';
	import { AppBar, LightSwitch } from '@skeletonlabs/skeleton';
	import Fa from 'svelte-fa';
	import { faGithub } from '@fortawesome/free-brands-svg-icons';
</script>

<AppBar>
	<svelte:fragment slot="lead"><span /></svelte:fragment>
	<svelte:fragment slot="trail">
		<a
			class="btn rounded-sm"
			href="https://github.com/joelseq/surreal-search"
			target="_blank"
			rel="noreferrer"
		>
			<Fa size="2x" icon={faGithub} class="mr-4" />
		</a>
		<LightSwitch rounded="rounded-full" />
	</svelte:fragment>
</AppBar>
<slot />
