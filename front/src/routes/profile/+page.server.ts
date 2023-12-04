import type { PageServerLoad, Actions } from './$types';
import { redirect, fail } from '@sveltejs/kit';
import {getPlayerTotalResults} from '$lib/services/player.service';


export const load: PageServerLoad = async (event) => {
	console.log('profile page.server.ts');
	const user = JSON.parse(event.locals.player);
	if (user) {
	} else {
		throw redirect(302, '/login');
	}

	let token = "";
	let playerResults = {};
	try {
		playerResults = await getPlayerTotalResults(user.accessToken, user._key) ;
	} catch (err: any) {
		console.log('profile error: ' + err.message);
		return fail(401, {
			error: err.message
		});
	}
	return playerResults;
};
