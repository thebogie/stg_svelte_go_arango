<script lang="ts">

  import { TabGroup, Tab, TabAnchor } from '@skeletonlabs/skeleton';
  import DatatableNemesis from '$lib/components/datatable/DatatableNemesis.svelte';
  import DatatableOwned from '$lib/components/datatable/DatatableOwned.svelte';


  let tabSet: number = 0;
</script>

<svelte:head>
  <title>Profile</title>
</svelte:head>

<TabGroup>
  <Tab bind:group={tabSet} name="tab1" value={0}>
    <svelte:fragment slot="lead">Nemesis</svelte:fragment>
  </Tab>
  <Tab bind:group={tabSet} name="tab2" value={1}>Owned</Tab>
  <Tab bind:group={tabSet} name="tab3" value={2}>Won Together</Tab>
  <!-- Tab Panels --->
  <svelte:fragment slot="panel">
    {#if tabSet === 0}
      <DatatableNemesis/>
    {:else if tabSet === 1}
      <DatatableOwned />
    {:else if tabSet === 2}

    {/if}
  </svelte:fragment>
</TabGroup>



