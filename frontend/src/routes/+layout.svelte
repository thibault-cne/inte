<script lang="ts">
	import '../app.css';
	import '../app.postcss';

	import { AppShell } from '@skeletonlabs/skeleton';
</script>

<AppShell>
	<svelte:fragment slot="header">Header</svelte:fragment>
	<slot />
</AppShell>
