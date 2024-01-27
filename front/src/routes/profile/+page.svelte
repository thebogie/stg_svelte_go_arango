<script lang="ts">
  import BarChart from '$lib/d3/BarChart.svelte';
  import {type PaginationSettings, Paginator, Table, tableMapperValues, type TableSource} from '@skeletonlabs/skeleton';


  export let data;


  let source = [
    { id: 1, name: 'Alice', email: 'alice@example.com' },
    { id: 2, name: 'Bob', email: 'bob@example.com' },
    { id: 3, name: 'Charlie', email: 'charlie@example.com' },
    // ... more data
  ];

  let npaginationSettings = {
    page: 0,
    limit: 5,
    size:  data.personal_leader_board.nemesis.length,
    amounts: [ 5, 10],
  } satisfies PaginationSettings;

  let pbpaginationSettings = {
    page: 0,
    limit: 5,
    size:  data.personal_leader_board.nemesis.length,
    amounts: [ 5, 10],
  } satisfies PaginationSettings;

  $: nPaginationSource = data.personal_leader_board.nemesis.slice(
    npaginationSettings.page * npaginationSettings.limit,
    npaginationSettings.page * npaginationSettings.limit + npaginationSettings.limit
  ).map(item => [item.player.firstname, item.winsAgainstYou]);

  $: pbPaginationSource = data.personal_leader_board.punchingBag.slice(
    pbpaginationSettings.page * pbpaginationSettings.limit,
    pbpaginationSettings.page * pbpaginationSettings.limit + pbpaginationSettings.limit
  ).map(item => [item.player.firstname, item.lostAgainstYou]);





</script>


<svelte:head>
  <title>Profile</title>
</svelte:head>

<!-- <BarChart /> -->

<div class="card">
  <header class="card-header">Nemesis</header>
  <section class="p-4">
    <Table source={{ head: ['id', 'name'], body: nPaginationSource }}  />
    <Paginator {...npaginationSettings} on:page={(e) => (npaginationSettings.page = e.detail)} />

  </section>
  <footer class="card-footer">(footer)</footer>
</div>

<div class="card">
  <header class="card-header">Punching Bag</header>
  <section class="p-4">
    <Table source={{ head: ['id', 'name'], body: pbPaginationSource }}  />
    <Paginator {...pbpaginationSettings} on:page={(e) => (pbpaginationSettings.page = e.detail)} />

  </section>
  <footer class="card-footer">(footer)</footer>
</div>


