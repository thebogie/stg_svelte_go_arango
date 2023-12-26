<script lang="ts">
  import BarChart from '$lib/d3/BarChart.svelte';

  export let data;
  // export let player;
  console.log('PLAYER' + JSON.stringify(data.props.signed_in));
  console.log('COUNTRIES' + JSON.stringify(data.props.countries));


</script>


<svelte:head>
  <title>Profile</title>
</svelte:head>
<p>The passed data is: {data.props.countries}</p>

<div class="card p-4">Games Played

 <!-- <BarChart {data} />  -->

</div>
<h1>Profile</h1>
<p></p>


