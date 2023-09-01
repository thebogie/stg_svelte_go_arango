<script lang="ts">
    // import { notificationData } from '$lib/store/notificationStore';
    // import { post, browserSet, browserGet } from '$lib/gql/requestGraphql';
    // import { goto } from '$app/navigation';
    // import { variables } from '$lib/utils/constants';
    import {fly} from 'svelte/transition';
    // import  { CustomError } from '$lib/interfaces/error.interface';
    import {changeText} from '$lib/helpers/buttonText';
    import {loginUser} from "$lib/gql/requestGraphql";

    let email = '';
    let password = '';
    let errors;
    // errors: Array<CustomError>;

    const handleLogin = async () => {
        console.log("FISH");
        const jsonData = await loginUser();
        console.log("jsonData" + JSON.stringify(jsonData));

        if (browserGet('refreshToken')) {
            localStorage.removeItem('refreshToken');
        }

        /*
        const [jsonRes, err] = await post(fetch, `${variables.BASE_API_URI}/login/`, {
            user: {
                email: email,
                password: password
            }
        });
        const response: UserResponse = jsonRes;

        if (err.length > 0) {
            errors = err;
        } else if (response.user) {
            if (response.user.tokens && response.user.tokens.refresh) {
                browserSet('refreshToken', response.user.tokens.refresh);
            }
            notificationData.update(() => 'Login successful...');
            await goto('/');
        }

         */


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