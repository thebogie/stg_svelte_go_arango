<script lang="ts">
    // import { notificationData } from '$lib/store/notificationStore';
    import {browserGet, browserSet} from '$lib/utils/requestUtils';
    // import { goto } from '$app/navigation';
    // import { variables } from '$lib/utils/constants';
    import {fly} from 'svelte/transition';
    // import  { CustomError } from '$lib/interfaces/error.interface';
    import {changeText} from '$lib/helpers/buttonText';
    import {loginUser} from "$lib/gql/requestGraphql";
    //import type IPlayerLogin from "$lib/interfaces/player.interface";
    import type {CustomError} from "$lib/interfaces/error.interface";
    import {notificationData} from "$lib/store/notificationStore";
    import {goto} from "$app/navigation";
    import type {IPlayer, IPlayerLogin} from "$lib/interfaces/player.interface";
    import {playerData} from "$lib/store/playerStore";
    //import type {IPlayer} from "$lib/interfaces/player.interface";

    let email = '';
    let password = '';
    let errors: CustomError[] = [];
    // errors: Array<CustomError>;

    const handleLogin = async () => {


        if (browserGet('refreshToken')) {
            localStorage.removeItem('refreshToken');
        }

        let responseuserdata: IPlayer = {};
        let response: IPlayerLogin = {
            token: "",
            userdata: responseuserdata,
        };
        let err: CustomError[] = [];
        const jsonData = await loginUser(email, password);
        // console.log("jsonData" + JSON.stringify(jsonData));
        [response, err] = jsonData;

        console.log("HERE" + JSON.stringify(response));


        if (err.length > 0) {
            errors = err;
        } else if (response.userdata) {
            if (response.token) {
                browserSet('refreshToken', response.token);
            }
            notificationData.update(() => 'Login successful...');
            playerData.set(() => response.userdata)
            await goto('/');
        }


    };
</script>

<svelte:head>
    <title>Login | Smacktalk Gaming</title>
</svelte:head>

<section
        class="container"
        in:fly={{ x: -100, duration: 500, delay: 500 }}
        out:fly={{ duration: 500 }}
>
    <h1>Login</h1>
    {#if errors}
        {#each errors as error}
            <p class="center error">{error.error}</p>
        {/each}
    {/if}
    <form class="form" on:submit|preventDefault={handleLogin}>
        <input
                aria-label="Email address"
                bind:value={email}
                name="email"
                placeholder="Email address"
                required
                type="email"
        />
        <input
                aria-label="password"
                bind:value={password}
                name="password"
                placeholder="password"
                required
                type="password"
        />
        <button class="btn" on:click={(e) => changeText(e, 'Signing in...')} type="submit">Login
        </button
        >
        <!-- <p class="center">No account yet? <a href="/accounts/register">Get started</a>.</p> -->
    </form>
</section>