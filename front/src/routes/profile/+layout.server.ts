import {createNemesisLeaderboard, getPlayerTotalResults} from '$lib/services/player.service';
import type {INemesisLeaderboard, IPlayer} from '$lib/interfaces/player';
import type {IContest, IProfile} from '$lib/interfaces/contest';
import {redirect} from '@sveltejs/kit';
import type {LayoutServerLoad} from '../../../.svelte-kit/types/src/routes/$types';

export const load: LayoutServerLoad = async ({ locals }) => {
    let total_results : IContest[] = [];
    let signed_in: IPlayer = {};
    let personal_leader_board: INemesisLeaderboard = {
        players: [],
        nemesis: [],
        owned: [],
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


    const profile: IProfile = {signed_in: signed_in, total_results: total_results, personal_leader_board: personal_leader_board};
    return {
        profile: profile
    };
};