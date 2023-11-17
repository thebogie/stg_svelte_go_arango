<script lang="ts">
    import type {ActionData} from './$types';
    import {enhance} from '$app/forms';


    console.log("login page.svelte");
    export let form: ActionData;
</script>

<svelte:head>
    <title>Login</title>
</svelte:head>

<h1>Login</h1>

<section>
    <form method="post" use:enhance>
        <div class="group">
            <label for="email">Email</label>
            <input id="email" name="email" required type="email"/>
        </div>

        <div class="group">
            <label for="password">Password</label>
            <input id="password" name="password" required type="password"/>
        </div>

        <div class="submit-container">
            <button type="submit">Login</button>
        </div>

        {#if form?.error}
            <div class="notice error">
                {form.error}
            </div>
        {/if}
    </form>

    <div class="actions">
        <a href="/signup">Sign Up</a>
    </div>
</section>