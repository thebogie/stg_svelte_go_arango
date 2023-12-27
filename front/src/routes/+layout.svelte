<script lang="ts">
  import {page} from '$app/stores';
  import '../app.postcss';
  import {AppShell, AppBar} from '@skeletonlabs/skeleton';
  import {computePosition, autoUpdate, flip, shift, offset, arrow} from '@floating-ui/dom';
  import {storePopup} from '@skeletonlabs/skeleton';
  import {enhance} from '$app/forms';

  storePopup.set({computePosition, autoUpdate, flip, shift, offset, arrow});
  console.log('LAYOUT  SVELTE');
  console.log('LAYOUT PROPS: ' + JSON.stringify($page.data.player));
</script>

<!-- App Shell -->
<AppShell>
  <svelte:fragment slot="header">
    <!-- App Bar -->
    <AppBar>
      <svelte:fragment slot="lead">
        <strong class="text-xl uppercase">Smacktalk Gaming</strong>
      </svelte:fragment>
      <svelte:fragment slot="trail">
        {#if $page.data.player}
          Welcome: {$page.data.player.email}
          <p></p>
          <form class="logout" action="/logout" method="POST" use:enhance>
            <button type="submit" > Logout </button>
          </form>

        {:else}

          <a class="text-xl" href="login">Login</a>
        {/if}


      </svelte:fragment>
    </AppBar>
  </svelte:fragment>
  <!-- Page Route Content -->
  <slot />
</AppShell>
