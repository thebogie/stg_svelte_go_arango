import {fail, redirect} from "@sveltejs/kit";
import {getPlayerGamesPlayed} from "$lib/services/player.service";

/*
export async function load({cookies}) {
    console.log("profile page.ts");

    const user  = cookies.get('loggedinplayer');
    if (user) {
    } else {
        throw redirect(302, '/login');
    }

    let token = "";
    let playerGamesPlayed = {};
    try {
        playerGamesPlayed = await getPlayerGamesPlayed(user.accessToken);
    } catch (err: any) {
        console.log('profile error: ' + err.message);
        return fail(401, {
            error: err.message
        });
    }
    return {games: playerGamesPlayed};
}

 */