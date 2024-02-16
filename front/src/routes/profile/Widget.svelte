<script lang="ts">
  export let title: string;
  export let content: any;
</script>

<div class="skeleton-wrapper">
  <div class="skeleton-header">{title}</div>
  <div class="skeleton-content" >{@html content}</div>
</div>

<style>
    /* Customize widget appearance */
</style>