<script lang="ts">
  import {DataHandler} from '@vincjo/datatables';
  import {page} from '$app/stores';
  import Search from '$lib/components/datatable/Search.svelte';
  import ThFilter from '$lib/components/datatable/ThFilter.svelte';
  import ThSort from '$lib/components/datatable/ThSort.svelte';
  import RowCount from '$lib/components/datatable/RowCount.svelte';
  import RowsPerPage from '$lib/components/datatable/RowsPerPage.svelte';
  import Pagination from '$lib/components/datatable/Pagination.svelte';

  const handler = new DataHandler($page.data.profile.personal_leader_board.nemesis, {rowsPerPage: 5});
  const rows = handler.getRows();

</script>


<div class="table-container space-y-4">
  <div class="text-center font-bold text-2xl mb-4">Nemesis: Those that have beaten you</div>
  <table class="table table-hover table-compact table-auto w-full ">
    <thead>
    <tr>
      <ThSort {handler} orderBy="player.email">Player</ThSort>
      <ThSort {handler} orderBy="winsAgainstYou"># of losses</ThSort>

    </tr>
    <tr>
      <!-- <ThFilter {handler} filterBy="first_name" />
      <ThFilter {handler} filterBy="last_name" />
      <ThFilter {handler} filterBy="email" /> -->
    </tr>
    </thead>
    <tbody>
    {#each $rows as row}
      <tr>
        <td>{row.player.email}</td>
        <td>{row.winsAgainstYou}</td>

      </tr>
    {/each}
    </tbody>
  </table>
  <footer class="flex justify-between">
    <Search {handler} />
    <RowsPerPage {handler} />
    <RowCount {handler} />
    <Pagination {handler} />
  </footer>
</div>

