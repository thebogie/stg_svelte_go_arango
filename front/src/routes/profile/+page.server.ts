import { redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	console.log('profile page.serfver.ts');

	if (locals?.player !== undefined) {
		const signed_in = locals.player;

		let countries = [
			{ country: 'China', population: 1439324 },
			{ country: 'India', population: 1380004 },
			{ country: 'United States of America', population: 331003 },
			{ country: 'Indonesia', population: 273524 },
			{ country: 'Pakistan', population: 220892 },
			{ country: 'Brazil', population: 212559 },
			{ country: 'Nigeria', population: 206140 },
			{ country: 'Bangladesh', population: 164689 },
			{ country: 'Russian Federation', population: 145934 },
			{ country: 'Mexico', population: 128933 }
		];

		return { props: { signed_in, countries } };
	}
};
