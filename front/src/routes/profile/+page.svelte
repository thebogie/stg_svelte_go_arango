<script lang="ts">
  import BarChart from '$lib/d3/BarChart.svelte';
  import {type PaginationSettings, Paginator, Table, tableMapperValues, type TableSource} from '@skeletonlabs/skeleton';
  import Widget from './Widget.svelte';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';

  export let data;

  let widgets = [
    { id: 1, name: 'Alice', email: 'alice@example.com' },
    { id: 2, name: 'Bob', email: 'bob@example.com' },
    { id: 3, name: 'Charlie', email: 'charlie@example.com' },
    // ... more data
  ];
  let currentPage = 1;
  let itemsPerPage = 5;

  const handlePageChange = (newPage: number) => {
    currentPage = newPage;
    goto(`/dashboard?page=${newPage}`);
  };

  let source = [
    { id: 1, name: 'Alice', email: 'alice@example.com' },
    { id: 2, name: 'Bob', email: 'bob@example.com' },
    { id: 3, name: 'Charlie', email: 'charlie@example.com' },
    // ... more data
  ];



</script>


<svelte:head>
  <title>Profile</title>
</svelte:head>

<div class="dashboard-container">
  {#each widgets.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage) as widget}
    <widget {...widget} />
  {/each}

  {#if widgets.length > itemsPerPage}
    <pagination
      totalPages={Math.ceil(widgets.length / itemsPerPage)}
      currentPage={currentPage}
      onPageChange={handlePageChange}
    />
  {/if}
</div>

<style>
    /* Customize dashboard layout */
</style>


