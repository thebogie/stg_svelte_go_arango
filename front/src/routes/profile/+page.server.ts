import { redirect } from '@sveltejs/kit';
import {getPlayerTotalResults, loginPlayer} from '$lib/services/player.service';

export const load = async ({ locals }) => {
	let total_results ;

	console.log('profile page.server.ts');

	if (locals?.player !== undefined) {
		const signed_in = locals.player;

		try {
			total_results = await getPlayerTotalResults(locals.player);


		} catch (err) {
			console.error(err); // Add your error handling here
		}

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

		return {  signed_in, countries, total_results  };
	}
};
