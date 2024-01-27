import {createNemesisLeaderboard, getPlayerTotalResults} from '$lib/services/player.service';
import type {INemesisLeaderboard, IPlayer} from '$lib/interfaces/player';
import type {IContest} from '$lib/interfaces/contest';
import {redirect} from '@sveltejs/kit';

export const load = async ({ locals }) => {
	let total_results : IContest[] = [];
	let signed_in: IPlayer = {};
	let personal_leader_board: INemesisLeaderboard = {
		players: [],
		nemesis: [],
		punchingBag: [],
		winnersTogether: [],
		losersTogether: []
	};
	console.log('profile page.server.ts');

	if (locals?.player !== undefined) {
		signed_in = locals.player;

		try {
			total_results = await getPlayerTotalResults(signed_in);

			personal_leader_board = await createNemesisLeaderboard(signed_in, total_results);


		} catch (err) {
			console.error(err); // Add your error handling here
			throw redirect(303, '/login');
		}
	}

	return {
		login: signed_in,
		total_results: total_results,
		personal_leader_board: personal_leader_board
	};
};
