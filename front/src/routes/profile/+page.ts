import {fail, redirect} from "@sveltejs/kit";
import {getPlayerTotalResults} from "$lib/services/player.service";

/*
export async function load({}) {
    console.log("profile page.ts");



    let playerGamesPlayed = {};
    try {
        playerGamesPlayed = await getPlayerTotalResults();
    } catch (err: any) {
        console.log('profile error: ' + err.message);
        return fail(401, {
            error: err.message
        });
    }
    return {games: playerGamesPlayed};
}
*/
