<script lang="ts">
    //import type {ActionData} from './$types';
    // import {enhance} from '$app/forms';

    console.log("login page.svelte");

</script>

<svelte:head>
    <title>Login</title>
</svelte:head>

<h1>Login</h1>

<section>

    <form action="?/login"  method="POST">
        <div role="group">
            <input name="email" placeholder="email" type="text">
            <input name="password" placeholder="Password" type="password">
            <button type="submit">Submit</button>
        </div>
    </form>

</section>

