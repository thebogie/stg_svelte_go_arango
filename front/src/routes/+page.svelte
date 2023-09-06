<script lang="ts">
    import {playerData} from '$lib/store/playerStore';

    import {fly} from 'svelte/transition';
</script>

<svelte:head>
    <title>Home | Smacktalk Gaming</title>
</svelte:head>

<section in:fly={{ y: -100, duration: 500, delay: 500 }} out:fly={{ duration: 500 }}>
    <h1>
        <div class="welcome">
            <picture>
                <source srcset="svelte-welcome.webp" type="image/webp"/>

            </picture>
        </div>
        {#if $playerData.email}
            <span style="text-transform: capitalize;">{$playerData.email},</span>
        {/if}

        to our new<br/>SvelteKit app
    </h1>
</section>