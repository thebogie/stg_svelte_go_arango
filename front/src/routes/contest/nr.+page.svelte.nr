<script lang="ts">
  import { superForm } from 'sveltekit-superforms/client';
  import type { PageData } from './$types.js';
  import SveltyPicker from 'svelty-picker';


  export let data: PageData;

  const { form, errors, enhance, message } = superForm(data.loginForm, {
    resetForm: true
  });

  const {
    form: registerForm,
    errors: registerErrors,
    enhance: registerEnhance,
    message: registerMessage
  } = superForm(data.registerForm, {
    resetForm: true
  });
</script>

<SveltyPicker  />

{#if $message}<h3>{$message}</h3>{/if}

<form method="POST" action="?/login" use:enhance>
  E-mail: <input name="email" type="email" bind:value={$form.email} />
  Password:
  <input name="password" type="password" bind:value={$form.password} />
  <button>Submit</button>
</form>

<hr />

{#if $registerMessage}<h3>{$registerMessage}</h3>{/if}

<form method="POST" action="?/register" use:registerEnhance>
  E-mail: <input name="email" type="email" bind:value={$registerForm.email} />
  Password:
  <input name="password" type="password" bind:value={$registerForm.password} />
  Confirm password:
  <input
    name="confirmPassword"
    type="password"
    bind:value={$registerForm.confirmPassword} />
  <button>Submit</button>
</form>