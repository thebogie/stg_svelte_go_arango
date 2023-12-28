<script lang="ts">
  import BarChart from '@lib/d3/BarChart.svelte';
  import {page} from '$app/stores';
  export let data;
  // export let player;
  //console.log('PLAYER' + JSON.stringify($page.data.signed_in));
  //console.log('COUNTRIES' + JSON.stringify($page.data.countries));


</script>


<svelte:head>
  <title>Profile</title>
</svelte:head>

<div class="card p-4">Total Wins, Draws and Losses

  <BarChart {$page} />

</div>



