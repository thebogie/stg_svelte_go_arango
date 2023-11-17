import type { PageServerLoad, Actions } from './$types';
import { redirect, fail } from '@sveltejs/kit';
import {getPlayerGamesPlayed} from '$lib/services/player.service';


export const load: PageServerLoad = async (event) => {
	console.log('profile page.server.ts');
	const user = JSON.parse(event.locals.player);
	if (user) {
	} else {
		throw redirect(302, '/login');
	}

	let token = "";
	let playerGamesPlayed = {};
	try {
		playerGamesPlayed = await getPlayerGamesPlayed(user.accessToken) ;
	} catch (err: any) {
		console.log('profile error: ' + err.message);
		return fail(401, {
			error: err.message
		});
	}
	return playerGamesPlayed;
};
