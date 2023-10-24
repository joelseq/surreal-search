<script>
	import '../app.postcss';
	import { AppBar, LightSwitch } from '@skeletonlabs/skeleton';
</script>

<AppBar>
	<svelte:fragment slot="lead"><span /></svelte:fragment>
	<svelte:fragment slot="trail"><LightSwitch rounded="rounded-full" /></svelte:fragment>
</AppBar>
<slot />
