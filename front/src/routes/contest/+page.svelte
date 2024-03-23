<script>
	import SveltyPicker from 'svelty-picker';

	import { superForm } from 'sveltekit-superforms';
	import SuperDebug from 'sveltekit-superforms';

	export let data;

	// Client API:
	const { form } = superForm(data.form);


</script>


<SuperDebug data={$form} />
<form method="POST">
	<label for="name">Name</label>
	<input class="text-purple-700 text-opacity-100" type="text" name="name" bind:value={$form.name} />

	<label for="start_date_time">Start Date+Time</label>
	<SveltyPicker formatType='standard' inputClasses="form-control text-purple-700 text-opacity-100"
								format="yyyy-mm-ddThh:ii:00Z" displayFormat="yyyy-mm-dd hh:ii" name="start_date_time" bind:value={$form.start_date_time}
	></SveltyPicker>

	<label for="end_date_time">End Date+Time</label>
	<SveltyPicker formatType='standard' inputClasses="form-control text-purple-700 text-opacity-100"
								format="yyyy-mm-ddThh:ii:00Z" displayFormat="yyyy-mm-dd hh:ii" name="end_date_time"  bind:value={$form.end_date_time}
	></SveltyPicker>

	<div>
		<button>Submit</button>
	</div>
</form>