<script lang="ts">
	import { userData } from '$lib/store/userStore';

	import { fly } from 'svelte/transition';
</script>

<svelte:head>
	<title>Home | Smacktalk Gaming</title>
</svelte:head>

<section in:fly={{ y: -100, duration: 500, delay: 500 }} out:fly={{ duration: 500 }}>
	<h1>
		<div class="welcome">
			<picture>
				<source srcset="svelte-welcome.webp" type="image/webp" />
				<img alt="Welcome" src="svelte-welcome.png" />
			</picture>
		</div>
		{#if $userData.username}
			<span style="text-transform: capitalize;">{$userData.username},</span>
		{/if}

		to our new<br />SvelteKit app
	</h1>
</section>