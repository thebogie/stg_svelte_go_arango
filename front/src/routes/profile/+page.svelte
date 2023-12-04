<script lang="ts">

    export let data : App.PageData;
    console.log("FUCKING DATA" + JSON.stringify(data));
</script>

<svelte:head>
    <title>Profile</title>
</svelte:head>


<div class="card p-4">Games Played</div>
<h1>Profile ${JSON.stringify(data)}</h1>
<p></p>