import type { LayoutServerLoad } from './$types';
import type {IPlayer} from '$lib/interfaces/player';

export const load: LayoutServerLoad = async ({ cookies }) => {
    console.log("LAYOUT SERVER");
    let readcookie = cookies.get('player');
    if ( readcookie === undefined || readcookie === "") {
        return {player: undefined}
    } else {
        let player: IPlayer = JSON.parse(readcookie);

        return {player: player};
    }
};