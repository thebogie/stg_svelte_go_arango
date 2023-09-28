<script lang="ts">
    import {page} from '$app/stores';
    import {userData} from '$lib/store/userStore';
    import {logOutUser} from '$lib/gql/requestGraphql';
    import {playerData} from "$lib/store/playerStore";
    import type {IPlayer} from "$lib/interfaces/player.interface";


 console.log("HEADER: Playerdata?" + JSON.stringify({$playerData}));

</script>

<header>
    <div class="corner">
        <a href="https://kit.svelte.dev">
            <!--  <img alt="SvelteKit" src={logo} /> -->
        </a>
    </div>

    <nav>
        <svg aria-hidden="true" viewBox="0 0 2 3">
            <path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z"/>
        </svg>
        <ul>
            <li class:active={$page.url.pathname === '/'}>
                <a href="/">Home</a>
            </li>
            {#if !$userData.username}
                <li class:active={$page.url.pathname === '/accounts/login'}>
                    <a href="/accounts/login">Login</a>
                </li>
               <!--  <li class:active={$page.url.pathname === '/accounts/register'}>
                    <a href="/accounts/register">Register</a>
                </li>
                -->
            {:else}
                <li>
                    Hi, <a href="/accounts/user/{$userData.username}-{$userData.id}"
                >{$userData.username}</a
                >
                </li>
                <li>
                    <a href={null} on:click={logOutUser} style="cursor: pointer;">Logout</a>
                </li>
            {/if}
        </ul>
        <svg aria-hidden="true" viewBox="0 0 2 3">
            <path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z"/>
        </svg>
    </nav>

    <div class="corner">
        <a href="https://github.com/Sirneij/">
            <!-- <img alt="John O. Idogun" src={john} /> -->
        </a>
    </div>
</header>