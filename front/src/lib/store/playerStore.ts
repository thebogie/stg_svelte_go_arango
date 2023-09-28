import {writable} from 'svelte/store';

import type {IPlayer} from '$lib/interfaces/player.interface';
const initialPlayers: IPlayer = {accessToken: undefined, email: "", password: ""};
export const playerData = writable<IPlayer>(initialPlayers)

{

    console.log("STORE:" + JSON.stringify(initialPlayers));

};

//{_id: "", _key: "", accessToken: undefined, email: "", firstname: "", password: "", rev: "", roles: undefined}