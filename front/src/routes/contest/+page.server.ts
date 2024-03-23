import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { z } from 'zod';
import { message } from 'sveltekit-superforms';
import { fail } from '@sveltejs/kit';

// Define outside the load function so the adapter can be cached
const schema = z.object({
	name: z.string().default('Hello world!'),
	start_date_time: z.string().datetime({ offset: true }),
	end_date_time: z.string().datetime({ offset: true })
});

export const load = async () => {
	const form = await superValidate(zod(schema));
	form.data.start_date_time = '2024-02-02T18:00:00+05:00';
	form.data.end_date_time = '2024-02-02T20:00:00+05:00';
	// Always return { form } in load functions
	return { form };
};

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, zod(schema));
		console.log(form);

		if (!form.valid) {
			// Again, return { form } and things will just work.
			return fail(400, { form });
		}

		// TODO: Do something with the validated form.data

		// Display a success status message
		return message(form, 'Form posted successfully!');
	}
};
