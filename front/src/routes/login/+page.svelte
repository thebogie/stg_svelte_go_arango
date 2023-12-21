<script lang="ts">
    //import type {ActionData} from './$types';
    // import {enhance} from '$app/forms';
    import {getCurrentPlayer, logged_in_player} from "$lib/stores/playerStore";
    import Cookies from "js-cookie";
    import {goto} from "$app/navigation";
    import {loginPlayer} from "$lib/services/player.service";


    console.log("login page.svelte");

</script>

<svelte:head>
    <title>Login</title>
</svelte:head>

<h1>Login</h1>

<section>

    <form method="POST">
        <div role="group">
            <input type="text" name="email" placeholder="email">
            <input type="password" name="password" placeholder="Password">
            <button type="submit">Submit</button>
        </div>
    </form>

</section>

