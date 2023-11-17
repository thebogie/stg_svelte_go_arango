<script lang="ts">

    export let data : App.PageData;
    console.log("FUCKING DATA" + JSON.stringify(data));
</script>

<svelte:head>
    <title>Profile</title>
</svelte:head>

<h1>Profile ${JSON.stringify(data)}</h1>
<p></p>