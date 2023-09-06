import {writable} from 'svelte/store';

import type {IPlayer} from '$lib/interfaces/player.interface';

export const playerData = writable<IPlayer>({_id: "", _key: "", accessToken: undefined, email: "", firstname: "", password: "", rev: "", roles: undefined});